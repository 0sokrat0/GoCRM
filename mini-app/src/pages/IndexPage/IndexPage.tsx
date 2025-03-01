import { FC, useEffect } from "react";
import { themeParams } from "@telegram-apps/sdk";
import { ProfileCard } from "@/components/HomePage/ProfileCard";
import { ScheduleCard } from "@/components/HomePage/ScheduleCard";
import { HoursCard } from "@/components/HomePage/HoursCard";
import { DynamicGradientBackground } from "@/components/DynamicGradientBackground";


export const IndexPage: FC = () => {
  useEffect(() => {
    if (themeParams.mount.isAvailable()) {
      themeParams.mount();
    }
    if (themeParams.bindCssVars.isAvailable()) {
      themeParams.bindCssVars();
    }
  }, []);

  return (
    <div className="relative min-h-screen p-4 bg-tg-bg text-tg-text overflow-hidden">
      <DynamicGradientBackground />
      <div className="space-y-6 relative z-10">
        <ProfileCard avatarSrc="/images/Graff.jpg" title="GRAFF" />
        <ScheduleCard />
        <HoursCard />
      </div>
    </div>
  );
};
