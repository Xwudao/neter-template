import '@/assets/styles/app-imports.scss';
import { createRoot } from 'react-dom/client';
import 'uno.css';
import App from './App.tsx';

// initialization
// initVChartSemiTheme();


createRoot(document.getElementById('root')!).render(
  // <StrictMode>
  <App />,
  // </StrictMode>,
);
