import { zodResolver } from '@hookform/resolvers/zod'
import clsx from 'clsx'
import { type PropsWithChildren } from 'react'
import {
  useForm,
  type DefaultValues,
  type FieldValues,
  type UseFormReturn,
  FormProvider as RHFProvider,
} from 'react-hook-form'
import type { ZodSchema } from 'zod'

import classes from './form.module.scss'

interface FormProps<T extends FieldValues> extends PropsWithChildren {
  schema: ZodSchema<T>
  defaultValues?: DefaultValues<T>
  onSubmit: (values: T, form: UseFormReturn<T>) => void | Promise<void>
  className?: string
}

/**
 * Form wrapper that integrates react-hook-form with Zod validation.
 * Provides the form context to all children.
 */
function Form<T extends FieldValues>({
  schema,
  defaultValues,
  onSubmit,
  className,
  children,
}: FormProps<T>) {
  const form = useForm<T>({
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    resolver: zodResolver(schema as any),
    defaultValues,
  })

  const handleSubmit = form.handleSubmit(async (values) => {
    await onSubmit(values as T, form as UseFormReturn<T>)
  })

  return (
    <RHFProvider {...form}>
      <form onSubmit={handleSubmit} className={clsx(classes.form, className)} noValidate>
        {children}
      </form>
    </RHFProvider>
  )
}

export default Form
