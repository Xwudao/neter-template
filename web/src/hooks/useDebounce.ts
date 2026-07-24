import { debounce } from 'lodash-es'
import { useMemo } from 'react'

/**
 * Returns a debounced version of the provided callback.
 */
function useDebounce<T extends (...args: unknown[]) => unknown>(
  callback: T,
  delay: number = 300,
): T {
  return useMemo(() => debounce(callback, delay) as unknown as T, [callback, delay])
}

export default useDebounce
