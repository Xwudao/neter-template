import { create } from 'zustand'
import { persist } from 'zustand/middleware'

export type ThemeMode = 'light' | 'dark' | 'system'
export type ResolvedTheme = Exclude<ThemeMode, 'system'>
export type AccentPreset =
  | 'indigo'
  | 'emerald'
  | 'amber'
  | 'rose'
  | 'sky'
  | 'violet'
  | 'orange'

type AppConfigState = {
  theme: ThemeMode
  resolvedTheme: ResolvedTheme
  accent: AccentPreset
  setTheme: (theme: ThemeMode) => void
  cycleTheme: () => void
  syncSystemTheme: () => void
  setAccent: (accent: AccentPreset) => void
}

const THEME_SEQUENCE: ThemeMode[] = ['light', 'dark', 'system']

export const getSystemTheme = (): ResolvedTheme => {
  if (
    typeof window !== 'undefined' &&
    window.matchMedia?.('(prefers-color-scheme: dark)').matches
  ) {
    return 'dark'
  }
  return 'light'
}

export const resolveTheme = (theme: ThemeMode): ResolvedTheme =>
  theme === 'system' ? getSystemTheme() : theme

const validAccents: AccentPreset[] = [
  'indigo',
  'emerald',
  'amber',
  'rose',
  'sky',
  'violet',
  'orange',
]

const useAppConfig = create<AppConfigState>()(
  persist(
    (set) => ({
      theme: 'system',
      resolvedTheme: getSystemTheme(),
      accent: 'indigo',
      setTheme: (theme) => set({ theme, resolvedTheme: resolveTheme(theme) }),
      cycleTheme: () =>
        set((state) => {
          const currentIndex = THEME_SEQUENCE.indexOf(state.theme)
          const nextTheme = THEME_SEQUENCE[(currentIndex + 1) % THEME_SEQUENCE.length]
          return {
            theme: nextTheme,
            resolvedTheme: resolveTheme(nextTheme),
          }
        }),
      syncSystemTheme: () =>
        set((state) =>
          state.theme === 'system'
            ? { resolvedTheme: getSystemTheme() }
            : { resolvedTheme: resolveTheme(state.theme) },
        ),
      setAccent: (accent) => set({ accent }),
    }),
    {
      name: 'app-config',
      version: 3,
      partialize: (state) => ({
        theme: state.theme,
        accent: state.accent,
      }),
      merge: (persisted, current) => {
        const p = persisted as Partial<AppConfigState>
        return {
          ...current,
          ...p,
          theme: p.theme ?? current.theme,
          resolvedTheme: resolveTheme(p.theme ?? current.theme),
          accent: validAccents.includes(p.accent ?? 'indigo')
            ? (p.accent ?? 'indigo')
            : 'indigo',
        }
      },
      migrate: (persisted: unknown) => {
        const state = (persisted as Partial<AppConfigState>) ?? {}
        const theme = ['light', 'dark', 'system'].includes(state.theme ?? 'system')
          ? (state.theme ?? 'system')
          : 'system'
        const accent = validAccents.includes(state.accent ?? 'indigo')
          ? (state.accent ?? 'indigo')
          : 'indigo'
        return {
          ...state,
          theme,
          accent,
          resolvedTheme: resolveTheme(theme),
        }
      },
    },
  ),
)

export default useAppConfig
