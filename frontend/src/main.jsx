import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.jsx'
const tg = window.Telegram?.WebApp;
tg?.ready(); 
tg?.expand(); 
tg?.ready();

document.documentElement.style.setProperty(
  "--tg-bg-color",
  tg?.colorScheme === "dark" ? "#1e1e1e" : "#ffffff"
);
document.documentElement.style.setProperty(
  "--tg-text-color",
  tg?.colorScheme === "dark" ? "#ffffff" : "#000000"
);


createRoot(document.getElementById('root')).render(
  <StrictMode>
    <App />
  </StrictMode>,
)
