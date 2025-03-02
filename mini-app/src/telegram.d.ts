// telegram.d.ts
interface TelegramWebAppUser {
  id: number;
  first_name: string;
  last_name?: string;
  username?: string;
  language_code?: string;
  is_premium?: boolean;
  allows_write_to_pm?: boolean;
  photo_url?: string;
}

interface TelegramWebAppInitData {
  query_id?: string;
  user?: TelegramWebAppUser;
  receiver?: TelegramWebAppUser;
  chat?: {
    id: number;
    type: string;
    title?: string;
    username?: string;
    photo_url?: string;
  };
  start_param?: string;
  can_send_after?: number;
  auth_date: number;
  hash: string;
}

interface TelegramWebApp {
  initData: string; // Строка с данными инициализации
  initDataUnsafe: TelegramWebAppInitData; // Парсированные данные
  openLink: (url: string) => void;
  ready: () => void;
  showAlert: (message: string, callback?: () => void) => void;
  showConfirm: (message: string, callback?: (confirmed: boolean) => void) => void;
  requestPhoneNumber?: () => void;
  // Добавьте другие методы и свойства, если они используются
}

interface TelegramInterface {
  WebApp: TelegramWebApp;
}

declare global {
  interface Window {
    Telegram?: TelegramInterface;
  }
}

export {};