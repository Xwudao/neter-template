import clsx from 'clsx'
import { forwardRef, type InputHTMLAttributes } from 'react'

import classes from './input.module.scss'

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  label?: string
  error?: string
}

const Input = forwardRef<HTMLInputElement, InputProps>(
  ({ label, error, className, id, ...props }, ref) => {
    const inputId = id || label?.toLowerCase().replace(/\s+/g, '-')

    return (
      <div className={clsx(classes.field)}>
        {label && (
          <label htmlFor={inputId} className={clsx(classes.label)}>
            {label}
          </label>
        )}
        <input
          ref={ref}
          id={inputId}
          className={clsx(classes.input, error && classes.inputError, className)}
          {...props}
        />
        {error && <span className={clsx(classes.error)}>{error}</span>}
      </div>
    )
  },
)

Input.displayName = 'Input'

export default Input
