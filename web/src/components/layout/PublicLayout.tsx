import { Link, Outlet, useRouterState } from '@tanstack/react-router'
import clsx from 'clsx'
import { type ReactNode } from 'react'

import { Logo } from '@/components/Logo'
import { ThemePicker } from '@/components/ThemePicker'
import { PUBLIC_NAV } from '@/constants/nav'

import classes from './public-layout.module.scss'

export interface PublicLayoutProps {
  children?: ReactNode
  className?: string
  maxWidth?: 'sm' | 'md' | 'lg' | 'full'
}

const maxWidthClass = {
  sm: classes.maxSm,
  md: classes.maxMd,
  lg: classes.maxLg,
  full: classes.maxFull,
}

/**
 * Public layout — sticky header, centered content container and footer.
 * Used by the front-facing (crawlable) routes.
 */
export function PublicLayout({ children, className, maxWidth = 'lg' }: PublicLayoutProps) {
  return (
    <div className={clsx(classes.layout, className)}>
      <PublicHeader />
      <main className={clsx(classes.main, maxWidthClass[maxWidth])}>{children ?? <Outlet />}</main>
      <PublicFooter />
    </div>
  )
}

function PublicHeader() {
  const pathname = useRouterState({ select: (s) => s.location.pathname })

  return (
    <header className={classes.header}>
      <div className={classes.headerInner}>
        <Link to="/" className={classes.brand}>
          <Logo />
          <span className={classes.brandName}>neter-template</span>
        </Link>

        <nav className={classes.nav} aria-label="主导航">
          {PUBLIC_NAV.map((item) => {
            const active = item.to === '/' ? pathname === '/' : pathname.startsWith(item.to)
            return (
              <Link
                key={item.to}
                to={item.to}
                className={clsx(classes.navLink, active && classes.navLinkActive)}
              >
                <span className={clsx(item.icon, classes.navIcon)} aria-hidden="true" />
                <span>{item.label}</span>
              </Link>
            )
          })}
        </nav>

        <div className={classes.headerRight}>
          <ThemePicker />
        </div>
      </div>
    </header>
  )
}

function PublicFooter() {
  return (
    <footer className={classes.footer}>
      <div className={classes.footerInner}>
        <div className={classes.footerBrand}>
          <Logo />
          <p className={classes.footerDesc}>
            A modern full-stack template with TanStack Router, React 19, UnoCSS and SSR.
          </p>
        </div>

        <div className={classes.footerNav}>
          <span className={classes.footerTitle}>导航</span>
          <div className={classes.footerLinks}>
            {PUBLIC_NAV.map((item) => (
              <Link key={item.to} to={item.to} className={classes.footerLink}>
                {item.label}
              </Link>
            ))}
          </div>
        </div>
      </div>

      <div className={classes.footerBottom}>
        <span>© {new Date().getFullYear()} neter-template. All rights reserved.</span>
      </div>
    </footer>
  )
}
