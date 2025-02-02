-- Удаляем таблицы в порядке зависимости
DROP TABLE IF EXISTS audit_logs CASCADE;
DROP TABLE IF EXISTS notifications CASCADE;
DROP TABLE IF EXISTS payments CASCADE;
DROP TABLE IF EXISTS bookings CASCADE;
DROP TABLE IF EXISTS schedule CASCADE;
DROP TABLE IF EXISTS masters_services CASCADE;
DROP TABLE IF EXISTS services CASCADE;
DROP TABLE IF EXISTS users CASCADE;

-- Удаляем расширения (если больше не нужны)
DROP EXTENSION IF EXISTS CITEXT;
DROP EXTENSION IF EXISTS pgcrypto;
DROP INDEX IF EXISTS idx_users_telegram_id;