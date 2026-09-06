import { createFileRoute } from '@tanstack/react-router'
import clsx from 'clsx'
import { useState } from 'react'

import { Button } from '@/components/Button'
import Input from '@/components/Input'
import { PageHeader } from '@/components/PageHeader'

import classes from './page.module.scss'

export const Route = createFileRoute('/admin/settings')({
  component: AdminSettingsPage,
})

function AdminSettingsPage() {
  const [saved, setSaved] = useState(false)

  const handleSave = (e: React.FormEvent) => {
    e.preventDefault()
    setSaved(true)
    window.setTimeout(() => setSaved(false), 2000)
  }

  return (
    <div className={clsx(classes.page)}>
      <PageHeader title="站点设置" subtitle="配置站点的基础信息与偏好。" />

      <form className={clsx(classes.form)} onSubmit={handleSave}>
        <Input label="站点名称" defaultValue="neter-template" placeholder="输入站点名称" />
        <Input label="站点描述" defaultValue="现代全栈模板" placeholder="一句话描述站点" />

        <div className={clsx(classes.fieldRow)}>
          <Input label="联系邮箱" type="email" defaultValue="admin@example.com" placeholder="admin@example.com" />
          <Input label="备案号" defaultValue="" placeholder="ICP 备案号（可选）" />
        </div>

        <div className="flex flex-col gap-2">
          <span className="text-sm font-medium text-[var(--color-text-strong)]">功能开关</span>
          <label className="flex items-center justify-between py-2 px-3 rounded-lg bg-[var(--color-surface)] border border-[var(--color-border-soft)]">
            <span className="text-sm">允许匿名搜索</span>
            <input type="checkbox" defaultChecked className="accent-[var(--color-accent)]" />
          </label>
          <label className="flex items-center justify-between py-2 px-3 rounded-lg bg-[var(--color-surface)] border border-[var(--color-border-soft)]">
            <span className="text-sm">开启注册</span>
            <input type="checkbox" defaultChecked className="accent-[var(--color-accent)]" />
          </label>
          <label className="flex items-center justify-between py-2 px-3 rounded-lg bg-[var(--color-surface)] border border-[var(--color-border-soft)]">
            <span className="text-sm">维护模式</span>
            <input type="checkbox" className="accent-[var(--color-accent)]" />
          </label>
        </div>

        <div className="flex items-center gap-3">
          <Button type="submit" variant="primary" size="md" icon={<span className="i-mdi-content-save-outline" />}>
            保存设置
          </Button>
          {saved && <span className="text-sm text-[var(--color-success)]">已保存 ✓</span>}
        </div>
      </form>
    </div>
  )
}
