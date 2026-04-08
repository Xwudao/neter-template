# neter-template Workspace Guidelines

## Repo Shape

- Go backend code lives under `cmd`, `internal`, and `pkg`.
- Frontend app lives under `web`.
- Frontend source code lives under `web/src`.

## Tech Stack (Frontend)

- React 19 + TanStack Router (file-based routes under `web/src/routes`)
- UnoCSS (`presetWind3`) + `@unocss/preset-icons` for utility classes and icons
- SCSS Modules for component styles (`*.module.scss`)
- Zustand for global state (`src/store/`)
- `ThemeProvider` + `useAppConfig` for theme/accent management
- oxlint + oxfmt for linting and formatting (no ESLint, no Prettier)
- tsgo (`@typescript/native-preview`) for type checking

## SCSS Theme System

- Tokens: `web/src/styles/_tokens.scss`
- Mixins: `web/src/styles/_mixins.scss` (auto-injected into every SCSS file)
- Global bootstrap: `web/src/styles/index.scss`
- Theme store: `web/src/store/useAppConfig.ts`
- DOM sync provider: `web/src/provider/ThemeProvider.tsx`

## Interaction Rule

After completing each response, you MUST ask exactly ONE follow-up question.

The purpose of this question is to check:

- whether the user is satisfied
- whether anything needs to be modified
- whether additional details or features are required

Requirements:

- Ask only ONE question (not multiple)
- Keep it concise and natural (like a real human)
- Do NOT sound repetitive or robotic

Examples:

- "Does this look good to you, or should I tweak anything?"
- "Want me to refine or add anything?"
- "Is this what you had in mind?"

Before finishing your response, verify:
→ A follow-up question is included

A response is NOT complete without this step.

<!-- intent-skills:start -->

# Skill mappings - when working in these areas, load the linked skill file into context.

skills:

- task: "adding SCSS styles, design tokens, dark mode, CSS variables, spacing, typography, radius, hover states, or ThemeProvider/Zustand-linked theme changes"
  load: ".github/skills/scss-theme-system/SKILL.md"
- task: "creating or refactoring React component styles with CSS Modules, converting .scss files to .module.scss, or using clsx class composition"
  load: ".github/skills/scss-module/SKILL.md"
- task: "working on TanStack Router setup, route trees, generated route files, or route file naming"
  load: "web/node_modules/@tanstack/router-core/skills/router-core/SKILL.md"
- task: "adding or changing route loaders, pending states, cache behavior, or deferred route data"
  load: "web/node_modules/@tanstack/router-core/skills/router-core/data-loading/SKILL.md"
- task: "changing links, navigate flows, preloading on intent, or scroll restoration"
  load: "web/node_modules/@tanstack/router-core/skills/router-core/navigation/SKILL.md"
- task: "working with dynamic route segments, path params, or search params in URLs"
  load: "web/node_modules/@tanstack/router-core/skills/router-core/path-params/SKILL.md"
- task: "changing the Vite TanStack Router plugin, generated routes, or automatic route code splitting"
  load: "web/node_modules/@tanstack/router-plugin/skills/router-plugin/SKILL.md"
- task: "designing or refactoring compact, tool-like, content-dense, or dashboard-style frontend pages"
  load: ".github/skills/front-compact-ui/SKILL.md"
- task: "adding or refactoring a neter-template page that has substantial state, handlers, or side effects, or extracting logic from a TSX file into a hook file"
  load: ".github/skills/page-hook-split/SKILL.md"
- task: "deciding where icons should be added or refined, polishing UI with UnoCSS icons for buttons, status labels, section headers, or empty states"
  load: ".github/skills/ui-ux-icon-usage/SKILL.md"
<!-- intent-skills:end -->
