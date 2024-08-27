import ThemeProvider from "@/provider/ThemeProvider.tsx";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { RouterProvider } from "@tanstack/react-router";
import { FC, PropsWithChildren } from "react";
import AuthProvider, { useAuth } from "./provider/AuthProvider.tsx";
import { router } from "./router/router.tsx";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
    },
  },
});

function ToAuthRoute() {
  const auth = useAuth();
  return <RouterProvider router={router} context={auth} />;
}

type IApp = object;
const App: FC<PropsWithChildren<IApp>> = () => {
  console.log("app render...");

  return (
    <>
      <QueryClientProvider client={queryClient}>
        <ThemeProvider>
          <AuthProvider>
            <ToAuthRoute />
          </AuthProvider>
        </ThemeProvider>
      </QueryClientProvider>
    </>
  );
};

export default App;
