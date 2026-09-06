import { createFileRoute } from '@tanstack/react-router'
import clsx from 'clsx'

import { Button } from '@/components/Button'
import { Card } from '@/components/Card'

import classes from './home.module.scss'

export const Route = createFileRoute('/_public/')({
  component: HomePage,
})

const stats = [
  { value: '12.4k', label: '资源总量' },
  { value: '98.6%', label: '可用率' },
  { value: '48h', label: '平均响应' },
]

const features = [
  {
    icon: 'i-mdi-lightning-bolt',
    title: '极致性能',
    desc: '基于 TanStack Router 的细粒度代码分割，配合 React Compiler 与 UnoCSS，渲染快、包体小。',
  },
  {
    icon: 'i-mdi-server-network',
    title: 'SSR 就绪',
    desc: '内置 gotossr 服务端渲染管线，为爬虫与首屏提供完整、可索引的 HTML。',
  },
  {
    icon: 'i-mdi-palette-outline',
    title: '主题系统',
    desc: '基于 CSS 变量的明暗双主题与强调色切换，全站一致的设计令牌与动效。',
  },
  {
    icon: 'i-mdi-shield-check-outline',
    title: '类型安全',
    desc: '端到端类型编排，由 oxlint + tsgo 保障代码质量，Zod 校验表单与 API。',
  },
]

function HomePage() {
  return (
    <div className={classes.page}>
      <section className={classes.hero}>
        <div className={classes.heroBadge}>
          <span className="i-mdi-rocket-launch-outline" aria-hidden="true" />
          现代全栈模板
        </div>
        <h1 className={classes.heroTitle}>
          一套模板，快速构建
          <br />
          <span className={classes.heroAccent}>前台 + 后台</span>
        </h1>
        <p className={classes.heroSubtitle}>
          neter-template 提供优雅的默认布局、完整的主题体系与开箱即用的 SSR，
          让你专注于业务本身，而不是重复的样板代码。
        </p>
        <div className={classes.heroActions}>
          <Button
            to="/latest"
            variant="primary"
            size="lg"
            icon={<span className="i-mdi-rocket-launch-outline" />}
          >
            立即体验
          </Button>
          <Button
            to="/docs"
            variant="outline"
            size="lg"
            icon={<span className="i-mdi-book-open-page-variant-outline" />}
          >
            阅读文档
          </Button>
        </div>

        <div className={classes.stats}>
          {stats.map((stat) => (
            <div key={stat.label} className={classes.stat}>
              <span className={classes.statValue}>{stat.value}</span>
              <span className={classes.statLabel}>{stat.label}</span>
            </div>
          ))}
        </div>
      </section>

      <section className={classes.section}>
        <h2 className={classes.sectionTitle}>为什么选择它</h2>
        <p className={classes.sectionSubtitle}>开箱即用，并且可以按需裁剪。</p>

        <div className={clsx(classes.grid, classes.gridCols4)}>
          {features.map((feature) => (
            <Card key={feature.title} className={classes.featureCard}>
              <span
                className={clsx(feature.icon, classes.featureIcon)}
                aria-hidden="true"
              />
              <h3 className={classes.featureTitle}>{feature.title}</h3>
              <p className={classes.featureDesc}>{feature.desc}</p>
            </Card>
          ))}
        </div>
      </section>

      <section className={classes.cta}>
        <div className={classes.ctaInner}>
          <h2 className={classes.ctaTitle}>准备好开始了吗？</h2>
          <p className={classes.ctaDesc}>
            前往最新资源，或进入管理后台体验完整的后台操作流程。
          </p>
          <div className={classes.ctaActions}>
            <Button to="/latest" variant="primary" size="lg">
              浏览最新
            </Button>
            <Button
              to="/admin"
              variant="secondary"
              size="lg"
              icon={<span className="i-mdi-shield-account-outline" />}
            >
              进入后台
            </Button>
          </div>
        </div>
      </section>
    </div>
  )
}
