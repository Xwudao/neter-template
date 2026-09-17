import clsx from 'clsx'

import useAppConfig, {
  type AccentPreset,
  type SurfaceStyle,
  type ThemeMode,
} from '@/store/useAppConfig'

import classes from './theme-picker.module.scss'

const accentLabels: Record<AccentPreset, string> = {
  indigo: 'Indigo',
  emerald: 'Emerald',
  amber: 'Amber',
  rose: 'Rose',
  sky: 'Sky',
  violet: 'Violet',
  orange: 'Orange',
}

const accentColors: Record<AccentPreset, string> = {
  indigo: '#818cf8',
  emerald: '#34d399',
  amber: '#fbbf24',
  rose: '#fb7185',
  sky: '#38bdf8',
  violet: '#a78bfa',
  orange: '#fb923c',
}

const surfaceOptions: { key: SurfaceStyle; label: string; icon: string }[] = [
  { key: 'solid', label: 'Solid', icon: 'i-mdi-square-rounded' },
  { key: 'glass', label: 'Glass', icon: 'i-mdi-blur' },
]

export function ThemePicker() {
  const theme = useAppConfig((s) => s.theme)
  const accent = useAppConfig((s) => s.accent)
  const cycleTheme = useAppConfig((s) => s.cycleTheme)
  const setAccent = useAppConfig((s) => s.setAccent)
  const surfaceStyle = useAppConfig((s) => s.surfaceStyle)
  const setSurfaceStyle = useAppConfig((s) => s.setSurfaceStyle)

  const themeIcon = (t: ThemeMode): string => {
    switch (t) {
      case 'light':
        return 'i-mdi-weather-sunny'
      case 'dark':
        return 'i-mdi-weather-night'
      case 'system':
        return 'i-mdi-theme-light-dark'
    }
  }

  return (
    <div className={clsx(classes.container)}>
      <button
        className={clsx(classes.toggle)}
        onClick={cycleTheme}
        title={`Theme: ${theme}`}
        type="button"
      >
        <span className={clsx(themeIcon(theme), classes.toggleIcon)} aria-hidden="true" />
        <span className="text-xs capitalize">{theme}</span>
      </button>

      <div className={clsx(classes.accentList)}>
        {(Object.keys(accentLabels) as AccentPreset[]).map((key) => (
          <button
            key={key}
            className={clsx(classes.accentBtn, accent === key && classes.accentBtnActive)}
            style={{ backgroundColor: accentColors[key] }}
            onClick={() => setAccent(key)}
            title={accentLabels[key]}
            type="button"
            aria-label={accentLabels[key]}
          />
        ))}
      </div>

      <div className={clsx(classes.surfaceList)}>
        {surfaceOptions.map(({ key, label, icon }) => (
          <button
            key={key}
            className={clsx(
              classes.surfaceBtn,
              surfaceStyle === key && classes.surfaceBtnActive,
            )}
            onClick={() => setSurfaceStyle(key)}
            title={label}
            type="button"
            aria-pressed={surfaceStyle === key}
          >
            <span className={clsx(icon, classes.surfaceIcon)} aria-hidden="true" />
            <span>{label}</span>
          </button>
        ))}
      </div>
    </div>
  )
}
