import { createContext, useContext, useEffect, type PropsWithChildren } from 'react'
import { Toaster } from 'react-hot-toast'

import useAppConfig, {
  type AccentPreset,
  type ResolvedTheme,
  type ThemeMode,
} from '@/store/useAppConfig'

type ThemeContextValue = {
  accent: AccentPreset
  theme: ThemeMode
  resolvedTheme: ResolvedTheme
}

type AccentTokenSet = {
  accent: string
  hover: string
  active: string
  soft: string
  border: string
  contrast: string
}

const accentTokens: Record<
  AccentPreset,
  {
    dark: AccentTokenSet
    light: AccentTokenSet
  }
> = {
  indigo: {
    dark: {
      accent: '#818cf8',
      hover: '#6366f1',
      active: '#4f46e5',
      soft: 'rgba(129, 140, 248, 0.14)',
      border: 'rgba(129, 140, 248, 0.42)',
      contrast: '#08070f',
    },
    light: {
      accent: '#4f46e5',
      hover: '#4338ca',
      active: '#3730a3',
      soft: 'rgba(79, 70, 229, 0.1)',
      border: 'rgba(79, 70, 229, 0.24)',
      contrast: '#ffffff',
    },
  },
  emerald: {
    dark: {
      accent: '#34d399',
      hover: '#10b981',
      active: '#059669',
      soft: 'rgba(52, 211, 153, 0.14)',
      border: 'rgba(52, 211, 153, 0.42)',
      contrast: '#022c22',
    },
    light: {
      accent: '#059669',
      hover: '#047857',
      active: '#065f46',
      soft: 'rgba(5, 150, 105, 0.1)',
      border: 'rgba(5, 150, 105, 0.24)',
      contrast: '#ffffff',
    },
  },
  amber: {
    dark: {
      accent: '#fbbf24',
      hover: '#f59e0b',
      active: '#d97706',
      soft: 'rgba(251, 191, 36, 0.14)',
      border: 'rgba(251, 191, 36, 0.4)',
      contrast: '#241302',
    },
    light: {
      accent: '#d97706',
      hover: '#b45309',
      active: '#92400e',
      soft: 'rgba(217, 119, 6, 0.1)',
      border: 'rgba(217, 119, 6, 0.24)',
      contrast: '#ffffff',
    },
  },
  rose: {
    dark: {
      accent: '#fb7185',
      hover: '#f43f5e',
      active: '#e11d48',
      soft: 'rgba(251, 113, 133, 0.14)',
      border: 'rgba(251, 113, 133, 0.42)',
      contrast: '#fff1f2',
    },
    light: {
      accent: '#e11d48',
      hover: '#be123c',
      active: '#9f1239',
      soft: 'rgba(225, 29, 72, 0.1)',
      border: 'rgba(225, 29, 72, 0.24)',
      contrast: '#ffffff',
    },
  },
  sky: {
    dark: {
      accent: '#38bdf8',
      hover: '#0ea5e9',
      active: '#0284c7',
      soft: 'rgba(56, 189, 248, 0.14)',
      border: 'rgba(56, 189, 248, 0.42)',
      contrast: '#082f49',
    },
    light: {
      accent: '#0284c7',
      hover: '#0369a1',
      active: '#075985',
      soft: 'rgba(2, 132, 199, 0.1)',
      border: 'rgba(2, 132, 199, 0.24)',
      contrast: '#ffffff',
    },
  },
  violet: {
    dark: {
      accent: '#a78bfa',
      hover: '#8b5cf6',
      active: '#7c3aed',
      soft: 'rgba(167, 139, 250, 0.14)',
      border: 'rgba(167, 139, 250, 0.42)',
      contrast: '#2e1065',
    },
    light: {
      accent: '#7c3aed',
      hover: '#6d28d9',
      active: '#5b21b6',
      soft: 'rgba(124, 58, 237, 0.1)',
      border: 'rgba(124, 58, 237, 0.24)',
      contrast: '#ffffff',
    },
  },
  orange: {
    dark: {
      accent: '#fb923c',
      hover: '#f97316',
      active: '#ea580c',
      soft: 'rgba(251, 146, 60, 0.14)',
      border: 'rgba(251, 146, 60, 0.42)',
      contrast: '#431407',
    },
    light: {
      accent: '#ea580c',
      hover: '#c2410c',
      active: '#9a3412',
      soft: 'rgba(234, 88, 12, 0.1)',
      border: 'rgba(234, 88, 12, 0.24)',
      contrast: '#ffffff',
    },
  },
}

const ThemeContext = createContext<ThemeContextValue | null>(null)

export function ThemeProvider({ children }: PropsWithChildren) {
  const accent = useAppConfig((state) => state.accent)
  const theme = useAppConfig((state) => state.theme)
  const resolvedTheme = useAppConfig((state) => state.resolvedTheme)
  const syncSystemTheme = useAppConfig((state) => state.syncSystemTheme)

  // Listen for system theme changes
  useEffect(() => {
    const mediaQuery = window.matchMedia?.('(prefers-color-scheme: dark)')
    if (!mediaQuery) return undefined

    syncSystemTheme()

    const handleChange = () => {
      syncSystemTheme()
    }

    mediaQuery.addEventListener('change', handleChange)

    return () => {
      mediaQuery.removeEventListener('change', handleChange)
    }
  }, [syncSystemTheme])

  // Sync theme to DOM
  useEffect(() => {
    const root = document.documentElement
    root.dataset.theme = resolvedTheme
    root.dataset.themeMode = theme
    root.style.colorScheme = resolvedTheme
  }, [resolvedTheme, theme])

  // Sync accent colors
  useEffect(() => {
    const root = document.documentElement
    const tokens = (accentTokens[accent] ?? accentTokens['indigo'])[resolvedTheme]

    root.style.setProperty('--color-accent', tokens.accent)
    root.style.setProperty('--color-accent-hover', tokens.hover)
    root.style.setProperty('--color-accent-active', tokens.active)
    root.style.setProperty('--color-accent-soft', tokens.soft)
    root.style.setProperty('--color-accent-border', tokens.border)
    root.style.setProperty('--color-accent-contrast', tokens.contrast)
    root.style.setProperty('--color-signal', tokens.accent)
    root.style.setProperty('--color-signal-soft', tokens.soft)
  }, [accent, resolvedTheme])

  return (
    <ThemeContext.Provider value={{ accent, theme, resolvedTheme }}>
      {children}
      <Toaster
        position="top-center"
        containerStyle={{ zIndex: 'var(--z-toast)' }}
        toastOptions={{
          duration: 3000,
          style: {
            background: 'var(--color-surface)',
            color: 'var(--color-text)',
            border: '1px solid var(--color-border-soft)',
            boxShadow: 'var(--shadow-md)',
            fontSize: 'var(--font-size-sm)',
          },
        }}
      />
    </ThemeContext.Provider>
  )
}

export function useTheme(): ThemeContextValue {
  const ctx = useContext(ThemeContext)
  if (!ctx) throw new Error('useTheme must be used inside <ThemeProvider>')
  return ctx
}
