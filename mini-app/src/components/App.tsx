import { useLaunchParams, miniApp, useSignal} from '@telegram-apps/sdk-react';
import { AppRoot } from '@telegram-apps/telegram-ui';
import { Navigate, Route, Routes, HashRouter } from 'react-router-dom';
import { routes } from '@/navigation/routes.tsx';
import { useEffect } from 'react';


export function App() {
  const lp = useLaunchParams();
  const isDark = useSignal(miniApp.isDark);

  useEffect(() => {
    console.log("initDataRaw from Telegram:", lp.initDataRaw);
}, [lp.initDataRaw])

  // Отправка данных при инициализации
  useEffect(() => {
    
    const sendLaunchData = async () => {
      try {
        const response = await fetch('/api/v1/auth/telegram', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            init_data: lp.initDataRaw, // Используем сырые данные из launch params
            platform: lp.platform,
            theme_params: lp.themeParams
          }),
        });

        
        if (!response.ok) throw new Error('Auth failed');
        
        const data = await response.json();
        console.log('User authenticated:', data);
        
       
        miniApp.ready(); 
        
      } catch (error) {
        console.error('Auth error:', error);
       
      }
    };

    sendLaunchData();
  }, [lp.initDataRaw, lp.platform, lp.themeParams]); // Зависимости от параметров запуска

  return (
    <AppRoot
      appearance={isDark ? 'dark' : 'light'}
      platform={['macos', 'ios'].includes(lp.platform) ? 'ios' : 'base'}
    >
      <HashRouter>
        <Routes>
          {routes.map((route) => <Route key={route.path} {...route} />)}
          <Route path="*" element={<Navigate to="/"/>}/>
        </Routes>
      </HashRouter>
    </AppRoot>
  );
}