import { request } from '@/api/request.ts';
import { NullResp } from '@/api/types.ts';

export interface SiteConfig {
  site_info: string;
  seo_config: string;
}

export interface GetGetSiteConfigRes {
  data: SiteConfig;
  code: number;
  msg: string;
}

const getApiGetSiteConfig = () => {
  return request<GetGetSiteConfigRes>({
    url: '/admin/v1/site_config/all',
    method: 'get',
  });
};

export interface PostUpdateSiteConfigReq {
  name: string;
  config: string;
}

const postApiUpdateSiteConfig = (payload: PostUpdateSiteConfigReq) => {
  return request<NullResp>({
    url: '/admin/v1/site_config/update',
    method: 'post',
    data: payload,
  });
};
export interface GetGenSitemapRes {
  code: number;
  msg: string;
  data: string;
}

const getApiGenSitemap = () => {
  return request<GetGenSitemapRes>({
    url: '/admin/v1/site_config/gen_sitemap',
    method: 'get',
  });
};

export interface GetGetCmnSiteConfigRes {
  code: number;
  msg: string;
  data: SiteConfig;
}

const getApiGetCmnSiteConfig = () => {
  return request<GetGetCmnSiteConfigRes>({
    url: '/v1/site_config/all',
    method: 'get',
  });
};

export { getApiGetSiteConfig, getApiGetCmnSiteConfig, getApiGenSitemap, postApiUpdateSiteConfig };
