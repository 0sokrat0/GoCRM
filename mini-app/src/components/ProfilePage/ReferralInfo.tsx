import { FC } from "react";

interface ReferralInfoProps {
  hasReferrer: boolean;
}

export const ReferralInfo: FC<ReferralInfoProps> = ({ hasReferrer }) => {
  return (
    <div className="flex items-center space-x-2">
      <span className="text-[var(--tg-theme-text-color)]">
        {hasReferrer ? "Есть реферер" : "Нет реферера"}
      </span>
    </div>
  );
};