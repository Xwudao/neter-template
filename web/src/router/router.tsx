import { routeTree } from '@/router/routes.tsx';
import { createHashHistory, createRouter } from '@tanstack/react-router';
const hashHistory = createHashHistory();

export const router = createRouter({
  routeTree,
  context: undefined!,
  history: hashHistory,
});
