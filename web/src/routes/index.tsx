import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/')({
  component: IndexPage,
})

function IndexPage() {
  return (
    <main style={{ padding: '2rem', fontFamily: 'var(--font-sans)' }}>
      <h1 style={{ color: 'var(--text-h)' }}>neter-template</h1>
      <p style={{ color: 'var(--text)' }}>Get started by editing src/routes/index.tsx</p>
    </main>
  )
}
