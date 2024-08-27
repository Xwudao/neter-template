import Cookies from 'js-cookie';
import { useEffect } from 'react';
const body = document.body;
const useTheme = () => {
  const [isDark, setIsDark] = useState(Cookies.get('is_dark') === 'true');
  const toggleTheme = () => {
    setIsDark(!isDark);
    Cookies.set('is_dark', String(!isDark));
  };

  useEffect(() => {
    if (isDark) {
      body.setAttribute('theme-mode', 'dark');
      document.documentElement.classList.add('dark');
    } else {
      body.removeAttribute('theme-mode');
      document.documentElement.classList.remove('dark');
    }
  }, [isDark]);

  return {
    isDark,
    toggleTheme,
  };
};

export default useTheme;
