import { createFileRoute } from '@tanstack/react-router'
import clsx from 'clsx'

import { Button } from '@/components/Button'
import { PageHeader } from '@/components/PageHeader'

import classes from './page.module.scss'

export const Route = createFileRoute('/admin/tags')({
  component: AdminTagsPage,
})

const tags = [
  { id: 1, name: 'React', count: 128, color: '#61dafb' },
  { id: 2, name: 'Go', count: 96, color: '#00add8' },
  { id: 3, name: 'TypeScript', count: 84, color: '#3178c6' },
  { id: 4, name: '数据库', count: 63, color: '#e91e63' },
  { id: 5, name: '云原生', count: 52, color: '#326ce5' },
  { id: 6, name: '教程', count: 47, color: '#ff9800' },
]

function AdminTagsPage() {
  return (
    <div className={clsx(classes.page)}>
      <PageHeader
        title="标签管理"
        subtitle="维护站内标签。"
        actions={
          <Button variant="primary" size="md" icon={<span className="i-mdi-plus" />}>
            新建标签
          </Button>
        }
      />

      <div className={clsx(classes.tableWrap)}>
        <table className={clsx(classes.table)}>
          <thead>
            <tr>
              <th>标签</th>
              <th>关联资源</th>
              <th>操作</th>
            </tr>
          </thead>
          <tbody>
            {tags.map((tag) => (
              <tr key={tag.id}>
                <td>
                  <div className="flex items-center gap-3">
                    <span
                      className="inline-block w-3 h-3 rounded-full"
                      style={{ backgroundColor: tag.color }}
                    />
                    <span className="font-medium text-[var(--color-text-strong)]">{tag.name}</span>
                  </div>
                </td>
                <td className="text-[var(--color-text-muted)]">{tag.count}</td>
                <td>
                  <div className="flex gap-2">
                    <Button variant="ghost" size="sm">编辑</Button>
                    <Button variant="ghost" size="sm" className="text-[var(--color-danger)]">删除</Button>
                  </div>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  )
}
