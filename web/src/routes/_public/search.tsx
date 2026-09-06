import { createFileRoute } from '@tanstack/react-router'
import clsx from 'clsx'
import { useState } from 'react'

import { Button } from '@/components/Button'
import { PageHeader } from '@/components/PageHeader'

import classes from './page.module.scss'

export const Route = createFileRoute('/_public/search')({
  component: SearchPage,
})

/** 简单的前端搜索示例（静态过滤）。 */
const items = [
  { title: 'React 19 新特性', tag: 'React' },
  { title: 'Go 服务端渲染', tag: 'Go' },
  { title: 'TanStack Router 指南', tag: 'TypeScript' },
  { title: 'CSS 变量主题系统', tag: '教程' },
  { title: 'UnoCSS 实战', tag: '教程' },
  { title: '数据库索引优化', tag: '数据库' },
]

function SearchPage() {
  const [query, setQuery] = useState('')
  const results = query.trim()
    ? items.filter(
        (item) =>
          item.title.toLowerCase().includes(query.toLowerCase()) ||
          item.tag.toLowerCase().includes(query.toLowerCase()),
      )
    : []

  return (
    <div className={clsx(classes.page)}>
      <PageHeader title="搜索" subtitle="输入关键字搜索站内内容。" />

      <div className={clsx(classes.searchBox)}>
        <span className={clsx('i-mdi-magnify', classes.searchIcon)} aria-hidden="true" />
        <input
          className={clsx(classes.searchInput)}
          placeholder="搜索资源、文档、标签…"
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          autoFocus
        />
        {query && (
          <Button variant="ghost" size="sm" onClick={() => setQuery('')}>
            清除
          </Button>
        )}
      </div>

      {query && (
        <div className="flex flex-col gap-3">
          {results.length === 0 ? (
            <div className={clsx(classes.empty)}>
              <span className={clsx('i-mdi-magnify-close', classes.emptyIcon)} aria-hidden="true" />
              <span className={classes.emptyTitle}>未找到结果</span>
              <span className={classes.emptyDesc}>换个关键词试试。</span>
            </div>
          ) : (
            results.map((item) => (
              <div key={item.title} className={clsx(classes.chip, 'self-start')}>
                <span>{item.title}</span>
                <span className="text-xs opacity-60">{item.tag}</span>
              </div>
            ))
          )}
        </div>
      )}
    </div>
  )
}
