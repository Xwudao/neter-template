import { createFileRoute } from '@tanstack/react-router'
import clsx from 'clsx'

import { PageHeader } from '@/components/PageHeader'

import classes from './page.module.scss'

export const Route = createFileRoute('/_public/about')({
  component: AboutPage,
})

function AboutPage() {
  return (
    <div className={clsx(classes.page)}>
      <PageHeader title="关于 neter-template" subtitle="一个现代化的全栈起步模板" />

      <div className={clsx(classes.article)}>
        <h2>项目理念</h2>
        <p>
          neter-template 致力于解决“从零搭建一个带前台与后台的全栈应用”的重复劳动。
          它把路由、状态、主题、样式与 SSR 等基础设施配置好，让你可以直接开始写业务。
        </p>

        <h2>技术栈</h2>
        <ul>
          <li>React 19 + TanStack Router（文件式路由，自动代码分割）</li>
          <li>Zustand 全局状态 + React Query 数据请求</li>
          <li>UnoCSS（presetWind3）用于原子样式，SCSS Modules 用于组件样式</li>
          <li>基于 CSS 变量的明暗主题与强调色体系</li>
          <li>gotossr 服务端渲染（SSR）管线</li>
        </ul>

        <h2>目录结构</h2>
        <ul>
          <li><code>web/src/routes/</code> — 前台与后台路由（<code>_public</code> 为前台布局，<code>admin</code> 为后台）</li>
          <li><code>web/src/components/</code> — 可复用的 UI 组件</li>
          <li><code>web/src/constants/</code> — 导航等常量定义</li>
          <li><code>web/src/ssr.tsx</code> — SSR 服务端入口</li>
        </ul>
      </div>
    </div>
  )
}
