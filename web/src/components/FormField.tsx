import clsx from 'clsx'
import { type ReactNode } from 'react'
import { type FieldValues, type Path, type UseFormReturn } from 'react-hook-form'

import classes from './form-field.module.scss'

interface FormFieldProps<T extends FieldValues> {
  name: Path<T>
  form: UseFormReturn<T>
  label?: string
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  children: (props: Record<string, any>) => ReactNode
}

/**
 * A generic form field wrapper that connects react-hook-form with custom inputs.
 */
function FormField<T extends FieldValues>({
  name,
  form,
  label,
  children,
}: FormFieldProps<T>) {
  const {
    register,
    formState: { errors },
  } = form

  const error = errors[name]
  const errorMessage = error?.message as string | undefined
  const { ref: _ref, ...fieldProps } = register(name)

  return (
    <div className={clsx(classes.field)}>
      {label && <label className={clsx(classes.label)}>{label}</label>}
      {children({
        ...fieldProps,
        error: errorMessage,
        name,
      })}
      {errorMessage && <span className={clsx(classes.error)}>{errorMessage}</span>}
    </div>
  )
}

export default FormField
