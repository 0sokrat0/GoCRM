-- Создаем необходимые расширения
CREATE EXTENSION IF NOT EXISTS pgcrypto;  -- Для функции gen_random_uuid()
CREATE EXTENSION IF NOT EXISTS CITEXT;    -- Для использования CITEXT (нечувствительный к регистру тип данных)

-- Таблица пользователей
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),  -- Первичный ключ, соответствует entity.ID
    role VARCHAR(10) DEFAULT 'client' CHECK (role IN ('client', 'master', 'admin')),
    telegram_id BIGINT UNIQUE NOT NULL,              -- Поле TelegramID
    username VARCHAR(50) NOT NULL,                    -- Username
    first_name VARCHAR(50) NOT NULL,                  -- FirstName
    last_name VARCHAR(50),                            -- LastName (может быть NULL)
    lang_code VARCHAR(10),                            -- LanguageCode (назван как lang_code в JSON)
    phone VARCHAR(20),                                -- Phone
    session_hash VARCHAR(250),                        -- SessionHash
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    login_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Таблица услуг
CREATE TABLE services (
    service_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) UNIQUE NOT NULL,
    description VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL CHECK (price >= 0),
    duration INT NOT NULL CHECK (duration > 0),  -- Длительность в минутах
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Таблица связи мастеров и услуг (связь многие ко многим)
CREATE TABLE masters_services (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    master_id UUID NOT NULL,
    service_id UUID NOT NULL,
    FOREIGN KEY (master_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (service_id) REFERENCES services(service_id) ON DELETE CASCADE
);

-- Таблица расписания мастеров (рабочее расписание)
CREATE TABLE schedule (
    schedule_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    master_id UUID NOT NULL,
    day_of_week INT NOT NULL CHECK (day_of_week BETWEEN 0 AND 6),  -- 0 = Воскресенье, 6 = Суббота
    start_time TIME NOT NULL,
    end_time TIME NOT NULL CHECK (end_time > start_time),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    FOREIGN KEY (master_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Таблица бронирований
CREATE TABLE bookings (
    booking_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    master_id UUID NOT NULL,
    service_id UUID NOT NULL,
    booking_time TIMESTAMP WITH TIME ZONE NOT NULL,
    status VARCHAR(20) NOT NULL CHECK (status IN ('pending', 'confirmed', 'canceled')),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (master_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (service_id) REFERENCES services(service_id) ON DELETE CASCADE
);

-- Таблица платежей
CREATE TABLE payments (
    payment_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id UUID NOT NULL,
    amount DECIMAL(10,2) NOT NULL CHECK (amount > 0),
    payment_method VARCHAR(50) NOT NULL CHECK (payment_method IN ('card', 'cash')),
    status VARCHAR(20) NOT NULL CHECK (status IN ('pending', 'paid', 'failed')),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    FOREIGN KEY (booking_id) REFERENCES bookings(booking_id) ON DELETE CASCADE
);

-- Таблица уведомлений
CREATE TABLE notifications (
    notification_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    type VARCHAR(50) NOT NULL CHECK (type IN ('email', 'sms', 'telegram')),
    message TEXT NOT NULL,
    status VARCHAR(20) NOT NULL CHECK (status IN ('sent', 'failed')),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Таблица аудиторских логов
CREATE TABLE audit_logs (
    log_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    action VARCHAR(255) NOT NULL,
    details TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Дополнительно можно создать уникальный индекс по telegram_id (хотя в определении уже стоит UNIQUE).
CREATE UNIQUE INDEX idx_users_telegram_id ON users (telegram_id);
