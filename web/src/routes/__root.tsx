import { Outlet, createRootRoute } from '@tanstack/react-router'
import { lazy, Suspense } from 'react'

const isProduction = import.meta.env.PROD

const TanStackRouterDevtools = isProduction
  ? () => null
  : lazy(() =>
      import('@tanstack/react-router-devtools').then((res) => ({
        default: res.TanStackRouterDevtools,
      })),
    )

export const Route = createRootRoute({
  component: () => (
    <>
      <Outlet />
      <Suspense>
        <TanStackRouterDevtools />
      </Suspense>
    </>
  ),
})
