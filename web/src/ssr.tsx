import {
  createRootRoute,
  createRoute,
  createRouter,
  Outlet,
  type RouterHistory,
} from '@tanstack/react-router'
import type { ReactNode } from 'react'

/**
 * Server-rendered entry used by gotossr. It intentionally does not import
 * client route modules: gotossr builds this file with esbuild, while the Vite
 * manifest links the global stylesheet that defines the `ssr-*` classes.
 */

export interface SSRProps {
  __requestPath?: string
  title?: string
  description?: string
  keywords?: string
}

const NAV = [
  { to: '/', label: '首页' },
  { to: '/latest', label: '最新' },
  { to: '/tags', label: '标签' },
  { to: '/search', label: '搜索' },
  { to: '/docs', label: '文档' },
  { to: '/about', label: '关于' },
]

function RootLayout() {
  return (
    <div className="ssr-layout">
      <header className="ssr-header">
        <div className="ssr-header-inner">
          <a href="/" className="ssr-brand">
            <span className="ssr-brand-mark">N</span>
            <span>neter-template</span>
          </a>
          <nav className="ssr-nav" aria-label="主导航">
            {NAV.map((item) => (
              <a key={item.to} href={item.to} className="ssr-nav-link">
                {item.label}
              </a>
            ))}
          </nav>
        </div>
      </header>

      <main className="ssr-main">
        <div className="ssr-container">
          <Outlet />
        </div>
      </main>

      <footer className="ssr-footer">
        © {new Date().getFullYear()} neter-template · Server-rendered by gotossr
      </footer>
    </div>
  )
}

function Page({ title, subtitle, children }: { title: string; subtitle?: string; children?: ReactNode }) {
  return (
    <>
      <h1 className="ssr-title">{title}</h1>
      {subtitle && <p className="ssr-subtitle">{subtitle}</p>}
      {children}
    </>
  )
}

const indexContent = () => (
  <Page title="neter-template" subtitle="一个现代的全栈模板：TanStack Router + React 19 + UnoCSS + SSR。">
    <div className="ssr-actions">
      <a href="/latest" className="ssr-action-primary">浏览最新资源</a>
      <a href="/docs" className="ssr-action-secondary">阅读文档</a>
    </div>
  </Page>
)

const aboutContent = () => <Page title="关于 neter-template" subtitle="一个现代化的全栈起步模板。" />
const docsContent = () => <Page title="文档" subtitle="引导你熟悉模板的各个部分。" />
const latestContent = () => <Page title="最新资源" subtitle="最近更新发布的资源列表。" />
const tagsContent = () => <Page title="标签" subtitle="按主题浏览站内内容。" />
const searchContent = () => <Page title="搜索" subtitle="输入关键字搜索站内内容。" />

function createRouteTree() {
  const rootRoute = createRootRoute({ component: RootLayout })

  const indexRoute = createRoute({ getParentRoute: () => rootRoute, path: '/', component: indexContent })
  const aboutRoute = createRoute({ getParentRoute: () => rootRoute, path: '/about', component: aboutContent })
  const docsRoute = createRoute({ getParentRoute: () => rootRoute, path: '/docs', component: docsContent })
  const latestRoute = createRoute({ getParentRoute: () => rootRoute, path: '/latest', component: latestContent })
  const tagsRoute = createRoute({ getParentRoute: () => rootRoute, path: '/tags', component: tagsContent })
  const searchRoute = createRoute({ getParentRoute: () => rootRoute, path: '/search', component: searchContent })

  return rootRoute.addChildren([indexRoute, aboutRoute, docsRoute, latestRoute, tagsRoute, searchRoute])
}

export function createSSRRouter({
  history,
}: {
  history: RouterHistory
  props?: SSRProps
}) {
  return createRouter({ routeTree: createRouteTree(), history })
}
