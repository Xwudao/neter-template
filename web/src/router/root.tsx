import NotfoundPage from "@/pages/NotfoundPage.tsx";
import {
  createRootRouteWithContext,
  createRoute,
  Navigate,
  Outlet,
} from "@tanstack/react-router";
import { AuthContextType } from "../provider/AuthProvider.tsx";

export const indexRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/",
  component: () => {
    return <Navigate to={`/admin`} />;
  },
  notFoundComponent: () => {
    return <NotfoundPage />;
  },
});

const rootRoute = createRootRouteWithContext<AuthContextType>()({
  notFoundComponent: () => {
    return <NotfoundPage />;
  },
  component: () => (
    <>
      <Outlet />
      {/*{!isProd && <TanStackRouterDevtools />}*/}
    </>
  ),
});

export default rootRoute;
