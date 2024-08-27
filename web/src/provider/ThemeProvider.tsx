import { createContext, FC, PropsWithChildren } from 'react';
import useTheme from '@/hooks/useTheme';
interface ThemeContextProp {
  isDark: boolean;
  toggleTheme: () => void;
}

const ThemeContext = createContext<ThemeContextProp>(null!);

const ThemeProvider: FC<PropsWithChildren> = ({ children }) => {
  const { isDark, toggleTheme } = useTheme();

  return (
    <ThemeContext.Provider value={{ isDark, toggleTheme }}>
      {children}
    </ThemeContext.Provider>
  );
};

export default ThemeProvider;
