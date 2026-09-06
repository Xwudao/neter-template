import { createFileRoute } from '@tanstack/react-router'
import clsx from 'clsx'

import { Card } from '@/components/Card'
import { StatCard } from '@/components/StatCard'

import classes from './dashboard.module.scss'

export const Route = createFileRoute('/admin/')({
  component: AdminDashboardPage,
})

const stats = [
  { label: '总用户', value: '1,284', icon: 'i-mdi-account-group-outline', tone: 'accent' as const, hint: '+12% 本周' },
  { label: '资源总量', value: '8,392', icon: 'i-mdi-file-multiple-outline', tone: 'neutral' as const, hint: '+96 今日' },
  { label: '今日访问', value: '3,548', icon: 'i-mdi-eye-outline', tone: 'success' as const, hint: '+8.4%' },
  { label: '活跃会话', value: '147', icon: 'i-mdi-access-point', tone: 'warning' as const, hint: '峰值 219' },
]

const recentUsers = [
  { name: '张伟', email: 'zhangwei@example.com', role: '管理员', status: '正常' },
  { name: '李娜', email: 'lina@example.com', role: '用户', status: '正常' },
  { name: '王强', email: 'wangqiang@example.com', role: '用户', status: '已禁用' },
  { name: '刘洋', email: 'liuyang@example.com', role: '用户', status: '正常' },
]

const activities = [
  { icon: 'i-mdi-file-plus', text: '新增资源「现代前端工程化实践指南」', time: '2 小时前' },
  { icon: 'i-mdi-account-plus', text: '用户「李娜」完成注册', time: '3 小时前' },
  { icon: 'i-mdi-alert-outline', text: '检测到异常登录尝试', time: '5 小时前' },
  { icon: 'i-mdi-cog', text: '站点设置已更新', time: '8 小时前' },
]

function AdminDashboardPage() {
  return (
    <div className={clsx(classes.page)}>
      <div className={clsx(classes.statsGrid)}>
        {stats.map((stat) => (
          <StatCard key={stat.label} {...stat} />
        ))}
      </div>

      <div className={clsx(classes.twoCol)}>
        <Card className={clsx(classes.section)}>
          <h2 className={classes.sectionTitle}>最近用户</h2>
          <table className={classes.table}>
            <thead>
              <tr>
                <th>用户</th>
                <th>角色</th>
                <th>状态</th>
              </tr>
            </thead>
            <tbody>
              {recentUsers.map((user) => (
                <tr key={user.email}>
                  <td>
                    <div className={classes.userCell}>
                      <span className={classes.avatar}>{user.name.charAt(0)}</span>
                      <div>
                        <div className={classes.userName}>{user.name}</div>
                        <div className={classes.userEmail}>{user.email}</div>
                      </div>
                    </div>
                  </td>
                  <td>{user.role}</td>
                  <td>
                    <span className={clsx(classes.status, user.status === '已禁用' && classes.statusDanger)}>
                      {user.status}
                    </span>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </Card>

        <Card className={clsx(classes.section)}>
          <h2 className={classes.sectionTitle}>近期动态</h2>
          <div className={classes.activityList}>
            {activities.map((activity, index) => (
              <div key={index} className={classes.activity}>
                <span className={clsx(activity.icon, classes.activityIcon)} aria-hidden="true" />
                <div className={classes.activityBody}>
                  <span className={classes.activityText}>{activity.text}</span>
                  <span className={classes.activityTime}>{activity.time}</span>
                </div>
              </div>
            ))}
          </div>
        </Card>
      </div>
    </div>
  )
}
