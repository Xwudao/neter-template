import clsx from 'clsx'
import { type HTMLAttributes } from 'react'

import classes from './card.module.scss'

export interface CardProps extends HTMLAttributes<HTMLDivElement> {
  padded?: boolean
}

/** Card — a surface container for grouping content. */
export function Card({ padded = true, className, children, ...props }: CardProps) {
  return (
    <div className={clsx(classes.card, padded && classes.padded, className)} {...props}>
      {children}
    </div>
  )
}

export default Card
