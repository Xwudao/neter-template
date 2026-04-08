import { StrictMode } from 'react'
import { RouterProvider } from '@tanstack/react-router'
import { createRoot } from 'react-dom/client'
import { ThemeProvider } from '@/provider/ThemeProvider'
import { router } from '@/router'
import 'virtual:uno.css'
import '@/styles/index.scss'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ThemeProvider>
      <RouterProvider router={router} />
    </ThemeProvider>
  </StrictMode>,
)
