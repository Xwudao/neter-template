import { createContext, useContext, useEffect, type PropsWithChildren } from 'react'
import { Toaster } from 'react-hot-toast'
import useAppConfig, { type AccentPreset, type ThemeMode } from '@/store/useAppConfig'

type ThemeContextValue = {
  accent: AccentPreset
  theme: ThemeMode
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
}

const ThemeContext = createContext<ThemeContextValue | null>(null)

export function ThemeProvider({ children }: PropsWithChildren) {
  const accent = useAppConfig((state) => state.accent)
  const theme = useAppConfig((state) => state.theme)

  useEffect(() => {
    const root = document.documentElement
    root.dataset.theme = theme
    root.style.colorScheme = theme
  }, [theme])

  useEffect(() => {
    const root = document.documentElement
    const tokens = accentTokens[accent][theme]

    root.style.setProperty('--color-accent', tokens.accent)
    root.style.setProperty('--color-accent-hover', tokens.hover)
    root.style.setProperty('--color-accent-active', tokens.active)
    root.style.setProperty('--color-accent-soft', tokens.soft)
    root.style.setProperty('--color-accent-border', tokens.border)
    root.style.setProperty('--color-accent-contrast', tokens.contrast)
    root.style.setProperty('--accent', tokens.accent)
    root.style.setProperty('--accent-bg', tokens.soft)
    root.style.setProperty('--accent-border', tokens.border)
    root.style.setProperty('--color-signal', tokens.accent)
    root.style.setProperty('--color-signal-soft', tokens.soft)
  }, [accent, theme])

  return (
    <ThemeContext.Provider value={{ accent, theme }}>
      {children}
      <Toaster
        position="top-right"
        toastOptions={{
          duration: 3000,
          style: {
            background: 'var(--color-surface)',
            color: 'var(--text)',
            border: '1px solid var(--border)',
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
