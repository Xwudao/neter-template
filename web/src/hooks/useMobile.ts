import { debounce } from 'lodash-es'
import { useEffect, useState } from 'react'

function useMobile() {
  const [isMobile, setIsMobile] = useState(false)
  const [width, setWidth] = useState(0)
  const [orientation, setOrientation] = useState<'portrait' | 'landscape'>('portrait')

  useEffect(() => {
    const userAgent = window.navigator.userAgent
    const isMobileDevice =
      /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(userAgent)
    setIsMobile(isMobileDevice)

    const handleResize = debounce(() => {
      setWidth(window.innerWidth)
      setOrientation(window.innerWidth > window.innerHeight ? 'landscape' : 'portrait')
    }, 250)

    handleResize()
    window.addEventListener('resize', handleResize)
    return () => window.removeEventListener('resize', handleResize)
  }, [])

  return { isMobile, width, orientation }
}

export default useMobile
