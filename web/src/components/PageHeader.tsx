import clsx from 'clsx'
import { type ReactNode } from 'react'

import classes from './page-header.module.scss'

export interface PageHeaderProps {
  title: string
  subtitle?: string
  actions?: ReactNode
  className?: string
}

/** PageHeader — consistent page title, subtitle and optional actions. */
export function PageHeader({ title, subtitle, actions, className }: PageHeaderProps) {
  return (
    <div className={clsx(classes.root, className)}>
      <div className={classes.text}>
        <h1 className={classes.title}>{title}</h1>
        {subtitle && <p className={classes.subtitle}>{subtitle}</p>}
      </div>
      {actions && <div className={classes.actions}>{actions}</div>}
    </div>
  )
}

export default PageHeader
