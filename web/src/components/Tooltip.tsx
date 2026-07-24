import clsx from 'clsx'
import { type PropsWithChildren, useState } from 'react'

import classes from './tooltip.module.scss'

interface TooltipProps extends PropsWithChildren {
  content: string
  position?: 'top' | 'bottom' | 'left' | 'right'
  className?: string
}

function Tooltip({ children, content, position = 'top', className }: TooltipProps) {
  const [visible, setVisible] = useState(false)

  return (
    <div
      className={clsx(classes.wrapper, className)}
      onMouseEnter={() => setVisible(true)}
      onMouseLeave={() => setVisible(false)}
    >
      {children}
      {visible && (
        <div className={clsx(classes.tooltip)} data-position={position}>
          {content}
        </div>
      )}
    </div>
  )
}

export default Tooltip
