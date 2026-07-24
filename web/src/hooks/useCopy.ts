import { useCopyToClipboard } from 'react-use'

const useCopy = () => {
  const [, copy] = useCopyToClipboard()
  return copy
}

export default useCopy
