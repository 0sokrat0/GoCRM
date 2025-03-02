import { FC, ReactNode, useEffect } from "react";
import { themeParams } from "@telegram-apps/sdk";
import { DynamicGradientBackground } from "@/components/DynamicGradientBackground";

interface LayoutProps {
  children: ReactNode;
}

export const Layout: FC<LayoutProps> = ({ children }) => {
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
        {children}
      </div>
    </div>
  );
};
