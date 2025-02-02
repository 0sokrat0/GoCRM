// telegram.d.ts (создайте этот файл в корне или в каталоге с типами)
interface TelegramWebApp {
    openLink: (url: string) => void;
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
  