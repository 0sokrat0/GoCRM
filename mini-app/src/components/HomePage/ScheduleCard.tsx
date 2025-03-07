import { FC } from "react";
import { Cell } from "@telegram-apps/telegram-ui";

export const ScheduleCard: FC = () => {
  return (
    <div className="bg-white/10 backdrop-blur-lg rounded-2xl p-4 shadow-md border border-white/20 transition-transform duration-200 hover:scale-105">
      <Cell 
        before={
          <img
            className="w-8 h-8 mr-3"
            src="/images/notifications_28.svg"
            alt="Иконка уведомлений"
          />
        }
        subtitle="Сегодня в 15:30 - Мужская стрижка"
        className="py-2 px-4 hover:block bg-[var(--tg-theme-secondary-bg-color)] rounded-2xl transition-colors duration-300"
      >
        Ближайшая запись
      </Cell>
    </div>
  );
};
