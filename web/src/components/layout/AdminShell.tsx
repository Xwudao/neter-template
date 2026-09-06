import { Link, Outlet, useRouterState } from '@tanstack/react-router'
import clsx from 'clsx'
import { type ReactNode } from 'react'

import { Logo } from '@/components/Logo'
import { ThemePicker } from '@/components/ThemePicker'
import { ADMIN_NAV, type AdminNavGroup, type AdminNavLink } from '@/constants/nav'

import classes from './admin-layout.module.scss'

export interface AdminShellProps {
  children?: ReactNode
}

/**
 * Admin shell — fixed sidebar navigation, top bar and scrollable content.
 * Admin routes are client-only and intentionally not server-rendered.
 */
export function AdminShell({ children }: AdminShellProps) {
  return (
    <div className={classes.shell}>
      <AdminSidebar />
      <div className={classes.body}>
        <AdminTopbar />
        <main className={classes.content}>{children ?? <Outlet />}</main>
      </div>
    </div>
  )
}

function AdminSidebar() {
  return (
    <aside className={classes.sidebar}>
      <Link to="/admin" className={classes.sidebarBrand}>
        <Logo />
        <span className={classes.sidebarBrandName}>管理后台</span>
      </Link>

      <nav className={classes.sidebarNav} aria-label="后台导航">
        {ADMIN_NAV.map((entry) =>
          entry.type === 'link' ? (
            <SidebarLink key={entry.to} link={entry} />
          ) : (
            <AdminGroup key={entry.label} group={entry} />
          ),
        )}
      </nav>

      <div className={classes.sidebarFooter}>
        <Link to="/" className={classes.backLink}>
          <span className="i-mdi-arrow-left" aria-hidden="true" />
          返回前台
        </Link>
      </div>
    </aside>
  )
}

function AdminGroup({ group }: { group: AdminNavGroup }) {
  return (
    <div className={classes.navGroup}>
      <span className={classes.navGroupLabel}>{group.label}</span>
      <div className={classes.navGroupItems}>
        {group.children.map((child) => (
          <SidebarLink key={child.to} link={child} />
        ))}
      </div>
    </div>
  )
}

function SidebarLink({ link }: { link: AdminNavLink }) {
  const pathname = useRouterState({ select: (s) => s.location.pathname })
  const active = link.exact ? pathname === link.to : pathname.startsWith(link.to)

  return (
    <Link
      to={link.to}
      className={clsx(classes.navLink, active && classes.navLinkActive)}
    >
      <span className={clsx(link.icon, classes.navLinkIcon)} aria-hidden="true" />
      <span>{link.label}</span>
    </Link>
  )
}

function AdminTopbar() {
  const pathname = useRouterState({ select: (s) => s.location.pathname })
  const segments = pathname.split('/').filter(Boolean)
  const title = segments.length ? segments[segments.length - 1] : '控制台'

  return (
    <header className={classes.topbar}>
      <div className={classes.topbarTitle}>{title}</div>
      <div className={classes.topbarActions}>
        <ThemePicker />
      </div>
    </header>
  )
}
