import { createFileRoute } from '@tanstack/react-router'
import clsx from 'clsx'

import { Button } from '@/components/Button'
import { Card } from '@/components/Card'
import { PageHeader } from '@/components/PageHeader'

import classes from './page.module.scss'

export const Route = createFileRoute('/_public/latest')({
  component: LatestPage,
})

interface Resource {
  id: number
  title: string
  category: string
  size: string
  updatedAt: string
  icon: string
}

const resources: Resource[] = [
  { id: 1, title: '现代前端工程化实践指南', category: '教程', size: '12.4 MB', updatedAt: '2 小时前', icon: 'i-mdi-file-document' },
  { id: 2, title: 'Go 微服务架构设计与实现', category: '开发', size: '8.1 MB', updatedAt: '5 小时前', icon: 'i-mdi-code-braces' },
  { id: 3, title: '设计系统全解析', category: '设计', size: '24.0 MB', updatedAt: '8 小时前', icon: 'i-mdi-palette' },
  { id: 4, title: '数据库索引优化手册', category: '数据库', size: '3.2 MB', updatedAt: '昨天', icon: 'i-mdi-database' },
  { id: 5, title: 'Kubernetes 实战部署', category: '运维', size: '18.7 MB', updatedAt: '昨天', icon: 'i-mdi-kubernetes' },
  { id: 6, title: 'TypeScript 高级类型技巧', category: '开发', size: '5.9 MB', updatedAt: '2 天前', icon: 'i-mdi-language-typescript' },
]

function LatestPage() {
  return (
    <div className={clsx(classes.page)}>
      <PageHeader
        title="最新资源"
        subtitle="最近更新发布的资源列表。"
        actions={
          <Button to="/search" variant="outline" size="md" icon={<span className="i-mdi-magnify" />}>
            搜索
          </Button>
        }
      />

      <div className={clsx(classes.cardList)}>
        {resources.map((resource) => (
          <Card key={resource.id} padded={false} className={clsx(classes.listItem)}>
            <span className={clsx(resource.icon, classes.listItemIcon)} aria-hidden="true" />
            <div className={clsx(classes.listItemBody)}>
              <span className={clsx(classes.listItemTitle)}>{resource.title}</span>
              <span className={clsx(classes.meta)}>
                {resource.category}
                <span className={classes.metaDot} />
                {resource.size}
                <span className={classes.metaDot} />
                {resource.updatedAt}
              </span>
            </div>
            <Button variant="ghost" size="sm" iconRight={<span className="i-mdi-arrow-right" />}>
              查看
            </Button>
          </Card>
        ))}
      </div>
    </div>
  )
}
