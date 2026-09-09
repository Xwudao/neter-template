---
name: scss-theme-system
description: 'Use when adding SCSS styles, design tokens, dark mode, CSS variables, spacing, typography, radius, hover states, or ThemeProvider/Zustand-linked theme changes in neter-template.'
---

# SCSS Theme System

Use this skill when working on the visual token system of neter-template.

## Source Of Truth

- Theme tokens live in `web/src/styles/_tokens.scss`
- Shared Sass helpers live in `web/src/styles/_mixins.scss` (auto-injected into every SCSS file via Vite `additionalData`)
- Global theme bootstrap lives in `web/src/styles/index.scss`
- Runtime theme state lives in `web/src/store/useAppConfig.ts`
- Runtime DOM sync lives in `web/src/provider/ThemeProvider.tsx`

## Rules

1. Prefer CSS variables over hard-coded values.
2. Use the existing token scales before inventing new spacing, radius, font, or color values.
3. Dark is the **default** theme. Light mode overrides live under `:root[data-theme='light']`.
4. If accent behavior changes at runtime, update `ThemeProvider.tsx` instead of scattering overrides across components.
5. For new SCSS files, rely on the shared helpers — they are auto-injected by Vite.
6. Preserve compatibility aliases like `--text`, `--bg`, and `--accent`.
7. Do not add `background-color` or `color` transitions to page-level surfaces or shared controls unless the motion is explicitly required.

## Token Quick Reference

### Typography
`--font-sans` · `--font-heading` · `--font-mono`
`--font-weight-{regular|medium|semibold|bold}`
`--font-size-{2xs|xs|sm|md|lg|xl|2xl|3xl|4xl|display}`
`--line-height-{tight|snug|normal|relaxed}`
`--letter-spacing-{tight|normal|wide}`

### Spacing
`--space-{0|0-5|1|1-5|2|2-5|3|4|5|6|8|10|12|16}`
Control padding: `--padding-control-{xs|sm|md|lg}`
Page/card: `--padding-card`, `--padding-page`
Gaps: `--stack-gap-{sm|md|lg}`
Margins: `--margin-inline-section`, `--margin-block-section`

### Radius
`--radius-xs` · `--radius-sm` · `--radius-md` · `--radius-lg` · `--radius-xl` · `--radius-2xl` · `--radius-pill`

### Colors
Surface: `--color-canvas`, `--color-surface`, `--color-surface-muted`, `--color-surface-elevated`
Text: `--color-text`, `--color-text-muted`, `--color-text-strong`
Border: `--color-border-soft`, `--color-border-strong`
Accent: `--color-accent`, `--color-accent-hover`, `--color-accent-active`, `--color-accent-soft`, `--color-accent-border`, `--color-accent-contrast`
Semantic: `--color-{success|warning|danger|info}`, `--color-signal`, `--color-signal-soft`

### Motion
`--duration-{fast|normal|slow}` · `--ease-standard`

### Elevation
`--shadow-xs` · `--shadow-sm` · `--shadow-md` · `--shadow-lg`

### Layout/Z-index
`--container-page`, `--header-h`
`--z-{raised|dropdown|header|modal|top}`

### Semantic Aliases
`--bg`, `--text`, `--text-h`, `--border`, `--code-bg`
`--accent`, `--accent-bg`, `--accent-border`
`--shadow`, `--sans`, `--heading`, `--mono`

## Token Safety — Tokens That Do NOT Exist

- `--color-text-primary` → use `--color-text-strong`
- `--color-text-secondary` → use `--color-text-muted`
- `--color-surface-raised` → use `--color-surface`
- `--color-border` → use `--color-border-strong` (hover) or `--color-border-soft` (default)
- `--color-error` → use `--color-danger`
- `--color-success-subtle` → use `color-mix(in srgb, var(--color-success) 12%, transparent)`

## Accent Presets (ThemeProvider)

Three runtime-switchable presets — `indigo` (default), `emerald`, `amber`.
Switch via `useAppConfig((s) => s.setAccent)`. `ThemeProvider.tsx` writes all `--color-accent-*` tokens to `document.documentElement`.

## Mixins Available (auto-injected)

```scss
// Token accessors
color-token($name)    → var(--color-#{$name})
space($step)          → var(--space-#{$step})
radius($size)         → var(--radius-#{$size})
font-size($size)      → var(--font-size-#{$size})

// Interaction & surfaces
@mixin focus-ring($color, $offset)
@mixin interactive-surface
@mixin card-surface
@mixin elevated-surface

// Text
@mixin text-truncate
@mixin text-clamp($lines)

// Layout
@mixin flex-center
@mixin flex-between

// Responsive breakpoints
@mixin sm   // ≥ 640px
@mixin md   // ≥ 768px
@mixin lg   // ≥ 1024px
@mixin xl   // ≥ 1280px
```

## Workflow

1. Start from tokens in `web/src/styles/_tokens.scss`.
2. Reuse helpers from `web/src/styles/_mixins.scss` — auto-injected, no import needed.
3. For theme-linked behavior, read from `useAppConfig` or `useTheme()` instead of duplicating state.
4. Validate with `pnpm build` and `pnpm check` after theme changes.

