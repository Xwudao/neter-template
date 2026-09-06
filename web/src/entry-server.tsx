import { createMemoryHistory, RouterProvider } from '@tanstack/react-router'
import { renderToString } from 'react-dom/server.browser'

import { ThemeProvider } from '@/provider/ThemeProvider'
import { createAppRouter } from '@/router'
import { SSRDataProvider, type SSRPageProps } from '@/ssr-data'

// gotossr prepends `var props = …` before evaluating this single-file Vite
// bundle. Keep this declaration ambient so it is erased from the output.
declare const props: SSRPageProps

function render(url: string): string {
  const router = createAppRouter(
    createMemoryHistory({ initialEntries: [url] }),
    true,
  )

  // The embedded gotossr runtime is synchronous. TanStack Router normally
  // resolves matches through router.load(); mark synchronous routes ready so
  // renderToString sees the same match as the browser's first render.
  router.isServer = true
  const location = router.latestLocation
  const matches = router.matchRoutes(location).map((match) => ({
    ...match,
    status: 'success' as const,
    _displayPending: false,
    _forcePending: false,
  }))
  router.stores.location.set(location)
  router.stores.resolvedLocation.set(location)
  router.stores.status.set('idle')
  router.stores.setMatches(matches)
  // RouterProvider reads these snapshots before an async router.load() cycle
  // has run in gotossr's synchronous runtime.
  Reflect.set(router, '_committed', matches)
  Reflect.set(router, '_rendered', [matches])

  return renderToString(
    <SSRDataProvider value={props}>
      <ThemeProvider>
        <RouterProvider router={router} />
      </ThemeProvider>
    </SSRDataProvider>, 
  )
}

const ssrGlobal = globalThis as typeof globalThis & { __ssr_result: string }

try {
  ssrGlobal.__ssr_result = render(props.__requestPath || '/')
} catch (error) {
  const detail = error instanceof Error ? error.stack || error.message : String(error)
  ssrGlobal.__ssr_result = `<!-- SSR bundle error: ${detail} -->`
}
