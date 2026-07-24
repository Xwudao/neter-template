import { createFileRoute } from '@tanstack/react-router'
import clsx from 'clsx'

import { Logo } from '@/components/Logo'
import { ThemePicker } from '@/components/ThemePicker'

import classes from './index.module.scss'

export const Route = createFileRoute('/')({
  component: HomePage,
})

function HomePage() {
  return (
    <main className={clsx(classes.page)}>
      <Logo />
      <h1 className={clsx(classes.title)}>neter-template</h1>
      <p className={clsx(classes.desc)}>
        Start editing <code className={clsx(classes.code)}>src/routes/index.tsx</code>
      </p>
      <ThemePicker />
    </main>
  )
}
