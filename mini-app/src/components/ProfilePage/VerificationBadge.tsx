import { FC } from "react";

interface VerificationBadgeProps {
  isVerified: boolean;
}

export const VerificationBadge: FC<VerificationBadgeProps> = ({ isVerified }) => {
  return (
    <div className="flex items-center space-x-2">
      {isVerified ? (
        <>
          <span className="text-green-500">✓</span>
          <span className="text-[var(--tg-theme-text-color)]">Верифицирован</span>
        </>
      ) : (
        <>
          <span className="text-red-500">✗</span>
          <span className="text-[var(--tg-theme-text-color)]">Не верифицирован</span>
        </>
      )}
    </div>
  );
};