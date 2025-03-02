import { FC } from "react";

interface UserLevelProps {
  level: string;
}

export const UserLevel: FC<UserLevelProps> = ({ level }) => {
  const getLevelColor = (level: string) => {
    switch (level) {
      case "new":
        return "bg-blue-500";
      case "regular":
        return "bg-green-500";
      case "vip":
        return "bg-purple-500";
      default:
        return "bg-gray-500";
    }
  };

  return (
    <div className="flex items-center space-x-2">
      <div className={`w-4 h-4 rounded-full ${getLevelColor(level)}`} />
      <span className="text-[var(--tg-theme-text-color)] capitalize">
        {level}
      </span>
    </div>
  );
};