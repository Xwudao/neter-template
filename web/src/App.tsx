import ThemeProvider from '@/provider/ThemeProvider.tsx';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { FC, PropsWithChildren } from 'react';
import AuthProvider from './provider/AuthProvider.tsx';
import { createRouter, RouterProvider } from '@tanstack/react-router';
import { routeTree } from '@/routeTree.gen.ts';
import useAuth from '@/provider/useAuth.tsx';

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

type IApp = object;
const App: FC<PropsWithChildren<IApp>> = () => {
  console.log('app render...');

  return (
    <>
      <QueryClientProvider client={queryClient}>
        <ThemeProvider>
          <AuthProvider>
            <AuthApp />
          </AuthProvider>
        </ThemeProvider>
      </QueryClientProvider>
    </>
  );
};

export default App;
