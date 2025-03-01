import { FC, useEffect, useState } from "react";

export const RandomGradientBackground: FC = () => {
  const gradients = [
    "bg-gradient-to-r from-amber-200 to-yellow-400",
    "bg-gradient-to-r from-indigo-400 to-cyan-400",
    "bg-gradient-to-r from-emerald-400 to-cyan-400",
  ];

  const [gradientClass, setGradientClass] = useState("");

  // Функция смены градиента
  const changeGradient = () => {
    const randomGradient = gradients[Math.floor(Math.random() * gradients.length)];
    setGradientClass(randomGradient);
  };

  useEffect(() => {
    // Начальная установка
    changeGradient();
    // Смена градиента каждые 10 секунд (10000 мс)
    const interval = setInterval(changeGradient, 1000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div
      aria-hidden="true"
      className={`absolute top-[-1rem] sm:top-[-2rem] left-1/2 -z-10 -translate-x-1/2 ${gradientClass} animate-gradientMove sm:blur-3xl blur-xl transition-all duration-500`}
      style={{
        width: "90vw",
        maxWidth: "72.1875rem",
        aspectRatio: "1155/678",
        opacity: 0.25,
      }}
    />
  );
};
