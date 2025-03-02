import { FC } from "react";
import { Avatar } from "@telegram-apps/telegram-ui";
import { User } from "@telegram-apps/sdk-react";

interface UserInfoProps {
  user: User;
}

export const UserInfo: FC<UserInfoProps> = ({ user }) => {
  return (
    <div className="flex flex-col items-center space-y-4">
      {/* Аватар пользователя */}
      <Avatar
        size={96}
        src={user.photoUrl}
        className="border-2 border-[var(--tg-theme-hint-color)]/30 shadow-md"
      />

      {/* Имя и фамилия */}
      <div className="text-center">
        <h2 className="text-2xl font-bold text-[var(--tg-theme-text-color)]">
          {user.firstName} {user.lastName}
        </h2>
        {user.username && (
          <p className="text-[var(--tg-theme-hint-color)]">@{user.username}</p>
        )}
      </div>

      {/* Дополнительная информация */}
      <div className="w-full space-y-2">
        <div className="flex justify-between">
          <span className="text-[var(--tg-theme-hint-color)]">ID:</span>
          <span className="text-[var(--tg-theme-text-color)]">{user.id}</span>
        </div>
        <div className="flex justify-between">
          <span className="text-[var(--tg-theme-hint-color)]">Премиум:</span>
          <span className="text-[var(--tg-theme-text-color)]">
            {user.isPremium ? "Да" : "Нет"}
          </span>
        </div>
        <div className="flex justify-between">
          <span className="text-[var(--tg-theme-hint-color)]">Язык:</span>
          <span className="text-[var(--tg-theme-text-color)]">
            {user.languageCode || "Не указано"}
          </span>
        </div>
      </div>
    </div>
  );
};