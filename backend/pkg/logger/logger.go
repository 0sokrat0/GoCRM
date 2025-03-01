package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Основной объект логгера.
	Logger *zap.Logger

	// Уровень логирования, который можно менять динамически.
	atomicLevel zap.AtomicLevel

	once       sync.Once
	logQueue   chan logMessage
	wg         sync.WaitGroup
	queueSize  = 1000
	shutdownCh = make(chan struct{})
)

// logMessage описывает одно сообщение в очереди.
type logMessage struct {
	level  zapcore.Level
	msg    string
	fields []zap.Field
}

// Config задаёт параметры инициализации логгера.
type Config struct {
	Environment      string
	Color            bool
	OutputPaths      []string
	ErrorOutputPaths []string
	// Параметры ниже можно убрать/игнорировать, если caller берём вручную.
	AddCaller          bool
	CallerSkip         int
	StacktraceLevel    zapcore.Level
	EnableSampling     bool
	SamplingInitial    int
	SamplingThereafter int
}

var levelColors = map[zapcore.Level]*color.Color{
	zapcore.DebugLevel: color.New(color.FgCyan),
	zapcore.InfoLevel:  color.New(color.FgGreen),
	zapcore.WarnLevel:  color.New(color.FgYellow),
	zapcore.ErrorLevel: color.New(color.FgRed),
	zapcore.FatalLevel: color.New(color.FgMagenta),
	zapcore.PanicLevel: color.New(color.FgMagenta),
}

// Init инициализирует логгер и запускает воркер для асинхронной записи.
func Init(cfg Config) {
	once.Do(func() {
		cfg = validateConfig(cfg)

		// Создаём AtomicLevel, чтобы можно было менять уровень логирования в runtime
		atomicLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
		if strings.EqualFold(cfg.Environment, "production") {
			atomicLevel.SetLevel(zap.InfoLevel)
		}

		encoder := selectEncoder(cfg)
		core := createCore(cfg, encoder)

		Logger = buildLogger(cfg, core)
		logQueue = make(chan logMessage, queueSize)

		startLogWorker()
	})
}

// validateConfig заполняет пропущенные поля Config значениями по умолчанию.
func validateConfig(cfg Config) Config {
	if len(cfg.OutputPaths) == 0 {
		cfg.OutputPaths = []string{"stdout"}
	}
	if len(cfg.ErrorOutputPaths) == 0 {
		cfg.ErrorOutputPaths = []string{"stderr"}
	}
	if cfg.StacktraceLevel == 0 {
		cfg.StacktraceLevel = zapcore.ErrorLevel
	}
	// Если мы добавляем caller вручную, то AddCaller/CallerSkip можно не использовать.
	return cfg
}

// selectEncoder выбирает формат вывода (JSON или консольный) в зависимости от окружения.
func selectEncoder(cfg Config) zapcore.Encoder {
	if strings.EqualFold(cfg.Environment, "production") {
		return newProductionEncoder()
	}
	return newDevelopmentEncoder(cfg)
}

// createCore создаёт zapcore.Core с учётом уровня, энкодера, а также семплирования.
func createCore(cfg Config, encoder zapcore.Encoder) zapcore.Core {
	ws := createWriteSyncer(cfg.OutputPaths)

	// Уровень будет динамически управляться atomicLevel
	core := zapcore.NewCore(
		encoder,
		ws,
		atomicLevel,
	)
	if cfg.EnableSampling {
		core = applySampling(cfg, core)
	}
	return core
}

// applySampling включает семплирование логов, чтобы ограничить объём в случае спам-сообщений.
func applySampling(cfg Config, core zapcore.Core) zapcore.Core {
	initial := 100
	if cfg.SamplingInitial > 0 {
		initial = cfg.SamplingInitial
	}
	thereafter := 100
	if cfg.SamplingThereafter > 0 {
		thereafter = cfg.SamplingThereafter
	}
	return zapcore.NewSampler(core, time.Second, initial, thereafter)
}

// buildLogger формирует итоговый zap.Logger из core и дополнительных опций.
func buildLogger(cfg Config, core zapcore.Core) *zap.Logger {
	opts := []zap.Option{
		zap.ErrorOutput(createWriteSyncer(cfg.ErrorOutputPaths)),
		zap.AddStacktrace(cfg.StacktraceLevel),
	}
	return zap.New(core, opts...)
}

// newProductionEncoder формирует JSON-энкодер (обычно используется на продакшене).
func newProductionEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
}

