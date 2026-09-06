import {
  createBrowserHistory,
  createRouter,
  type RouterHistory,
} from '@tanstack/react-router'

import { routeTree } from './routeTree.gen'

export function createAppRouter(history: RouterHistory) {
  return createRouter({
    routeTree,
    history,
    defaultPreload: 'intent',
    defaultPendingMs: 120,
    defaultPendingMinMs: 240,
    scrollRestoration: true,
  })
}

export const router = createAppRouter(createBrowserHistory())

declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router
  }
}
