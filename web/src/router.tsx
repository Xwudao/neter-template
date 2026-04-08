import { createBrowserHistory, createRouter } from '@tanstack/react-router'
import { routeTree } from './routeTree.gen'

const history = createBrowserHistory()

export const router = createRouter({
  routeTree,
  history,
  defaultPreload: 'intent',
  defaultPendingMs: 120,
  defaultPendingMinMs: 240,
  scrollRestoration: true,
})

declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router
  }
}
