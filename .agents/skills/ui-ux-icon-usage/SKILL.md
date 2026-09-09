---
name: ui-ux-icon-usage
description: 'Use when polishing neter-template frontend UI or UX and deciding where UnoCSS icons should be added or refined, especially for buttons, status labels, section headers, helper text, and empty states.'
---

# UI/UX Icon Usage

Use this skill when adjusting neter-template frontend interfaces and deciding whether icons should be introduced, changed, or removed.

## Goal

Add icons when they improve recognition speed, action clarity, or scanning rhythm. Do not add icons mechanically to every label.

## Project Rules

### Icon source

- Icons come from UnoCSS `@unocss/preset-icons` with `i-` prefix.
- Prefer staying in the Material Design Icons (MDI) family: `i-mdi-*`.
- Verify icon names exist before using them — do not guess.

### Common icons

- Run/play: `i-mdi-play-circle-outline`
- Stop: `i-mdi-stop-circle-outline`
- Copy: `i-mdi-content-copy`
- Delete/Clear: `i-mdi-delete-outline`
- Success: `i-mdi-check-circle-outline`
- Warning: `i-mdi-alert-outline`
- Info: `i-mdi-information-outline`
- Settings: `i-mdi-cog-outline`
- Dark/Light toggle: `i-mdi-weather-night` / `i-mdi-weather-sunny`

### Rendering pattern

```tsx
<span className="i-mdi-code-json" aria-hidden="true" />
```

When mixing with module styles:

```tsx
<span className={clsx('i-mdi-check-bold', classes.icon)} aria-hidden="true" />
```

## Default Decision Rules

### Good places to add icons

- Primary and secondary action buttons
- Section headers when they help distinguish blocks
- Status badges: ready, running, success, warning, failed
- Empty states and helper text

### Use sparingly

- Every field label in a long form
- Every row in repeated data
- Long paragraphs

## Style Guidance

- Prefer outline-style icons when a filled icon feels visually heavy.
- Keep icon size close to nearby text size (`1em` is the configured default).
- Use `aria-hidden="true"` for purely decorative icons.
- Merge utility classes with module classes using `clsx`.

## Review Checklist

- [ ] Icon improves scanability or action clarity
- [ ] Consistent icon family within the same UI area
- [ ] `aria-hidden="true"` on decorative icons
- [ ] Icon name verified to exist in MDI collection
- [ ] Classes composed with `clsx` when combining with module styles