## Avoid

- Hard-coded hex values inside component SCSS unless adding a new shared token at the same time
- Separate dark-mode classes when a token override is enough
- Replacing the ThemeProvider/store contract with local component state for global theme concerns
- Adding generic background/color transitions to body, panels, or form controls when the page does not need that motion


## Source Of Truth

- Theme tokens live in `src/styles/_tokens.scss`
- Shared Sass helpers live in `src/styles/_mixins.scss`
- Global theme bootstrap lives in `src/styles/index.scss`
- Runtime theme state lives in `src/store/useAppConfig.ts`
- Runtime DOM sync lives in `src/provider/ThemeProvider.tsx`

## Rules

1. Prefer CSS variables over hard-coded values.
2. Use the existing token scales before inventing new spacing, radius, font, or color values.
3. Keep light and dark mode in the same token system by overriding variables under `:root[data-theme='light']` (dark is the default).
4. If accent behavior changes at runtime, update `ThemeProvider.tsx` instead of scattering overrides across components.
5. For new SCSS files, rely on the shared helpers injected by Vite from `@/styles/mixins` (auto-injected via `additionalData`).
6. Preserve compatibility aliases like `--text`, `--bg`, and `--accent` unless the old consumers are migrated together.
7. Do not add `background-color` or `color` transitions to page-level surfaces or shared controls unless the motion is explicitly required; theme switches will animate those token changes and cause visible flicker.

## Token Families

### Typography

- Font families: `--font-sans`, `--font-heading`, `--font-mono`
- Weights: `--font-weight-{regular|medium|semibold|bold}`
- Sizes: `--font-size-{2xs|xs|sm|md|lg|xl|2xl|3xl|4xl|display}`
- Line heights: `--line-height-{tight|snug|normal|relaxed}`
- Letter spacing: `--letter-spacing-{tight|normal|wide}`

### Spacing

- `--space-{0|0-5|1|1-5|2|2-5|3|4|5|6|8|10|12|16}`
- Control padding presets: `--padding-control-{xs|sm|md|lg}`
- Page/card padding: `--padding-card`, `--padding-page`, `--padding-section`
- Stack gaps: `--stack-gap-{xs|sm|md|lg}`
- Section margins: `--margin-inline-section`, `--margin-block-section`

### Radius

- `--radius-{none|xs|sm|md|lg|xl|2xl|pill}`

### Colors

- Canvas/surface/muted/elevated
- Text: `--color-text`, `--color-text-muted`, `--color-text-strong`
- Border: `--color-border-soft`, `--color-border-strong`
- Accent: `--color-accent`, `--color-accent-hover`, `--color-accent-active`, `--color-accent-soft`, `--color-accent-border`, `--color-accent-contrast`
- Hover overlays: `--color-hover-overlay`, `--color-active-overlay`
- Semantic: `--color-{success|warning|danger|info}` + `--color-{success|warning|danger|info}-soft`

### Motion

- `--duration-{fast|normal|slow}`, `--ease-standard`

### Elevation

- `--shadow-{xs|sm|md|lg}`

### Layout

- `--container-page`, `--header-h`, `--sidebar-w`
- Z-index: `--z-{raised|dropdown|header|modal|top}`

### Semantic Aliases

- `--bg`, `--text`, `--text-h`, `--border`, `--code-bg`
- `--accent`, `--accent-bg`, `--accent-border`
- `--shadow`, `--sans`, `--heading`, `--mono`

## Accent Presets (ThemeProvider)

Three runtime-switchable accent presets in `src/store/useAppConfig.ts`:

- `indigo` — Default. Blue-violet.
- `emerald` — Green teal.
- `amber` — Warm orange-gold.

`ThemeProvider.tsx` writes the accent tokens to `document.documentElement` on accent/theme change.

## Mixins Available (auto-injected)

```scss
// Token accessors
color-token($name)   → var(--color-#{$name})
space($step)         → var(--space-#{$step})
radius($size)        → var(--radius-#{$size})
font-size($size)     → var(--font-size-#{$size})

// Interaction
@mixin focus-ring($color, $offset)
@mixin interactive-surface

// Surfaces
@mixin card-surface
@mixin elevated-surface

// Text
@mixin text-truncate
@mixin text-clamp($lines)

// Layout
@mixin flex-center
@mixin flex-between

// Breakpoints
@mixin sm / md / lg / xl
```

## Workflow

1. Start from tokens in `src/styles/_tokens.scss`.
2. Reuse helpers from `src/styles/_mixins.scss` — they are auto-injected into every SCSS file.
3. If a component needs theme-aware behavior, read from `useAppConfigStore` or `useAppConfig` instead of duplicating state.
4. Validate with `pnpm build` and `pnpm check` after theme changes.

## Avoid

- Hard-coded hex values inside component SCSS unless adding a new shared token at the same time
- Separate dark-mode classes when a token override is enough
- Replacing the provider/store contract with local component state for global theme concerns
- Adding generic background/color transitions to body, panels, buttons, or form controls when the page does not need that motion
