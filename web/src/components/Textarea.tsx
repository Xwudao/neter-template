import clsx from 'clsx'
import { forwardRef, type TextareaHTMLAttributes } from 'react'

import classes from './textarea.module.scss'

interface TextareaProps extends TextareaHTMLAttributes<HTMLTextAreaElement> {
  label?: string
  error?: string
}

const Textarea = forwardRef<HTMLTextAreaElement, TextareaProps>(
  ({ label, error, className, id, ...props }, ref) => {
    const textareaId = id || label?.toLowerCase().replace(/\s+/g, '-')

    return (
      <div className={clsx(classes.field)}>
        {label && (
          <label htmlFor={textareaId} className={clsx(classes.label)}>
            {label}
          </label>
        )}
        <textarea
          ref={ref}
          id={textareaId}
          className={clsx(classes.textarea, error && classes.textareaError, className)}
          {...props}
        />
        {error && <span className={clsx(classes.error)}>{error}</span>}
      </div>
    )
  },
)

Textarea.displayName = 'Textarea'

export default Textarea
