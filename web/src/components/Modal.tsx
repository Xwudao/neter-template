import clsx from 'clsx'
import {
  type MouseEvent,
  type PropsWithChildren,
  createContext,
  useContext,
  useEffect,
  useRef,
  useState,
} from 'react'

import classes from './modal.module.scss'

// Nested modals mount outer-first, so their document keydown listeners fire
// outer-first as well. Keep an explicit stack of open dialogs, ordered by
// nesting depth (then open order). The stack index drives both the z-index, so
// same-depth dialogs still stack by open order, and which dialog consumes
// Escape.
export const ModalLayerContext = createContext(0)

interface ModalLayer {
  depth: number
  close: (() => void) | null
}

const modalLayers: ModalLayer[] = []
let escapeListenerAttached = false

function handleEscapeKey(event: KeyboardEvent) {
  if (event.key !== 'Escape') return
  for (let index = modalLayers.length - 1; index >= 0; index -= 1) {
    const { close } = modalLayers[index]
    if (!close) continue
    event.preventDefault()
    event.stopPropagation()
    close()
    return
  }
}

function registerModalLayer(layer: ModalLayer) {
  const upperIndex = modalLayers.findIndex((item) => item.depth > layer.depth)
  const index = upperIndex === -1 ? modalLayers.length : upperIndex
  modalLayers.splice(index, 0, layer)

  if (!escapeListenerAttached) {
    document.addEventListener('keydown', handleEscapeKey, true)
    escapeListenerAttached = true
  }

  return index
}

function unregisterModalLayer(layer: ModalLayer) {
  const index = modalLayers.indexOf(layer)
  if (index >= 0) modalLayers.splice(index, 1)

  if (modalLayers.length === 0 && escapeListenerAttached) {
    document.removeEventListener('keydown', handleEscapeKey, true)
    escapeListenerAttached = false
  }
}

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
  const modalDepth = useContext(ModalLayerContext)
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

  // The stack index keeps the z-index and the Escape handling in the same order.
  const [layerIndex, setLayerIndex] = useState(0)
  // Keep the latest handler reachable without re-registering the layer on every
  // render, which would wrongly move an existing dialog to the top of the stack.
  const onCloseRef = useRef(onClose)
  onCloseRef.current = onClose

  useEffect(() => {
    if (!open) {
      setLayerIndex(0)
      return
    }

    const layer: ModalLayer = {
      depth: modalDepth,
      close: () => onCloseRef.current?.(),
    }
    setLayerIndex(registerModalLayer(layer))

    return () => unregisterModalLayer(layer)
  }, [open, modalDepth])

  const handleOverlayClick = (e: MouseEvent) => {
    if (e.target === overlayRef.current && onClose) {
      onClose()
    }
  }

  if (!open) return null

  return (
    <ModalLayerContext.Provider value={layerIndex + 1}>
      <div
        ref={overlayRef}
        onClick={handleOverlayClick}
        className={clsx(classes.overlay)}
        style={{ zIndex: `calc(var(--z-modal) + ${layerIndex})` }}
      >
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
    </ModalLayerContext.Provider>
  )
}

export default Modal
