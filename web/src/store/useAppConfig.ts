import { create } from 'zustand'
import { persist } from 'zustand/middleware'

export type ThemeMode = 'light' | 'dark'
export type AccentPreset = 'indigo' | 'emerald' | 'amber'

type AppConfigState = {
  theme: ThemeMode
  accent: AccentPreset
  setTheme: (theme: ThemeMode) => void
  toggleTheme: () => void
  setAccent: (accent: AccentPreset) => void
}

const getSystemTheme = (): ThemeMode => {
  if (
    typeof window !== 'undefined' &&
    window.matchMedia?.('(prefers-color-scheme: dark)').matches
  ) {
    return 'dark'
  }

  return 'light'
}

const useAppConfig = create<AppConfigState>()(
  persist(
    (set) => ({
      theme: getSystemTheme(),
      accent: 'indigo',
      setTheme: (theme) => set({ theme }),
      toggleTheme: () =>
        set((state) => ({
          theme: state.theme === 'dark' ? 'light' : 'dark',
        })),
      setAccent: (accent) => set({ accent }),
    }),
    {
      name: 'app-config',
    },
  ),
)

export default useAppConfig
