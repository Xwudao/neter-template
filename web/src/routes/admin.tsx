import { createFileRoute } from '@tanstack/react-router'

import { AdminShell } from '@/components/layout/AdminShell'

export const Route = createFileRoute('/admin')({
  component: AdminShell,
})
