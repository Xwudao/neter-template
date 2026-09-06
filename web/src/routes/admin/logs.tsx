import { createFileRoute } from '@tanstack/react-router'
import clsx from 'clsx'

import { PageHeader } from '@/components/PageHeader'

import classes from './page.module.scss'

export const Route = createFileRoute('/admin/logs')({
  component: AdminLogsPage,
})

const logs = [
  { level: 'info', time: '2026-09-06 14:22:05', message: '用户 zhangwei 登录成功', source: 'auth' },
  { level: 'info', time: '2026-09-06 14:20:31', message: 'GET /api/v1/resources 200 12ms', source: 'http' },
  { level: 'warn', time: '2026-09-06 14:18:44', message: '异常登录尝试已记录', source: 'security' },
  { level: 'error', time: '2026-09-06 14:15:09', message: '数据库连接超时 (timeout after 5s)', source: 'db' },
  { level: 'info', time: '2026-09-06 14:10:52', message: '站点设置已更新', source: 'config' },
  { level: 'info', time: '2026-09-06 14:08:17', message: '任务调度器启动完成', source: 'cron' },
]

const levelClass: Record<string, string> = {
  info: classes.statusSuccess,
  warn: classes.statusWarning,
  error: classes.statusDanger,
}

function AdminLogsPage() {
  return (
    <div className={clsx(classes.page)}>
      <PageHeader title="系统日志" subtitle="查看系统运行与安全事件。" />

      <div className={clsx(classes.tableWrap)}>
        <table className={clsx(classes.table)}>
          <thead>
            <tr>
              <th>级别</th>
              <th>时间</th>
              <th>消息</th>
              <th>来源</th>
            </tr>
          </thead>
          <tbody>
            {logs.map((log, index) => (
              <tr key={index}>
                <td>
                  <span className={clsx(classes.status, levelClass[log.level])}>{log.level}</span>
                </td>
                <td className="text-[var(--color-text-muted)] whitespace-nowrap">{log.time}</td>
                <td className="font-mono text-xs">{log.message}</td>
                <td className="text-[var(--color-text-muted)]">{log.source}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  )
}
