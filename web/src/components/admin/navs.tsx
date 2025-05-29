import AppIcon from '@/components/AppIcon';
import type { NavItemPropsWithItems, SubNavPropsWithItems } from '@douyinfe/semi-ui/lib/es/navigation';
// import MaterialSymbolsRecommendOutline from '~icons/material-symbols/recommend-outline';

export type NavItems = (string | SubNavPropsWithItemsWithPath | NavItemPropsWithItemsWithPath)[];

export interface SubNavPropsWithItemsWithPath extends SubNavPropsWithItems {
  items?: (SubNavPropsWithItemsWithPath | string)[];
  path?: string;
}

export interface NavItemPropsWithItemsWithPath extends NavItemPropsWithItems {
  items?: (NavItemPropsWithItemsWithPath | string)[];
  path?: string;
}

const Navs: NavItems = [
  {
    itemKey: 'admin-index',
    text: '首页',
    icon: <AppIcon icon={'i-ant-design-home-outlined'} />,
    path: '/admin',
  },
  // {
  //   itemKey: 'list-category',
  //   text: '分类管理',
  //   icon: <Icon svg={<IconamoonCategoryLight />} size={`large`} />,
  //   path: '/admin/list_category',
  // },
  // {
  //   itemKey: 'list-site',
  //   text: '站点列表',
  //   icon: <Icon svg={<MaterialSymbolsAttachFile />} size={`large`} />,
  //   path: '/admin/list_site',
  // },
  {
    itemKey: 'data-list',
    text: '数据管理',
    icon: <AppIcon icon={'i-ant-design-ordered-list-outlined'} />,
    path: '/admin/data_list',
  },
  {
    itemKey: 'site-config',
    text: '站点配置',
    icon: <AppIcon icon={'i-ant-design-setting-outlined'} />,
    path: '/admin/config',
  },
];
const findKeyByPath = (navs: NavItems, path: string): string | undefined => {
  for (const navItem of navs) {
    if (typeof navItem === 'string') {
      continue;
    }

    if (navItem.path && navItem.path === path) {
      return navItem.itemKey as string;
    }

    if (navItem.items) {
      const keyInSubItems = findKeyByPath(navItem.items, path);
      if (keyInSubItems) {
        return keyInSubItems;
      }
    }
  }

  return undefined;
};

const findPathByKey = (navs: NavItems, key: string): string | undefined => {
  for (const navItem of navs) {
    if (typeof navItem === 'string') {
      continue;
    }

    if (navItem.itemKey === key && navItem.path) {
      return navItem.path;
    }

    if (navItem.items) {
      const pathInSubItems = findPathByKey(navItem.items, key);
      if (pathInSubItems) {
        return pathInSubItems;
      }
    }
  }

  return undefined;
};

export default Navs;
// eslint-disable-next-line react-refresh/only-export-components
export { findKeyByPath, findPathByKey };
