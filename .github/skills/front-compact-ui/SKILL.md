---
name: front-compact-ui
description: 'Use when designing or refactoring neter-template frontend pages that should feel compact, refined, content-dense, tool-like, doc-like, or suitable for search/detail/dashboard style content pages.'
---

# Front Compact UI

Use this skill when working on neter-template frontend pages that should feel small, precise, and efficient.

## Primary Goal

Build frontend pages that read like capable product surfaces, not promotional hero pages.

## Baseline

- Default body copy and labels should usually stay within 12px to 16px (`--font-size-xs` to `--font-size-md`).
- Spacing, card padding, and control sizes should stay compact before they become airy.
- Visual polish should come from hierarchy, border rhythm, muted surfaces, and restrained emphasis.
- Tool pages, dashboards, and detail pages should optimize scanning efficiency.
- Do not introduce poster-like hero sections, oversized type, or decorative empty space unless explicitly required.

## neter-template Conventions

- TanStack Router uses the Vite plugin workflow with file-based routes under `web/src/routes`.
- Keep `web/src/router.tsx` limited to `createRouter(...)` and the generated route tree import.
- Use `useAppConfig` from `web/src/store/useAppConfig.ts` for theme/accent state.
- Use `useTheme()` from `web/src/provider/ThemeProvider.tsx` for reading theme in components.

## Interaction Rules

- Prefer native-feeling button feedback and simple state changes over hover lift, glow sweeps, or dramatic shadows.
- Keep action groups dense but readable.
- Use light surface separation, subtle status pills, and compact metadata rows.

## Layout Rules

- Start from shared tokens in `web/src/styles/_tokens.scss`.
- Prefer compact cards, split panels, and two-column utility layouts over full-bleed compositions.
- Keep explanatory copy short; tool pages should show state, controls, and output early.

## Tool Page Layout Pattern

### Structure (top to bottom)

```
<main page>
  <div panel toolbar>       ← compact strip, ~40px tall
  <div editors-grid>        ← equal 2-column on desktop
  <div bottom-row>          ← options left + summary/log right
</main>
```

### Toolbar

- Left side: brand name / page title (icon + text, 600 weight) + status badge.
- Right side: all action buttons, grouped by function.
- Status badge uses a `data-*` attribute for state variants — colour in SCSS, not inline style.

### Token Safety

Only use CSS variables defined in `web/src/styles/_tokens.scss`. See the scss-theme-system skill for the full safe token list.

### Toolbar Wrapping (Mobile)

Always add `flex-wrap: wrap` to the outer toolbar and actions wrapper. Never use `flex-shrink: 0` on the actions wrapper — this prevents wrapping and causes overflow at 375px.

## Avoid

- Hero sections (gradient banner, kicker, large title, highlight pills)
- Preset/quick-fill drop-down UIs unless explicitly requested
- Tips or usage sidebars in the page body
- Marketing copy when project-status or workbench copy should be used
