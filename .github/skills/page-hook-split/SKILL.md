---
name: page-hook-split
description: 'Use when a neter-template page component has substantial state, handlers, or side effects that should be extracted into a dedicated hook file so the TSX file is reduced to pure rendering.'
---

# Page Hook Split

Use this skill whenever a page or complex component holds enough logic that the TSX file becomes hard to scan.
The goal is a clean separation: **logic lives in the hook, TSX only renders**.

## When to apply

- A page component has multiple `useState`, `useMemo`, `useEffect`, or `useRef` calls alongside non-trivial handler functions.
- The TSX file is long enough that the JSX is buried beneath type definitions, helper functions, and data constants.
- The same logic or data shape needs to be reused across more than one component.
- You are adding new logic to an existing page that already has a `.hook.ts` file.

## File naming convention

| File | Purpose |
|---|---|
| `some-page.tsx` | JSX rendering only — imports from the hook |
| `hooks/some-page.hook.ts` | All state, derived values, handlers, side effects, constants, and types |
| `some-page.module.scss` | Scoped styles (unchanged) |

Place the hook file under a `hooks/` subdirectory inside the same `pages/` folder. This keeps hook files separated from TSX and SCSS files without moving them out of the pages tree.

## What belongs in the hook file

- All `useState`, `useMemo`, `useCallback`, `useEffect`, `useRef` declarations
- Handler functions (`handleXxx`, `applyXxx`, `onXxx`)
- Derived values computed from state
- localStorage persistence logic (read on init, write on change via `useEffect`)
- Static data constants used only by this page
- Shared pure helper functions used by the above
- All type and interface definitions for the page's data shapes
- The named `useXxx()` hook that bundles everything and returns a flat object

## What stays in the TSX file

- The `export default function` component — JSX and nothing else
- A single destructure of the hook's return value at the top of the component body
- Imports for the hook, SCSS module, `clsx`, and any UI components

## Hook return shape

Return a **flat object**. Avoid nested objects so the TSX destructure stays readable.

```ts
return {
  // state
  input, setInput,
  output,
  activeTab, setActiveTab,
  // derived
  inputStats,
  // refs
  outputRef,
  // handlers
  handleSubmit,
  handleCopy,
  handleClear,
}
```

## localStorage persistence

When a user choice should survive a page refresh, persist it in localStorage.

- Use a namespaced key: `neter:<page-name>:<field>`, e.g. `neter:dashboard:tab`.
- Initialize state from a **lazy initializer function** so localStorage is only read once.
- Validate the stored value before accepting it; fall back to the default if missing or unknown.
- Write back on every change via a dedicated `useEffect`.

```ts
const STORAGE_KEY = 'neter:some-page:tab'

function readStoredTab(): TabKey {
  const stored = localStorage.getItem(STORAGE_KEY)
  const valid: TabKey[] = ['a', 'b', 'c']
  return valid.includes(stored as TabKey) ? (stored as TabKey) : 'a'
}

const [activeTab, setActiveTab] = useState<TabKey>(readStoredTab) // lazy init
```

## Minimal TSX skeleton after extraction

```tsx
import clsx from 'clsx'
import { useMyPage } from './hooks/my-page.hook'
import classes from './my-page.module.scss'

export default function MyPage() {
  const {
    input, setInput,
    activeTab, setActiveTab,
  } = useMyPage()

  return (
    <div className={classes.page}>
      {/* pure JSX — no logic here */}
    </div>
  )
}
```
