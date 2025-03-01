import { FC, useEffect, useState } from "react";

export const DynamicGradientBackground: FC = () => {
  const gradients = [
    "bg-gradient-to-r from-indigo-400 to-cyan-400",
    "bg-gradient-to-r from-emerald-400 to-cyan-400",
    "bg-gradient-to-r from-slate-900 to-slate-700",
    "bg-gradient-to-r from-violet-500 to-purple-500",
    "bg-gradient-to-r from-amber-500 to-pink-500",
  ];

  const [currentGradient, setCurrentGradient] = useState(gradients[0]);
  const [nextGradient, setNextGradient] = useState<string | null>(null);
  const [fade, setFade] = useState(false);

  useEffect(() => {
    const changeGradient = () => {
      let newGradient = gradients[Math.floor(Math.random() * gradients.length)];
      // Гарантируем, что новый градиент отличается от текущего
      while (newGradient === currentGradient) {
        newGradient = gradients[Math.floor(Math.random() * gradients.length)];
      }
      setNextGradient(newGradient);
      setFade(true);
      // Через 1 секунду завершаем переход, устанавливая новый градиент как текущий
      setTimeout(() => {
        setCurrentGradient(newGradient);
        setFade(false);
        setNextGradient(null);
      }, 1000);
    };

    const interval = setInterval(changeGradient, 100000);
    return () => clearInterval(interval);
  }, [currentGradient, gradients]);

  return (
    <div className="absolute top-0 left-0 -z-10 w-screen h-screen opacity-30">
      {/* Текущий градиент */}
      <div
        className={`absolute top-0 left-0 w-full h-full transition-opacity duration-1000 ${currentGradient} sm:blur-3xl blur-xl`}
      />
      {/* Следующий градиент, который появляется при смене */}
      {nextGradient && (
        <div
          className={`absolute top-0 left-0 w-full h-full transition-opacity duration-1000 ${nextGradient} sm:blur-3xl blur-xl ${
            fade ? "opacity-100" : "opacity-0"
          }`}
        />
      )}
    </div>
  );
};
