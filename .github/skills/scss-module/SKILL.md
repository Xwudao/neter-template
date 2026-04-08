---
name: scss-module
description: 'Use when creating or refactoring React component styles in neter-template with CSS Modules, especially when converting local .scss files to .module.scss with classes imports and clsx class composition.'
---

# SCSS Module

Use this skill when a neter-template React component or page should use local scoped styles instead of a plain `.scss` import.

## When to use

- Converting component-level or page-level `.scss` files to `.module.scss`
- Adding new scoped styles for a TSX component
- Refactoring class-heavy JSX into `classes + clsx` usage
- Styling local UI pieces without leaking selectors globally

## Core rules

1. Name the stylesheet with `.module.scss` and keep it next to the component or page.
2. Import the module as `classes`.
3. Import `clsx` and wrap every local class reference with `clsx(...)`, even when there is only one class.
4. Prefer kebab-case selectors in SCSS. Access them from TSX with the generated camelCase names (Vite CSS Modules `localsConvention: 'camelCase'` is configured).
5. Keep global `.scss` imports only for true app-wide styles such as `src/styles/index.scss`.

## Import pattern

```tsx
import clsx from 'clsx';
import classes from './example.module.scss';

export function Example() {
  return <div className={clsx(classes.root)}>...</div>;
}
```

## Class composition

```tsx
<button className={clsx(classes.action, isPrimary && classes.actionPrimary)}>
  Save
</button>
```

For icon utility classes or UnoCSS utilities, merge through `clsx`:

```tsx
<span className={clsx('i-mdi-check', classes.icon)} aria-hidden="true" />
```

## Styling guidance

- Start from tokens in `src/styles/_tokens.scss`
- Mixins from `@/styles/_mixins.scss` are auto-injected — use `card-surface`, `focus-ring`, `interactive-surface`, etc. without importing
- Prefer CSS variables over hard-coded values
- Keep surfaces compact and tool-like unless the task explicitly asks for a more decorative layout
- Keep selectors shallow and local; avoid deep nesting beyond 2 levels

## Allowed exceptions

Use `:global(...)` only when styling DOM owned by a third-party library.

```scss
.root {
  :global(.some-third-party-class) {
    min-height: 16rem;
  }
}
```

## Avoid

- Importing local `.module.scss` files for side effects only
- Mixing `styles`, `s`, or other import names when `classes` is the project convention
- Building class strings manually when `clsx` is clearer
- Leaving page/component-scoped styles in plain `.scss` files if they are not intended to be global
