import {
  createRouter,
  type RouterHistory,
} from '@tanstack/react-router'

import { routeTree } from './routeTree.gen'

export function createAppRouter(history: RouterHistory, isServer = false) {
  return createRouter({
    routeTree,
    history,
    isServer,
    defaultPreload: 'intent',
    defaultPendingMs: 120,
    defaultPendingMinMs: 240,
    scrollRestoration: true,
  })
}

declare module '@tanstack/react-router' {
  interface Register {
    router: ReturnType<typeof createAppRouter>
  }
}
