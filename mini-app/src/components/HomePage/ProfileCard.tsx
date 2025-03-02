import { FC } from "react";
import { Avatar, List, Section, Cell } from "@telegram-apps/telegram-ui";
import { Link } from "@/components/Link/Link";

interface ProfileCardProps {
  avatarSrc: string;
  title: string;
}

export const ProfileCard: FC<ProfileCardProps> = ({ avatarSrc, title }) => {
  return (
    <div className="bg-white/10 backdrop-blur-lg rounded-3xl p-4 sm:p-6 shadow-md border border-white/20 transition-transform duration-200 hover:scale-105">
      <div className="flex flex-col sm:flex-row items-center gap-4 sm:gap-6 mb-4 sm:mb-6">
        <Avatar 
          size={96}
          src={avatarSrc} 
          className="border-2 border-[var(--tg-theme-hint-color)]/30 shadow-md w-20 h-20 sm:w-24 sm:h-24"
        />
        <div className="flex-1 text-center sm:text-left">
          <h1 className="text-2xl sm:text-3xl font-bold text-[var(--tg-theme-text-color)] leading-tight">
            {title}
          </h1>
        </div>
      </div>
      
      <List className="rounded-2xl sm:rounded-3xl overflow-hidden border border-[var(--tg-theme-hint-color)]/20 bg-[var(--tg-theme-secondary-bg-color)] backdrop-blur-sm">
        <Section 
          header={
            <div className="bg-[var(--tg-theme-secondary-bg-color)] text-[var(--tg-theme-text-color)] px-4 py-3 sm:px-6 sm:py-4">
              <span className="text-lg font-medium">Действия</span>
            </div>
          }
        >
          {[
            { to: "/servicePage", icon: "/images/calendar_24.svg", text: "Записаться", subtitle: "Выберите время и мастера" },
            { to: "/service", icon: "/images/actions_24.svg", text: "Услуги", subtitle: "Список услуг" },
            { to: "/history", icon: "/images/person_24.svg", text: "Профиль", subtitle: "История записей" },
          ].map((item) => (
            <Link key={item.to} to={item.to} className="block">
              <Cell
                before={
                  <img
                    className="w-7 h-7 sm:w-8 sm:h-8 mr-3"
                    src={item.icon}
                    alt={`Иконка ${item.text}`}
                  />
                }
                subtitle={
                  <span className="text-[var(--tg-theme-hint-color)] opacity-90">
                    {item.subtitle}
                  </span>
                }
                className="px-4 sm:px-6 py-3 hover:bg-[var(--tg-theme-bg-color)]/10 active:bg-[var(--tg-theme-bg-color)]/20 transition-colors"
              >
                <span className="text-[var(--tg-theme-text-color)]">{item.text}</span>
              </Cell>
            </Link>
          ))}
        </Section>
      </List>
    </div>
  );
};