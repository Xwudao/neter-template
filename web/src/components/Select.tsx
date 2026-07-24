import clsx from 'clsx'
import { forwardRef, type SelectHTMLAttributes } from 'react'

import classes from './select.module.scss'

interface SelectOption {
  value: string
  label: string
}

interface SelectProps extends SelectHTMLAttributes<HTMLSelectElement> {
  label?: string
  error?: string
  options: SelectOption[]
  placeholder?: string
}

const Select = forwardRef<HTMLSelectElement, SelectProps>(
  ({ label, error, options, placeholder, className, id, ...props }, ref) => {
    const selectId = id || label?.toLowerCase().replace(/\s+/g, '-')

    return (
      <div className={clsx(classes.field)}>
        {label && (
          <label htmlFor={selectId} className={clsx(classes.label)}>
            {label}
          </label>
        )}
        <select
          ref={ref}
          id={selectId}
          className={clsx(classes.select, error && classes.selectError, className)}
          {...props}
        >
          {placeholder && (
            <option value="" disabled>
              {placeholder}
            </option>
          )}
          {options.map((opt) => (
            <option key={opt.value} value={opt.value}>
              {opt.label}
            </option>
          ))}
        </select>
        {error && <span className={clsx(classes.error)}>{error}</span>}
      </div>
    )
  },
)

Select.displayName = 'Select'

export default Select
