import Cookies from 'js-cookie';
import { useEffect, useState } from 'react';
const body = document.body;

// 创建一个全局事件来同步主题状态
const THEME_CHANGE_EVENT = 'theme-change';

const useTheme = () => {
  const [isDark, setIsDark] = useState(Cookies.get('is_dark') === 'true');

  const toggleTheme = () => {
    const newTheme = !isDark;
    setIsDark(newTheme);
    Cookies.set('is_dark', String(newTheme));
    // 派发自定义事件通知其他组件
    window.dispatchEvent(new CustomEvent(THEME_CHANGE_EVENT, { detail: newTheme }));
  };

  useEffect(() => {
    // 监听其他组件的主题变化
    const handleThemeChange = (event: CustomEvent) => {
      setIsDark(event.detail);
    };

    window.addEventListener(THEME_CHANGE_EVENT, handleThemeChange as EventListener);

    return () => {
      window.removeEventListener(THEME_CHANGE_EVENT, handleThemeChange as EventListener);
    };
  }, []);

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
