module.exports = {
  darkMode: "class",
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
    "./public/index.html"
  ],
  theme: {
    extend: {
      colors: {
        'tg-bg': 'var(--tg-theme-bg-color)',
        'tg-secondary-bg': 'var(--tg-theme-secondary-bg-color)',
        'tg-text': 'var(--tg-theme-text-color)',
        'tg-hint': 'var(--tg-theme-hint-color)',
        'tg-link': 'var(--tg-theme-link-color)',
      },      
      letterSpacing: {
        'tg-title': '0.5px',
      },
      fontFamily: {
        'tg-default': 'var(--tg-theme-font-family)',
      },
      borderRadius: {
        'tg-rounded': 'var(--tg-theme-border-radius, 0.25rem)',
      },
    },
  },
  plugins: [],
};