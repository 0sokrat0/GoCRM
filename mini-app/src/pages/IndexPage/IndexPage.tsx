import { Section, Cell, List, Avatar } from '@telegram-apps/telegram-ui';
import type { FC } from 'react';
import { Link } from '@/components/Link/Link.tsx';
import { Page } from '@/components/Page.tsx';
import styles from './IndexPage.module.css'; // CSS-модуль

export const IndexPage: FC = () => {
  return (
    <Page back={false}>
      <div className={styles.container}>
        {/* Блок профиля */}
        <div className={`${styles.profileCard} ${styles.gradientBorder} fadeIn`}>
          <div className={styles.profileHeader}>
            <Avatar
              size={96}
              src="/images/Graff.jpg" // путь от корня public
              className={styles.avatar}
            />
            <div className={styles.profileInfo}>
              <h1 className={styles.title}>GRAFF</h1>
            </div>
          </div>

          {/* Список действий */}
          <List className={`${styles.actionsList} ${styles.glowEffect} fadeIn`}>
            <Section header="Действия">
              <Link to="/booking" className={styles.link}>
                <Cell
                  before={<span className={styles.icon}>✂️</span>}
                  subtitle="Выберите время и мастера"
                  className={styles.actionCell}
                >
                  Записаться
                </Cell>
              </Link>
              <Link to="/service" className={styles.link}>
                <Cell
                  before={<span className={styles.icon}>💈</span>}
                  subtitle="Список услуг"
                  className={styles.actionCell}
                >
                  Услуги
                </Cell>
              </Link>
              <Link to="/history" className={styles.link}>
                <Cell
                  before={<span className={styles.icon}>📖</span>}
                  subtitle="Просмотр предыдущих записей"
                  className={styles.actionCell}
                >
                  История посещений
                </Cell>
              </Link>
            </Section>
          </List>
        </div>

        {/* Блок текущих записей */}
        <div className={`${styles.scheduleCard} ${styles.gradientBorder}  fadeIn`}>
          <Cell
            before={<span className={styles.scheduleIcon}>⏰</span>}
            subtitle="Сегодня в 15:30 - Мужская стрижка"
          >
            Ближайшая запись
          </Cell>
        </div>

        {/* Кнопка "Написать в Telegram" */}
        <div className={`${styles.hoursCard} ${styles.glowEffect}`}>
          <Link
            to="https://t.me/YourTelegramUsername"  // замените на ваш никнейм
            target="_blank"
            rel="noopener noreferrer"
            className={styles.link}
          >
            <Cell
              before={<span className={`${styles.icon} ${styles.telegramIcon} fadeIn`}>✈️</span>}
              subtitle="Свяжитесь с нами через Telegram"
              className={`${styles.actionCell} ${styles.telegramButton}`}
            >
              Написать в Telegram
            </Cell>
          </Link>
        </div>

        {/* Часы работы */}
        <div className={`${styles.hoursCard} ${styles.glowEffect} fadeIn`}>
          <h3 className={styles.hoursTitle}>Часы работы</h3>
          <p className={styles.hoursText}>Пн-Пт: 10:00 – 20:00</p>
          <p className={styles.hoursText}>Сб-Вс: 11:00 – 18:00</p>
        </div>
      </div>

      {/* Подвал */}
      <footer className={styles.footer}>
        <p>GRAFF © 2024</p>
      </footer>
    </Page>
  );
};
