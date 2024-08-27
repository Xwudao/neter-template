import { Icon } from "@douyinfe/semi-ui";
import type {
  NavItemPropsWithItems,
  SubNavPropsWithItems,
} from "@douyinfe/semi-ui/lib/es/navigation";
// import MaterialSymbolsRecommendOutline from '~icons/material-symbols/recommend-outline';
import AntDesignHomeOutlined from "~icons/ant-design/home-outlined";
import AntDesignOrderedListOutlined from "~icons/ant-design/ordered-list-outlined";
import AntDesignSettingOutlined from "~icons/ant-design/setting-outlined";

export type NavItems = (
  | string
  | SubNavPropsWithItemsWithPath
  | NavItemPropsWithItemsWithPath
)[];

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
    itemKey: "admin-index",
    text: "首页",
    icon: <Icon svg={<AntDesignHomeOutlined />} size={`large`} />,
    path: "/admin",
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
    itemKey: "data-list",
    text: "数据管理",
    icon: <Icon svg={<AntDesignOrderedListOutlined />} size={`large`} />,
    path: "/admin/data_list",
  },
  {
    itemKey: "site-config",
    text: "站点配置",
    icon: <Icon svg={<AntDesignSettingOutlined />} size={`large`} />,
    path: "/admin/site_config",
  },
];
export const findKeyByPath = (
  navs: NavItems,
  path: string,
): string | undefined => {
  for (const navItem of navs) {
    if (typeof navItem === "string") {
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

export const findPathByKey = (
  navs: NavItems,
  key: string,
): string | undefined => {
  for (const navItem of navs) {
    if (typeof navItem === "string") {
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
