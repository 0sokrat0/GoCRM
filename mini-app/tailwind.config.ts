// tailwind.config.js
module.exports = {
  darkMode: "class", // включаем режим dark через класс
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
    "./public/index.html"
  ],
  theme: {
    extend: {
      colors: {
        // Основной фон Telegram
        'tg-bg': 'var(--tg-theme-bg-color)',
        // Дополнительный фон
        'tg-secondary-bg': 'var(--tg-theme-secondary-bg-color)',
        // Цвет текста
        'tg-text': 'var(--tg-theme-text-color)',
        // Цвет подсказок
        'tg-hint': 'var(--tg-theme-hint-color)',
        // Цвет ссылок (если понадобится)
        'tg-link': 'var(--tg-theme-link-color)',
      },
      letterSpacing: {
        'tg-title': '0.5px',
      },
      fontFamily: {
        'tg-default': 'var(--tg-theme-font-family)',
      },
      borderRadius: {
        'tg-rounded': 'var(--tg-theme-border-radius, 0.25rem)',
      },
      keyframes: {
        gradientMove: {
          '0%': { backgroundPosition: '0% 50%' },
          '50%': { backgroundPosition: '100% 50%' },
          '100%': { backgroundPosition: '0% 50%' },
        },
      },
      animation: {
        floating: 'floating 20s ease-in-out infinite',
        
      },
    },
  },
  plugins: [],
};
