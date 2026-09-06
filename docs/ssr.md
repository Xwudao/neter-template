# Hybrid SSR

Public TanStack Router routes are rendered by Go through `gotossr`; `/admin` and `/admin/*` remain Vite SPA routes. Go uses memory history to render a crawlable public-page shell, then the Vite bundle mounts the full client application.

## Build and run

```bash
make web-build
go run ./cmd/app
```

`make web-build` emits the Vite manifest and assets to `assets/dist`, which are embedded in the Go binary. The SSR engine also evaluates `web/src/ssr.tsx` at runtime, so deploy the `web` source directory alongside the binary.

The current `go.mod` uses the same local `gotossr` checkout as the reference project (`/Users/tim/Codes/ai-project/gotossr`) for its TanStack Router compatibility fixes. Point that `replace` directive at your checkout until those fixes are released upstream.

During frontend development, use `pnpm --dir web dev` for Vite HMR. The Go SSR response is available from the application port after rebuilding assets with `make web-build`.

## Adding an SSR page

Add its server-renderable shell to `web/src/ssr.tsx`; this file deliberately avoids SCSS Modules because gotossr uses esbuild directly. Keep `/admin` routes out of that entry: `internal/routes/ssr.go` always returns the normal Vite SPA shell for `/admin` and `/admin/*`.

## Dynamic title, meta, and page data

`internal/routes/ssr.go` resolves `SSRPageData` before calling `RenderRoute`. Its hard-coded `ssrPageDataForPath` entries are working examples: `Title` becomes `<title>`, while `Description` and `Keywords` become named meta tags. The same struct is passed as `RenderConfig.Props` to `web/src/ssr.tsx`.

For database-backed pages, replace that function with a business-layer query using `c.Request.Context()` and route parameters, then pass the queried title, summary, tags, and body data as `Props`. Do this in Go before `RenderRoute`; TanStack client loaders cannot query the database for this SSR integration.
