import { FC, useEffect, useState } from "react";

interface Color {
  h: number;
  s: number;
  l: number;
}

const lerp = (a: number, b: number, t: number) => a * (1 - t) + b * t;
const clamp = (num: number, min: number, max: number) =>
  Math.min(Math.max(num, min), max);

const interpolateColor = (current: Color, target: Color, t: number): Color => ({
  h: lerp(current.h, target.h, t),
  s: clamp(lerp(current.s, target.s, t), 30, 100),
  l: clamp(lerp(current.l, target.l, t), 20, 80),
});

const generateColor = (base?: Color): Color => {
  if (!base) {
    return {
      h: Math.floor(Math.random() * 360),
      s: 50 + Math.random() * 50,
      l: 40 + Math.random() * 20,
    };
  }
  
  return {
    h: (base.h + 30 + Math.random() * 60) % 360,
    s: clamp(base.s + (Math.random() * 20 - 10), 40, 100),
    l: clamp(base.l + (Math.random() * 10 - 5), 30, 70),
  };
};

const generateGradient = (current?: { start: Color; mid: Color; end: Color }) => ({
  start: generateColor(current?.start),
  mid: generateColor(current?.mid),
  end: generateColor(current?.end),
});

export const DynamicGradientBackground: FC = () => {
  const [currentGradient, setCurrentGradient] = useState(() => generateGradient());
  const [targetGradient, setTargetGradient] = useState(() => generateGradient(currentGradient));

  useEffect(() => {
    let animationFrameId: number;
    const duration = 10000;
    let startTime = Date.now();

    const animate = () => {
      const now = Date.now();
      const progress = (now - startTime) / duration;

      if (progress < 1) {
        setCurrentGradient((prev) => ({
          start: interpolateColor(prev.start, targetGradient.start, progress),
          mid: interpolateColor(prev.mid, targetGradient.mid, progress),
          end: interpolateColor(prev.end, targetGradient.end, progress),
        }));
        animationFrameId = requestAnimationFrame(animate);
      } else {
        const newTarget = generateGradient(targetGradient);
        setCurrentGradient(targetGradient);
        setTargetGradient(newTarget);
        startTime = Date.now();
        animationFrameId = requestAnimationFrame(animate);
      }
    };

    animationFrameId = requestAnimationFrame(animate);
    return () => cancelAnimationFrame(animationFrameId);
  }, [targetGradient]);

  const gradientStyle = {
    background: `linear-gradient(to right,
      hsl(${currentGradient.start.h}, ${currentGradient.start.s}%, ${currentGradient.start.l}%),
      hsl(${currentGradient.mid.h}, ${currentGradient.mid.s}%, ${currentGradient.mid.l}%),
      hsl(${currentGradient.end.h}, ${currentGradient.end.s}%, ${currentGradient.end.l}%))`,
  };

  return (
    <div className="absolute top-0 left-0 -z-10 w-screen h-screen opacity-30">
      <div
        className="absolute top-0 left-0 w-full h-full sm:blur-3xl blur-xl transition-all duration-10000 ease-in-out"
        style={gradientStyle}
      />
    </div>
  );
};