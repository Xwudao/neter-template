import ConfigProvider from '@/provider/ConfigProvider.tsx';
import useAuth from '@/provider/useAuth.tsx';
import { routeTree } from '@/routeTree.gen.ts';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { createRouter, RouterProvider, createBrowserHistory } from '@tanstack/react-router';
import AuthProvider from './provider/AuthProvider.tsx';
import ContentLoading from '@/components/loading/ContentLoading.tsx';
import { ErrorHolder } from '@/components/error/ErrorHolder.tsx';
import NotFound from '@/components/admin/layout/NotFound.tsx';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
    },
  },
});

const history = createBrowserHistory();

// Set up a Router instance
const router = createRouter({
  routeTree,
  defaultPreload: 'intent',
  defaultPendingComponent: ContentLoading,
  defaultPendingMs: 1, // 200ms 后显示 loading
  defaultPendingMinMs: 1000, // 最少显示 500ms
  defaultErrorComponent: ErrorHolder,
  defaultNotFoundComponent: NotFound,
  context: {
    queryClient,
    auth: undefined!,
  },
  history,
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
