import { FC } from "react";

export const FloatingEffect: FC = () => {
  return (
    <div 
      aria-hidden="true"
      className="absolute top-0 left-0 w-32 h-32 bg-gradient-to-br from-blue-400 to-cyan-400 opacity-30 rounded-full animate-floatAround pointer-events-none"
    />
  );
};
