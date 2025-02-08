import { useState, useEffect, useCallback } from 'react';
import { AppRoot, Cell, List, Button, Progress } from '@telegram-apps/telegram-ui';
import { FaChevronRight } from 'react-icons/fa';
import { useNavigate } from 'react-router-dom';
import styles from './ServicePage.module.css';

interface Service {
  ServiceID: string;
  Name: string;
  Description: string;
  Price: number;
  Duration: number;
}

// ⏬ Основная компонента
export const ServicePage = () => {
  const [services, setServices] = useState<Service[]>([]);
  const [loading, setLoading] = useState(true);
  const [selectedService, setSelectedService] = useState<string | null>(null);
  const navigate = useNavigate();

  const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080/api/v1";

  useEffect(() => {
    const fetchServices = async () => {
      try {
        console.log("🔍 API URL:", API_URL); // Лог для проверки
        const response = await fetch(`${API_URL}/services`);
        const result = await response.json();
  
        if (result.status === "success" && Array.isArray(result.data)) {
          setServices(result.data);
        } else {
          console.error("❌ API вернул некорректные данные:", result);
        }
      } catch (error) {
        console.error("Ошибка при загрузке услуг:", error);
      } finally {
        setLoading(false);
      }
    };
  
    fetchServices();
  }, []);

  // ✅ Функция выбора услуги (оптимизировано)
  const handleSelectService = useCallback((serviceId: string) => {
    setSelectedService((prev) => (prev === serviceId ? null : serviceId));
  }, []);

  return (
    <AppRoot>
      {/* ✅ Линия прогресса (фиксированная) */}
      <div className={styles.progressBar}>
        <Progress value={selectedService ? 40 : 20} />
      </div>

      <div className={styles.container}>
        <h1 className={styles.title}>Выберите услугу</h1>

        {loading ? (
          <p className={styles.loadingText}>Загрузка...</p>
        ) : services.length > 0 ? (
          <List className={styles.list}>
            {services.map((service) => (
              <Cell
                key={service.ServiceID}
                className={`${styles.serviceCard} ${selectedService === service.ServiceID ? styles.selected : ''}`}
                subtitle={
                  selectedService === service.ServiceID
                    ? <span className={styles.selectedSubtitle}>✅ Выбрано</span> // ✅ Улучшено визуально
                    : `${service.Duration} мин · ${service.Price} руб`
                }
                after={<FaChevronRight className={styles.chevron} />}
                onClick={() => handleSelectService(service.ServiceID)}
              >
                <span className={styles.serviceName}>{service.Name}</span>
              </Cell>
            ))}
          </List>
        ) : (
          <p className={styles.noServicesText}>Нет доступных услуг</p>
        )}

        {/* ✅ Кнопки навигации (исправлено) */}
        <div className={styles.footer}>
          <Button
            size="l"
            className={styles.continueButton}
            disabled={!selectedService}
            onClick={() => navigate('/next-step', { state: { serviceId: selectedService } })} // ✅ Передаем `state`
          >
            Продолжить
          </Button>
          <Button
            size="l"
            color="secondary"
            className={styles.continueButton}
            onClick={() => navigate(-1)}
          >
            Назад
          </Button>
        </div>
      </div>
    </AppRoot>
  );
};
