import { FC } from "react";

export const HoursCard: FC = () => {
  return (
    <div className="bg-white/10 backdrop-blur-lg rounded-xl p-4 text-center shadow-md border border-white/20 transition-transform duration-200 hover:scale-105">
      <h3 className="text-lg font-semibold text-[var(--tg-theme-text-color, #000000)] mb-2">
        Часы работы
      </h3>
      <p className="text-sm text-[var(--tg-theme-text-color, #000000)]">
        Пн-Пт: 10:00 – 20:00
      </p>
      <p className="text-sm text-[var(--tg-theme-text-color, #000000)]">
        Сб-Вс: 11:00 – 18:00
      </p>
    </div>
  );
};
