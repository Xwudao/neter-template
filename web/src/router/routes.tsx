import AdminIndexPageRoute from '@/pages/admin/AdminIndexPage.tsx';
import dataListPage from '@/pages/admin/DataListPage.tsx';
import listCategoryPage from '@/pages/admin/ListCategoryPage.tsx';
import listSitePage from '@/pages/admin/ListSitePage.tsx';
import siteConfigPage from '@/pages/admin/SiteConfigPage.tsx';
import rootRoute, { indexRoute } from '@/router/root.tsx';
import { createRoute, redirect } from '@tanstack/react-router';

export const adminPageRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/admin',
  beforeLoad: ({ context, location }) => {
    if (!context.logged) {
      throw redirect({ to: '/login', search: { redirect: location.href } });
    }
  },
}).lazy(() => import('../pages/AdminPage').then((d) => d.default));

export const loginPageRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/login',
}).lazy(() => import('../pages/LoginPage').then((d) => d.default));

export const routeTree = rootRoute.addChildren([
  indexRoute,
  loginPageRoute,
  adminPageRoute,
  AdminIndexPageRoute,
  listCategoryPage,
  listSitePage,
  siteConfigPage,
  dataListPage,
]);
