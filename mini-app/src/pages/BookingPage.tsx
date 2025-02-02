import { AppRoot, Cell, List, Button, Progress } from '@telegram-apps/telegram-ui';
import { FaChevronRight } from 'react-icons/fa';
import { services } from '../mock/services';
import styles from './BookingPage.module.css';

interface Service {
  id: number;
  name: string;
  duration: number;
  price: number;
}

export const BookingPage = () => {
  return (
    <AppRoot>
      <div className={styles.container}>
        {/* Заголовок */}
        <h1 className={styles.title}>Выберите услугу</h1>

        {/* Прогресс-бар */}
        <div className={styles.progressContainer}>
          <Progress value={20} />
          <span className={styles.progressText}>20%</span>
        </div>

        {/* Список услуг */}
        <List className={styles.list}>
          {services.map((service: Service) => (
            <Cell
              key={service.id}
              className={styles.serviceCard}
              before={
                <div className={styles.serviceIcon}>
                  <div className={styles.scissorsIcon}>
                    <svg viewBox="0 0 24 24">
                      <path d="M6 21h12v-2H6v2zM5 4.5v3h6.25v16h1.5v-16H19v-3H5z"/>
                    </svg>
                  </div>
                </div>
              }
              subtitle={`${service.duration} мин · ${service.price} руб`}
              after={<FaChevronRight className={styles.chevron} />}
            >
              <span className={styles.serviceName}>{service.name}</span>
            </Cell>
          ))}
        </List>

        {/* Кнопка "Продолжить" */}
        <div className={styles.footer}>
          <Button 
            size="l"
            className={styles.continueButton}
          >
            Продолжить
          </Button>
        </div>
      </div>
    </AppRoot>
  );
};