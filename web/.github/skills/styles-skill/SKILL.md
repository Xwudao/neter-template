---
name: styles-skill
description: 'Use when creating or modifying styles, theme, design tokens, CSS variables, SCSS, and visual components in neter-template.'
---

# Styles System — neter-template

This skill documents the complete style infrastructure for neter-template.

## Design Tokens

All design tokens are defined as CSS variables in `web/src/styles/_tokens.scss`.

### Token Reference

#### Typography

- `--font-{sans|heading|mono}` — Font families
- `--font-weight-{regular|medium|semibold|bold}` — Weights
- `--font-size-{2xs|xs|sm|md|lg|xl|2xl|3xl|4xl|display}` — Sizes
- `--line-height-{tight|snug|normal|relaxed}` — Line heights
- `--letter-spacing-{tight|normal|wide}` — Letter spacing

#### Spacing

- `--space-{0|0-5|1|1-5|2|2-5|3|4|5|6|7|8|9|10|12|16}` — Spacing scale
- `--padding-control-{xs|sm|md|lg}` — Control padding presets
- `--padding-{card|page}` — Card/page padding
- `--stack-gap-{xs|sm|md|lg}` — Stack gaps
- `--radius-{none|xs|sm|md|lg|xl|2xl|pill}` — Border radii

#### Colors (Dark Default)

- Surface: `--color-canvas`, `--color-surface`, `--color-surface-muted`, `--color-surface-elevated`
- Text: `--color-text`, `--color-text-muted`, `--color-text-strong`
- Border: `--color-border-soft`, `--color-border-strong`
- Accent: `--color-accent`, `--color-accent-hover`, `--color-accent-active`, `--color-accent-soft`, `--color-accent-border`, `--color-accent-contrast`
- Semantic: `--color-{success|warning|danger|info}`, `--color-signal`
- Soft variants: `--color-{success|warning|danger|info|signal}-soft`

Light mode overrides live under `:root[data-theme='light']`.

#### Elevation

- `--shadow-{xs|sm|md|lg}`

#### Motion

- `--duration-{fast|normal|slow}`, `--ease-standard`

#### Z-index

- `--z-{raised|dropdown|header|modal|top|toast}`

## SCSS Mixins

Located in `web/src/styles/_mixins.scss` — auto-injected into every SCSS file via Vite.

### Functions

- `color-token($name)` → `var(--color-#{$name})`
- `space($step)` → `var(--space-#{$step})`
- `radius($size)` → `var(--radius-#{$size})`
- `font-size($size)` → `var(--font-size-#{$size})`

### Mixins

- `focus-ring($color, $offset)` — Focus ring styles
- `interactive-surface` — Transition defaults
- `card-surface` — Card with border, radius, shadow
- `elevated-surface` — Elevated surface
- `text-truncate` — Single line truncation
- `text-clamp($lines)` — Multi-line clamp
- `flex-center` / `flex-between` — Flex helpers
- `sm`, `md`, `lg`, `xl` — Responsive breakpoints

## UnoCSS

Configuration in `web/uno.config.ts`.

- Presets: `presetWind3` (without preflight) + `presetIcons`
- Icon prefix: `i-` (e.g., `i-mdi-close`)
- Icon scale: 1em

## Component Styling Patterns

### Inline styles (preferred for simple components)

Use CSS variable references for theming:

```tsx
style={{
  background: 'var(--color-surface)',
  color: 'var(--color-text)',
  border: '1px solid var(--color-border-soft)',
}}
```

### SCSS Modules (for complex components)

- File: `*.module.scss` next to component
- Import: `import classes from './example.module.scss'`
- Use `clsx()` for class composition

### UnoCSS utilities (for layout/quick styling)

- Use `className` with UnoCSS utility classes
- Combine with `clsx()` when dynamic

## Theme System

State managed by Zustand store `web/src/store/useAppConfig.ts`.
DOM synced by `web/src/provider/ThemeProvider.tsx`.

### Theme modes

- `light`, `dark`, `system` (auto-detect prefers-color-scheme)
- Cycled via `useAppConfig((s) => s.cycleTheme)`

### Accent colors

7 presets: `indigo` (default), `emerald`, `amber`, `rose`, `sky`, `violet`, `orange`
Set via `useAppConfig((s) => s.setAccent('amber'))`

## Rules

1. Always use CSS variables instead of hard-coded colors, spacing, or radii.
2. Dark is the default theme. Light overrides go under `:root[data-theme='light']`.
3. For runtime accent changes, update `ThemeProvider.tsx` not individual components.
4. Don't add background/color transitions to body or panels.
5. Keep SCSS shallow — max 2 levels of nesting.
6. Reuse existing tokens before adding new ones.
