import { createFileRoute, Link } from '@tanstack/react-router'
import clsx from 'clsx'

import { Card } from '@/components/Card'
import { PageHeader } from '@/components/PageHeader'

import classes from './page.module.scss'

export const Route = createFileRoute('/_public/docs')({
  component: DocsPage,
})

const docs = [
  { icon: 'i-mdi-route', title: '路由约定', to: '/docs', desc: '文件式路由、布局路由与自动代码分割' },
  { icon: 'i-mdi-palette-outline', title: '主题系统', to: '/docs', desc: 'CSS 变量令牌、明暗主题与强调色' },
  { icon: 'i-mdi-server-network', title: 'SSR 服务端渲染', to: '/docs', desc: 'gotossr 的接入与配置' },
  { icon: 'i-mdi-integration', title: '数据请求', to: '/docs', desc: 'React Query 与 API 类型生成' },
  { icon: 'i-mdi-form-select', title: '表单与校验', to: '/docs', desc: 'react-hook-form 与 Zod' },
  { icon: 'i-mdi-code-tags', title: '工程化', to: '/docs', desc: 'oxlint / oxfmt / tsgo 工作流' },
]

function DocsPage() {
  return (
    <div className={clsx(classes.page)}>
      <PageHeader title="文档" subtitle="引导你熟悉模板的各个部分。" />

      <div className={clsx(classes.grid, classes.grid3)}>
        {docs.map((doc) => (
          <Link key={doc.title} to={doc.to} className="no-underline">
            <Card className="h-full">
              <div className="flex flex-col gap-3">
                <span className={clsx(doc.icon, 'text-2xl text-[var(--color-accent)]')} aria-hidden="true" />
                <h3 className="m-0 text-[var(--color-text-strong)] font-semibold">{doc.title}</h3>
                <p className="m-0 text-sm text-[var(--color-text-muted)] leading-relaxed">{doc.desc}</p>
              </div>
            </Card>
          </Link>
        ))}
      </div>
    </div>
  )
}
