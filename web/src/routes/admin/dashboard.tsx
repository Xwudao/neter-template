import { createFileRoute } from '@tanstack/react-router';

export const Route = createFileRoute('/admin/dashboard')({
  component: () => <div>Hello /admin/dashboard!</div>,
});
