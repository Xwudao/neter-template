import { User } from '@/api/userApi';
import NotFound from '@/components/admin/layout/NotFound';
import ContentLoading from '@/components/loading/ContentLoading';
import { type QueryClient } from '@tanstack/react-query';
import { Outlet, createRootRouteWithContext } from '@tanstack/react-router';
import { Suspense } from 'react';

interface MyRouterContext {
  // The ReturnType of your useAuth hook or the value of your AuthContext
  auth: User;
  queryClient: QueryClient;
}

export const Route = createRootRouteWithContext<MyRouterContext>()({
  component: () => (
    <Suspense fallback={<ContentLoading />}>
      <Outlet />
    </Suspense>
  ),
  notFoundComponent: NotFound,
  pendingComponent: ContentLoading,
  pendingMinMs: 500, // 最少显示 500ms loading
  pendingMs: 0, // 如果超过 1s 才显示 loading
});
