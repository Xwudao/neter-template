import { create } from 'zustand';
import { devtools, persist } from 'zustand/middleware';

export type SiteInfo = {
  site_name: string;
  site_url: string;
  site_title: string;
  sub_title: string;
  site_desc: string;
  site_logo: string;
  site_image: string;
  main_title: string;
  site_keywords: string[];
  site_meta_script: string;
  disclaimer: string;
};

export type SiteSeo = {
  meta_title: string;
  meta_description: string;
  meta_keywords: string[];
  og_title: string;
  og_description: string;
  og_image: string;
};

// 定义哪些字段需要从 JSON 字符串解析
export type RawConfig = {
  site_info: string;
  site_seo: string;
  // 可以继续添加其他需要解析的字段
};

// 解析后的配置类型
export type Config = {
  site_info: SiteInfo;
  site_seo: SiteSeo;
};

type ConfigState = {
  config: Config;
  hasCache: () => boolean; // 计算属性，检查配置是否已缓存
  updateConfig: (config: Partial<Config>) => void;
  resetConfig: () => void;
};

// 配置默认值工厂
const createDefaultConfig = (): Config => ({
  site_info: {
    site_name: '',
    site_url: '',
    site_title: '',
    sub_title: '',
    site_desc: '',
    site_logo: '',
    site_image: '',
    main_title: '',
    site_keywords: [],
    site_meta_script: '',
    disclaimer: '',
  },
  site_seo: {
    meta_title: '',
    meta_description: '',
    meta_keywords: [],
    og_title: '',
    og_description: '',
    og_image: '',
  },
});

export const initConfig = createDefaultConfig;

const useConfigState = create<ConfigState>()(
  devtools(
    persist(
      (set, get) => ({
        config: initConfig(),
        // hasCache: computed property to check if configuration has been cached
        hasCache() {
          return !!get().config.site_info.site_name;
        },
        updateConfig: (config: Partial<Config>) => set({ config: { ...get().config, ...config } }),
        resetConfig: () => set({ config: initConfig() }),
      }),
      {
        name: 'app-config',
      },
    ),
  ),
);

export default useConfigState;
