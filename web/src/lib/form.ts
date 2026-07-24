import { type FieldErrors, type FieldValues } from 'react-hook-form'
import toast from 'react-hot-toast'

/**
 * Maps Zod validation errors into a flat record of field → message.
 */
export type ValidationErrors = Record<string, string>

/**
 * Find the first error message from a nested error object (react-hook-form style).
 */
function findFirstErrorMessage(value: unknown): string | undefined {
  if (!value || typeof value !== 'object') return undefined
  if ('message' in value && typeof value.message === 'string' && value.message) {
    return value.message
  }
  if (Array.isArray(value)) {
    for (const item of value) {
      const msg = findFirstErrorMessage(item)
      if (msg) return msg
    }
    return undefined
  }
  for (const nestedValue of Object.values(value)) {
    const msg = findFirstErrorMessage(nestedValue)
    if (msg) return msg
  }
  return undefined
}

/**
 * Show the first form validation error as a toast notification.
 */
export function showFormError<T extends FieldValues>(
  errors: FieldErrors<T>,
  fallback = 'Please check the form fields',
) {
  toast.error(findFirstErrorMessage(errors) ?? fallback)
}
