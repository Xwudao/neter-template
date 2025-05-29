import { User } from '@/api/userApi';
import NotFound from '@/components/admin/layout/NotFound';
import { type QueryClient } from '@tanstack/react-query';
import { Outlet, createRootRouteWithContext } from '@tanstack/react-router';

interface MyRouterContext {
  // The ReturnType of your useAuth hook or the value of your AuthContext
  auth: User;
  queryClient: QueryClient;
}

export const Route = createRootRouteWithContext<MyRouterContext>()({
  component: () => <Outlet />,
  notFoundComponent: () => <NotFound />,
});
