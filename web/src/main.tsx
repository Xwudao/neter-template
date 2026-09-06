import { createBrowserHistory, RouterProvider } from '@tanstack/react-router'
import { StrictMode } from 'react'
import { hydrateRoot } from 'react-dom/client'

import { ThemeProvider } from '@/provider/ThemeProvider'
import { createAppRouter } from '@/router'
import { readSSRPageData, SSRDataProvider } from '@/ssr-data'

import '@unocss/reset/tailwind.css'
import 'virtual:uno.css'
import '@/styles/index.scss'

const router = createAppRouter(createBrowserHistory())
const ssrPageData = readSSRPageData()

hydrateRoot(document.getElementById('root')!,
  <StrictMode>
    <SSRDataProvider value={ssrPageData}>
      <ThemeProvider>
        <RouterProvider router={router} />
      </ThemeProvider>
    </SSRDataProvider>
  </StrictMode>, 
)
