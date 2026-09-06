import { createFileRoute } from '@tanstack/react-router'
import clsx from 'clsx'

import { Button } from '@/components/Button'
import { PageHeader } from '@/components/PageHeader'

import classes from './page.module.scss'

export const Route = createFileRoute('/admin/users')({
  component: AdminUsersPage,
})

const users = [
  { id: 1, name: '张伟', email: 'zhangwei@example.com', role: '管理员', status: '正常', joined: '2024-03-01' },
  { id: 2, name: '李娜', email: 'lina@example.com', role: '用户', status: '正常', joined: '2024-05-14' },
  { id: 3, name: '王强', email: 'wangqiang@example.com', role: '用户', status: '已禁用', joined: '2024-06-02' },
  { id: 4, name: '刘洋', email: 'liuyang@example.com', role: '用户', status: '正常', joined: '2024-07-19' },
  { id: 5, name: '陈静', email: 'chenjing@example.com', role: '用户', status: '待验证', joined: '2024-08-05' },
]

function AdminUsersPage() {
  return (
    <div className={clsx(classes.page)}>
      <PageHeader
        title="用户管理"
        subtitle="管理平台注册用户。"
        actions={
          <Button variant="primary" size="md" icon={<span className="i-mdi-account-plus-outline" />}>
            新建用户
          </Button>
        }
      />

      <div className={clsx(classes.toolbar)}>
        <div className={clsx(classes.toolbarLeft)}>
          <div className="flex items-center gap-2 px-3 py-2 rounded-lg bg-[var(--color-surface)] border border-[var(--color-border-strong)]">
            <span className="i-mdi-magnify text-[var(--color-text-muted)]" aria-hidden="true" />
            <input className="bg-transparent outline-none text-sm" placeholder="搜索用户名或邮箱…" />
          </div>
        </div>
        <Button variant="ghost" size="md" iconRight={<span className="i-mdi-filter-variant" />}>
          筛选
        </Button>
      </div>

      <div className={clsx(classes.tableWrap)}>
        <table className={clsx(classes.table)}>
          <thead>
            <tr>
              <th>用户</th>
              <th>角色</th>
              <th>状态</th>
              <th>注册时间</th>
              <th>操作</th>
            </tr>
          </thead>
          <tbody>
            {users.map((user) => (
              <tr key={user.id}>
                <td>
                  <div className="flex items-center gap-3">
                    <span className="inline-flex items-center justify-center w-8 h-8 rounded-full bg-[var(--color-accent-soft)] text-[var(--color-accent)] font-bold text-xs">
                      {user.name.charAt(0)}
                    </span>
                    <div>
                      <div className="font-medium text-[var(--color-text-strong)]">{user.name}</div>
                      <div className="text-xs text-[var(--color-text-muted)]">{user.email}</div>
                    </div>
                  </div>
                </td>
                <td>{user.role}</td>
                <td>
                  <span
                    className={clsx(
                      classes.status,
                      user.status === '正常' && classes.statusSuccess,
                      user.status === '已禁用' && classes.statusDanger,
                      user.status === '待验证' && classes.statusWarning,
                    )}
                  >
                    {user.status}
                  </span>
                </td>
                <td className="text-[var(--color-text-muted)]">{user.joined}</td>
                <td>
                  <div className="flex gap-2">
                    <Button variant="ghost" size="sm">编辑</Button>
                    <Button variant="ghost" size="sm" className="text-[var(--color-danger)]">禁用</Button>
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
