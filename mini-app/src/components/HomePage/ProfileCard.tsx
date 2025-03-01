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
      <div className="flex flex-col sm:flex-row items-center gap-5 mb-6">
        <Avatar 
          size={96} 
          src={avatarSrc} 
          className="border-2 border-white/30 shadow-md" 
        />
        <div className="flex-1">
          <h1 className="text-3xl sm:text-4xl font-bold text-[var(--tg-theme-text-color, #000000)]">
            {title}
          </h1>
        </div>
      </div>
      <List  className="rounded-xl overflow-hidden border border-white/20 bg-[var(--tg-theme-secondary-bg-color, #f5f5f5)] bg-opacity-50 ">
        <Section header="Действия">
          <Link to="/servicePage" className="block">
            <Cell
              before={
                <img 
                  className="w-6 h-6 mr-3 " 
                  src="public/images/calendar_24.svg" 
                  alt="Иконка календаря" 
                />
              }
              subtitle="Выберите время и мастера"
              className="py-2 px-4 hover:bg-white/10 transition-colors duration-300 "
            >
              Записаться
            </Cell>
          </Link>
          <Link to="/service" className="block">
            <Cell
              before={
                <img 
                  className="w-6 h-6 mr-3" 
                  src="public/images/actions_24.svg" 
                  alt="Иконка услуг" 
                />
              }
              subtitle="Список услуг"
              className="py-2 px-4 hover:bg-white/10 transition-colors duration-300"
            >
              Услуги
            </Cell>
          </Link>
          <Link to="/history" className="block">
            <Cell
              before={
                <img 
                  className="w-6 h-6 mr-3" 
                  src="public/images/person_24.svg" 
                  alt="Иконка профиля" 
                />
              }
              subtitle="История записей"
              className="py-2 px-4 hover:bg-white/10 transition-colors duration-300"
            >
              Профиль
            </Cell>
          </Link>
        </Section>
      </List>
    </div>
  );
};