// newDevelopmentEncoder формирует «читаемый» консольный энкодер (удобно при разработке).
func newDevelopmentEncoder(cfg Config) zapcore.Encoder {
	encCfg := zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	if cfg.Color {
		encCfg.EncodeLevel = colorLevelEncoder
	} else {
		encCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	}
	return zapcore.NewConsoleEncoder(encCfg)
}

// colorLevelEncoder добавляет цвет для каждого уровня логов.
func colorLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	c, ok := levelColors[l]
	if !ok {
		c = color.New(color.Reset)
	}
	enc.AppendString(c.Sprint(l.CapitalString()))
}

// createWriteSyncer создаёт общий WriteSyncer для указанных путей (stdout, stderr или файл).
func createWriteSyncer(paths []string) zapcore.WriteSyncer {
	syncers := make([]zapcore.WriteSyncer, 0, len(paths))
	for _, path := range paths {
		switch path {
		case "stdout":
			syncers = append(syncers, zapcore.Lock(os.Stdout))
		case "stderr":
			syncers = append(syncers, zapcore.Lock(os.Stderr))
		default:
			syncers = append(syncers, newFileSyncer(path))
		}
	}
	return zap.CombineWriteSyncers(syncers...)
}

// newFileSyncer открывает файл для логирования или пишет в stderr при ошибке.
func newFileSyncer(path string) zapcore.WriteSyncer {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open log file %s: %v\n", path, err)
		return zapcore.AddSync(os.Stderr)
	}
	return zapcore.AddSync(file)
}

// getCaller — вспомогательная функция для захвата информации о месте вызова.
func getCaller(skip int) zap.Field {
	_, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return zap.String("caller", "unknown")
	}
	return zap.String("caller", fmt.Sprintf("%s:%d", path.Base(file), line))
}

// enqueueLog — кладёт сообщение в очередь на асинхронную обработку.
func enqueueLog(level zapcore.Level, msg string, fields ...zap.Field) {
	select {
	case logQueue <- logMessage{level: level, msg: msg, fields: fields}:
	default:
		// Если очередь переполнена, лог теряем (или можно блокировать по вашему выбору).
		fmt.Printf("[DROPPED LOG] %s\n", msg)
	}
}

// startLogWorker — запускает одну горутину-воркер для обработки очереди.
func startLogWorker() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case logMsg := <-logQueue:
				writeLog(logMsg)
			case <-shutdownCh:
				// Дочитываем оставшиеся логи
				for len(logQueue) > 0 {
					writeLog(<-logQueue)
				}
				return
			}
		}
	}()
}

// writeLog — непосредственно вызывает методы zap.Logger (уже синхронно).
func writeLog(msg logMessage) {
	if Logger == nil {
		return
	}
	switch msg.level {
	case zapcore.DebugLevel:
		Logger.Debug(msg.msg, msg.fields...)
	case zapcore.InfoLevel:
		Logger.Info(msg.msg, msg.fields...)
	case zapcore.WarnLevel:
		Logger.Warn(msg.msg, msg.fields...)
	case zapcore.ErrorLevel:
		Logger.Error(msg.msg, msg.fields...)
	case zapcore.DPanicLevel, zapcore.PanicLevel:
		Logger.Panic(msg.msg, msg.fields...)
	case zapcore.FatalLevel:
		Logger.Fatal(msg.msg, msg.fields...)
	}
}

// Публичные методы для логирования с захватом caller вручную.
// Skip=2, чтобы пропустить сами обёртки Debug/Info/... и getCaller.
func Debug(msg string, fields ...zap.Field) {
	fields = append(fields, getCaller(2))
	enqueueLog(zapcore.DebugLevel, msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	fields = append(fields, getCaller(2))
	enqueueLog(zapcore.InfoLevel, msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	fields = append(fields, getCaller(2))
	enqueueLog(zapcore.WarnLevel, msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	fields = append(fields, getCaller(2))
	enqueueLog(zapcore.ErrorLevel, msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	fields = append(fields, getCaller(2))
	enqueueLog(zapcore.FatalLevel, msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	fields = append(fields, getCaller(2))
	enqueueLog(zapcore.PanicLevel, msg, fields...)
}

// Shutdown корректно завершает работу: закрывает канал и дожидается очистки очереди.
func Shutdown() {
	close(shutdownCh)
	wg.Wait()
}

// SetLogLevel позволяет менять уровень логирования в runtime.
// Например: logger.SetLogLevel(zapcore.ErrorLevel)
func SetLogLevel(l zapcore.Level) {
	atomicLevel.SetLevel(l)
}
