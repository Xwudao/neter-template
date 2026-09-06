import { createContext, useContext, type PropsWithChildren } from 'react'

export interface SSRResource {
  id: number
  title: string
  category: string
  size: string
  updatedAt: string
  icon: string
}

export interface SSRPageProps {
  __requestPath?: string
  featuredResource?: SSRResource
}

const SSRDataContext = createContext<SSRPageProps>({})

export function SSRDataProvider({ children, value }: PropsWithChildren<{ value: SSRPageProps }>) {
  return <SSRDataContext.Provider value={value}>{children}</SSRDataContext.Provider>
}

export function useSSRPageData() {
  return useContext(SSRDataContext)
}

export function readSSRPageData(): SSRPageProps {
  const el = document.getElementById('__SSR_PROPS__')
  if (!el?.textContent) return {}

  try {
    return JSON.parse(el.textContent) as SSRPageProps
  } catch {
    return {}
  }
}
