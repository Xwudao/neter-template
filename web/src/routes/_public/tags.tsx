import { createFileRoute, Link } from '@tanstack/react-router'
import clsx from 'clsx'

import { PageHeader } from '@/components/PageHeader'

import classes from './page.module.scss'

export const Route = createFileRoute('/_public/tags')({
  component: TagsPage,
})

const tags = [
  { name: 'React', count: 128 },
  { name: 'Go', count: 96 },
  { name: 'TypeScript', count: 84 },
  { name: '数据库', count: 63 },
  { name: '云原生', count: 52 },
  { name: '教程', count: 47 },
  { name: '设计', count: 38 },
  { name: '性能优化', count: 31 },
  { name: '部署', count: 27 },
  { name: '工程化', count: 24 },
  { name: '安全', count: 19 },
  { name: '微服务', count: 16 },
]

function TagsPage() {
  return (
    <div className={clsx(classes.page)}>
      <PageHeader title="标签" subtitle="按主题浏览站内内容。" />

      <div className="flex flex-wrap gap-3">
        {tags.map((tag) => (
          <Link
            key={tag.name}
            to="/tags"
            className={clsx(classes.chip)}
          >
            <span className="i-mdi-tag-outline" aria-hidden="true" />
            <span>{tag.name}</span>
            <span className="text-xs opacity-60">{tag.count}</span>
          </Link>
        ))}
      </div>
    </div>
  )
}
