export interface NavItem {
  label: string
  to: string
  icon?: string
}

/** Public (front-facing) navigation — used by the header and footer. */
export const PUBLIC_NAV: NavItem[] = [
  { label: '首页', to: '/', icon: 'i-mdi-home-outline' },
  { label: '最新', to: '/latest', icon: 'i-mdi-clock-outline' },
  { label: '标签', to: '/tags', icon: 'i-mdi-tag-multiple-outline' },
  { label: '搜索', to: '/search', icon: 'i-mdi-magnify' },
  { label: '文档', to: '/docs', icon: 'i-mdi-file-document-outline' },
  { label: '关于', to: '/about', icon: 'i-mdi-information-outline' },
]

/** Admin navigation — grouped by section, used by the admin sidebar. */
export interface AdminNavGroup {
  type: 'group'
  icon: string
  label: string
  children: AdminNavChild[]
}

export interface AdminNavLink {
  type: 'link'
  icon: string
  label: string
  to: string
  exact?: boolean
}

export type AdminNavChild = AdminNavLink

export type AdminNavEntry = AdminNavLink | AdminNavGroup

export const ADMIN_NAV: AdminNavEntry[] = [
  {
    type: 'link',
    icon: 'i-mdi-view-dashboard-outline',
    label: '控制台',
    to: '/admin',
    exact: true,
  },
  {
    type: 'group',
    icon: 'i-mdi-account-group-outline',
    label: '内容管理',
    children: [
      { type: 'link', icon: 'i-mdi-account-outline', label: '用户', to: '/admin/users' },
      { type: 'link', icon: 'i-mdi-tag-multiple-outline', label: '标签', to: '/admin/tags' },
      { type: 'link', icon: 'i-mdi-format-list-bulleted', label: '资源', to: '/admin/resources' },
    ],
  },
  {
    type: 'group',
    icon: 'i-mdi-cog-outline',
    label: '系统',
    children: [
      { type: 'link', icon: 'i-mdi-tune', label: '站点设置', to: '/admin/settings' },
      { type: 'link', icon: 'i-mdi-list-box-outline', label: '系统日志', to: '/admin/logs' },
    ],
  },
]
