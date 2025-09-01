import { User } from '@/api/userApi';
import ContentLoading from '@/components/loading/ContentLoading';
import { type QueryClient } from '@tanstack/react-query';
import { Outlet, createRootRouteWithContext } from '@tanstack/react-router';
import { Suspense } from 'react';

interface MyRouterContext {
  // The ReturnType of your useAuth hook or the value of your AuthContext
  auth: User;
  isAdmin: boolean;
  logged: boolean;
  queryClient: QueryClient;
}

export const Route = createRootRouteWithContext<MyRouterContext>()({
  component: () => (
    <Suspense fallback={<ContentLoading />}>
      <Outlet />
    </Suspense>
  ),
});
