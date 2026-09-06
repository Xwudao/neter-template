import {
  createRootRoute,
  Link,
  Outlet,
} from '@tanstack/react-router'
import { lazy, Suspense } from 'react'

import { PublicLayout } from '@/components/layout/PublicLayout'

const isProduction = import.meta.env.PROD

const TanStackRouterDevtools = isProduction
  ? () => null
  : lazy(() =>
      import('@tanstack/react-router-devtools').then((res) => ({
        default: res.TanStackRouterDevtools,
      })),
    )

function NotFound() {
  return (
    <PublicLayout>
      <div className="flex flex-col items-center justify-center min-h-[50vh] gap-4 text-center">
        <span className="i-mdi-compass-outline text-4xl text-[var(--color-accent)]" />
        <h1 className="text-2xl font-bold text-[var(--color-text-strong)] m-0">
          页面不存在
        </h1>
        <p className="text-sm text-[var(--color-text-muted)] m-0">
          你访问的页面不存在或已被移除。
        </p>
        <Link to="/" className="mt-2">
          <span className="i-mdi-arrow-left mr-2" />
          返回首页
        </Link>
      </div>
    </PublicLayout>
  )
}

export const Route = createRootRoute({
  component: () => (
    <>
      <Outlet />
      <Suspense>
        <TanStackRouterDevtools />
      </Suspense>
    </>
  ),
  notFoundComponent: NotFound,
})
