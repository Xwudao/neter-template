import ConfigProvider from '@/provider/ConfigProvider.tsx';
import useAuth from '@/provider/useAuth.tsx';
import { routeTree } from '@/routeTree.gen.ts';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { createRouter, RouterProvider } from '@tanstack/react-router';
import AuthProvider from './provider/AuthProvider.tsx';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
    },
  },
});

// Set up a Router instance
const router = createRouter({
  routeTree,
  defaultPreload: 'intent',
  context: {
    queryClient,
    auth: undefined!,
  },
});

// Register things for typesafety
declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router;
  }
}

const AuthApp = () => {
  const { user } = useAuth();
  return <RouterProvider router={router} context={{ auth: user }} />;
};

const App = () => {
  return (
    <>
      <QueryClientProvider client={queryClient}>
        <AuthProvider>
          <ConfigProvider>
              <AuthApp />
          </ConfigProvider>
        </AuthProvider>
      </QueryClientProvider>
    </>
  );
};

export default App;
