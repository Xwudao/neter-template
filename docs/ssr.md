# Hybrid SSR

Public TanStack Router routes are rendered by Go through `gotossr`; `/admin` and `/admin/*` remain Vite SPA routes. Vite emits both a browser bundle and a single-file SSR bundle. gotossr evaluates the latter with memory history, while the browser hydrates that exact route tree with browser history.

## Build and run

```bash
make web-build
go run ./cmd/app
```

`make web-build` emits the Vite browser manifest/assets to `assets/dist` and the single-file server artifact to `assets/ssr/entry-server.js`. `assets/dist` is embedded in the Go binary; deploy `assets/ssr/entry-server.js` alongside it. The server artifact is intentionally external because gotossr must evaluate it at request time.

The current `go.mod` uses the same local `gotossr` checkout as the reference project (`/Users/tim/Codes/ai-project/gotossr`) for its TanStack Router compatibility fixes. Point that `replace` directive at your checkout until those fixes are released upstream.

During frontend development, use `pnpm --dir web dev` for Vite HMR. The Go SSR response is available from the application port after rebuilding assets with `make web-build`.

## Adding an SSR page

Add a normal file-based TanStack route under `web/src/routes`; do not create a second SSR-only component. The Vite SSR build resolves UnoCSS, SCSS Modules, aliases, and the same route tree used by the browser. Keep `/admin` routes out of SSR: `internal/routes/ssr.go` returns the normal Vite SPA shell for `/admin` and `/admin/*`.

## Dynamic title, meta, and page data

`internal/routes/ssr.go` resolves `SSRPageData` before calling `RenderRoute`. `Title` becomes `<title>`, while `Description` and `Keywords` become named meta tags. `FeaturedResource` on `/latest` is the complete data-flow example: Go serializes it into `__SSR_PROPS__`; `SSRDataProvider` supplies it to both `entry-server.tsx` and `main.tsx`; `LatestPage` reads it with `useSSRPageData()`. Hydration therefore sees identical initial data and does not issue a duplicate request.

For database-backed pages, replace the example value with a business-layer query using `c.Request.Context()` and route parameters. For later client navigation, fetch through an API route or a TanStack Router loader; browser-only APIs must not run during SSR.
