import { FC, useMemo, useEffect, useState, useCallback } from "react";
import { Layout } from "@/components/Layout";
import { DisplayData, type DisplayDataRow } from "@/components/DisplayData/DisplayData";
import { useSignal, initData, type User, backButton } from "@telegram-apps/sdk-react";
import { Avatar, List, Button } from "@telegram-apps/telegram-ui";
import { TypeAnimation } from "react-type-animation";
import { getUserByTelegramID } from "@/api/userApi";
import { requestContact } from "@telegram-apps/sdk";

// Функция для преобразования данных пользователя в строки для отображения
function getUserRows(
  user: User,
  phone: string | null,
  handlePhoneRequest: () => void,
  loadingPhone: boolean
): DisplayDataRow[] {
  const rows: DisplayDataRow[] = [
    { title: "ID", value: user.id.toString() },
    { title: "Имя пользователя", value: user.username || "Не указано" },
    { title: "Имя", value: user.firstName || "Не указано" },
    { title: "Фамилия", value: user.lastName || "Не указано" },
  ];

  // Если номер телефона получен – отображаем его, иначе кнопку запроса номера
  rows.push({
    title: "Телефон",
    value: phone ? phone : (
      <Button 
        size="m" 
        onClick={handlePhoneRequest} 
        loading={loadingPhone}
        className="w-full"
      >
        Запросить номер
      </Button>
    ),
  });

  return rows;
}

export const ProfilePage: FC = () => {
  const initDataState = useSignal(initData.state);
  const [phone, setPhone] = useState<string | null>(null);
  const [loadingPhone, setLoadingPhone] = useState<boolean>(false);

  useEffect(() => {
    if (window.Telegram && window.Telegram.WebApp) {
      console.log("Telegram WebApp is initialized!");
    } else {
      console.log("Telegram WebApp is NOT initialized!");
    }
  }, []);

  useEffect(() => {
    if (backButton.isSupported()) {
      console.log("Back Button is supported!");
      backButton.mount();
      console.log("Back Button mounted:", backButton.isMounted());
      backButton.show();
      console.log("Back Button visible:", backButton.isVisible());
    } else {
      console.log("Back Button is NOT supported!");
    }
  }, []);

  useEffect(() => {
    const handleBackButtonClick = () => {
      console.log("Back button clicked!");
      window.history.back();
    };

    if (backButton.isSupported()) {
      backButton.onClick(handleBackButtonClick);
    }

    return () => {
      if (backButton.isSupported()) {
        backButton.offClick(handleBackButtonClick);
      }
    };
  }, []);

  // Функция обновления данных пользователя с сервера
  const updateUserData = useCallback(async () => {
    if (!initDataState || !initDataState.user) return;
    try {
      const updatedUser = await getUserByTelegramID(initDataState.user.id);
      setPhone(updatedUser.Phone);
    } catch (error) {
      console.error("Ошибка при обновлении данных пользователя:", error);
    }
  }, [initDataState]);

  // Инициализация данных пользователя при загрузке компонента
  useEffect(() => {
    updateUserData();
  }, [updateUserData]);

  // Обработчик для запроса номера телефона с использованием requestContact
  const handlePhoneRequest = useCallback(async () => {
    console.log("Button clicked: Requesting phone number");
    
    if (!requestContact.isAvailable()) {
      console.error("Метод запроса контактных данных не поддерживается");
      return;
    }
    
    try {
      setLoadingPhone(true);
      console.log("Starting contact request...");
      
      const contactData = await requestContact();
      console.log("Contact data received:", contactData);
      
      if (contactData && contactData.contact && contactData.contact.phone_number) {
        // Устанавливаем номер телефона напрямую
        setPhone(contactData.contact.phone_number);
        
        // Можно также отправить данные на сервер
        // Здесь может быть вызов API для сохранения номера телефона
        console.log("Phone number set:", contactData.contact.phone_number);
      }
    } catch (error) {
      console.error("Ошибка при запросе контактных данных:", error);
    } finally {
      setLoadingPhone(false);
    }
  }, []);

  const userRows = useMemo<DisplayDataRow[] | undefined>(() => {
    return initDataState && initDataState.user
      ? getUserRows(initDataState.user, phone, handlePhoneRequest, loadingPhone)
      : undefined;
  }, [initDataState, phone, handlePhoneRequest, loadingPhone]);

  return (
    <Layout>
      <div className="bg-white/10 backdrop-blur-lg rounded-3xl p-4 sm:p-6 shadow-md border border-white/20 transition-transform duration-200 hover:scale-105 max-w-3xl mx-auto">
        {/* Заголовок страницы */}
        <h1 className="text-2xl sm:text-3xl font-bold text-[var(--tg-theme-text-color)] mb-4 text-center">
          Профиль
        </h1>
        
        {/* Центрированный аватар */}
        <div className="flex justify-center mb-4">
          <Avatar
            size={96}
            src={initDataState?.user?.photoUrl}
            className="border-2 border-[var(--tg-theme-hint-color)]/30 shadow-md w-20 h-20 sm:w-24 sm:h-24"
          />
        </div>
        
        {/* Приветствие с эффектом печатания */}
        <div className="mb-6 text-center">
          <TypeAnimation
            sequence={[
              `Привет, ${initDataState?.user?.username || initDataState?.user?.firstName || "друг"}! 👋`,
              2000
            ]}
            wrapper="span"
            cursor={true}
            repeat={0}
            className="text-lg sm:text-xl font-medium text-[var(--tg-theme-text-color)]"
          />
        </div>
        
        {/* Отображение данных пользователя */}
        {userRows ? (
          <List className="rounded-2xl sm:rounded-3xl overflow-hidden border border-[var(--tg-theme-hint-color)]/20 bg-[var(--tg-theme-secondary-bg-color)] bg-opacity-30 backdrop-blur-sm">
            <DisplayData header="Информация о пользователе" rows={userRows} />
          </List>
        ) : (
          <div className="text-center text-[var(--tg-theme-hint-color)]">
            Данные пользователя недоступны.
          </div>
        )}
      </div>
    </Layout>
  );
};