import { useCallback, useState } from 'react'

/**
 * A hook for managing boolean state with toggle/on/off helpers.
 */
function useBoolean(initialValue = false) {
  const [value, setValue] = useState(initialValue)

  const setTrue = useCallback(() => setValue(true), [])
  const setFalse = useCallback(() => setValue(false), [])
  const toggle = useCallback(() => setValue((v) => !v), [])

  return { value, setTrue, setFalse, toggle, setValue }
}

export default useBoolean
