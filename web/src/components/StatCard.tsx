import clsx from 'clsx'
import { type ReactNode } from 'react'

import classes from './stat-card.module.scss'

export interface StatCardProps {
  label: string
  value: ReactNode
  icon?: string
  hint?: string
  tone?: 'accent' | 'success' | 'warning' | 'danger' | 'neutral'
  className?: string
}

/** StatCard — compact metric display for dashboards. */
export function StatCard({
  label,
  value,
  icon,
  hint,
  tone = 'neutral',
  className,
}: StatCardProps) {
  return (
    <div className={clsx(classes.card, className)}>
      <div className={clsx(classes.iconBadge, classes[tone])}>
        {icon && <span className={clsx(icon, classes.icon)} aria-hidden="true" />}
      </div>
      <div className={classes.body}>
        <span className={classes.label}>{label}</span>
        <span className={classes.value}>{value}</span>
        {hint && <span className={classes.hint}>{hint}</span>}
      </div>
    </div>
  )
}

export default StatCard
