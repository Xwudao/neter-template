import clsx from 'clsx'
import { type MouseEvent, type PropsWithChildren, useEffect, useRef } from 'react'

import classes from './modal.module.scss'

interface ModalProps extends PropsWithChildren {
  open: boolean
  onClose?: () => void
  title?: string
  width?: string
  className?: string
  showClose?: boolean
}

function Modal({
  open,
  onClose,
  title,
  width = '480px',
  className,
  showClose = true,
  children,
}: ModalProps) {
  const overlayRef = useRef<HTMLDivElement>(null)

  useEffect(() => {
    if (open) {
      document.body.style.overflow = 'hidden'
    } else {
      document.body.style.overflow = ''
    }
    return () => {
      document.body.style.overflow = ''
    }
  }, [open])

  useEffect(() => {
    const handleEscape = (e: KeyboardEvent) => {
      if (e.key === 'Escape' && onClose) {
        onClose()
      }
    }

    if (open) {
      document.addEventListener('keydown', handleEscape)
    }
    return () => document.removeEventListener('keydown', handleEscape)
  }, [open, onClose])

  const handleOverlayClick = (e: MouseEvent) => {
    if (e.target === overlayRef.current && onClose) {
      onClose()
    }
  }

  if (!open) return null

  return (
    <div ref={overlayRef} onClick={handleOverlayClick} className={clsx(classes.overlay)}>
      <div
        className={clsx(classes.content, className)}
        style={{ '--modal-width': width } as React.CSSProperties}
      >
        {title && (
          <div className={clsx(classes.header)}>
            <h2 className={clsx(classes.title)}>{title}</h2>
            {showClose && (
              <button
                onClick={onClose}
                type="button"
                aria-label="Close"
                className={clsx(classes.closeBtn)}
              >
                <span className="i-mdi-close" aria-hidden="true" />
              </button>
            )}
          </div>
        )}
        {children}
      </div>
    </div>
  )
}

export default Modal
