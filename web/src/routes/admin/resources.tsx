import { createFileRoute } from '@tanstack/react-router'
import clsx from 'clsx'

import { Button } from '@/components/Button'
import { PageHeader } from '@/components/PageHeader'

import classes from './page.module.scss'

export const Route = createFileRoute('/admin/resources')({
  component: AdminResourcesPage,
})

const resources = [
  { id: 1, title: '现代前端工程化实践指南', category: '教程', size: '12.4 MB', status: '已发布', icon: 'i-mdi-file-document' },
  { id: 2, title: 'Go 微服务架构设计', category: '开发', size: '8.1 MB', status: '已发布', icon: 'i-mdi-code-braces' },
  { id: 3, title: '设计系统全解析', category: '设计', size: '24.0 MB', status: '草稿', icon: 'i-mdi-palette' },
  { id: 4, title: '数据库索引优化手册', category: '数据库', size: '3.2 MB', status: '已发布', icon: 'i-mdi-database' },
  { id: 5, title: 'Kubernetes 实战部署', category: '运维', size: '18.7 MB', status: '待审核', icon: 'i-mdi-kubernetes' },
]

function AdminResourcesPage() {
  return (
    <div className={clsx(classes.page)}>
      <PageHeader
        title="资源管理"
        subtitle="管理站内资源内容。"
        actions={
          <Button variant="primary" size="md" icon={<span className="i-mdi-plus" />}>
            新增资源
          </Button>
        }
      />

      <div className={clsx(classes.tableWrap)}>
        <table className={clsx(classes.table)}>
          <thead>
            <tr>
              <th>资源</th>
              <th>分类</th>
              <th>大小</th>
              <th>状态</th>
              <th>操作</th>
            </tr>
          </thead>
          <tbody>
            {resources.map((resource) => (
              <tr key={resource.id}>
                <td>
                  <div className="flex items-center gap-3">
                    <span className={clsx(resource.icon, classes.iconCell)} aria-hidden="true" />
                    <span className="font-medium text-[var(--color-text-strong)]">{resource.title}</span>
                  </div>
                </td>
                <td>{resource.category}</td>
                <td className="text-[var(--color-text-muted)]">{resource.size}</td>
                <td>
                  <span
                    className={clsx(
                      classes.status,
                      resource.status === '已发布' && classes.statusSuccess,
                      resource.status === '待审核' && classes.statusWarning,
                      resource.status === '草稿' && classes.statusDanger,
                    )}
                  >
                    {resource.status}
                  </span>
                </td>
                <td>
                  <div className="flex gap-2">
                    <Button variant="ghost" size="sm">编辑</Button>
                    <Button variant="ghost" size="sm">下架</Button>
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
