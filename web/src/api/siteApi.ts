import { Category } from '@/api/categoryApi.ts';
import { request } from '@/api/request.ts';
import { NullResp, StringResp } from '@/api/types.ts';

export interface Site {
  name: string;
  url: string;
  description: string;
  edges: {
    cates: Category[];
  };
  create_time: string;
  update_time: string;
  keywords: Array<string>;
  category_id: number[];
  id: number;
  icon: string;
}

export interface PostUpdateSiteReq {
  name: string;
  url: string;
  icon: string;
  description: string;
  keywords: Array<string>;
  category_id?: number;
  item_order?: number;
  id: number;
}

export interface PostUpdateSiteRes {
  code: number;
  msg: string;
  data: Site;
}

const postApiUpdateSite = (payload: PostUpdateSiteReq) => {
  return request<PostUpdateSiteRes>({
    url: '/admin/v1/site/update',
    method: 'post',
    data: payload,
  });
};
export interface PostCreateSiteReq {
  icon: string;
  description: string;
  keywords: Array<string>;
  category_id: number;
  name: string;
  url: string;
  item_order?: number;
}

export interface PostCreateSiteRes {
  code: number;
  msg: string;
  data: Site;
}

const postApiCreateSite = (payload: PostCreateSiteReq) => {
  return request<PostCreateSiteRes>({
    url: '/admin/v1/site/create',
    method: 'post',
    data: payload,
  });
};

export interface GetGetSiteMetaQuery {
  url: string;
}

export interface GetGetSiteMetaRes {
  msg: string;
  data: {
    title: string;
    description: string;
    keywords: string;
    favicon: string;
  };
  code: number;
}

const getApiGetSiteMeta = (query: GetGetSiteMetaQuery) => {
  return request<GetGetSiteMetaRes>({
    url: '/admin/v1/site/meta',
    method: 'get',
    params: query,
  });
};
export interface GetListSitesQuery {
  category_id: number;
  page: number;
  size: number;
  by_order?: string;
}

export interface GetListSitesRes {
  code: number;
  msg: string;
  data: {
    total: number;
    list: Array<Site>;
  };
}

const getApiListSites = (query: GetListSitesQuery) => {
  return request<GetListSitesRes>({
    url: '/admin/v1/site/list',
    method: 'get',
    params: query,
  });
};
export interface PostUploadIconReq {
  icon_url: string;
  site_url: string;
}

const postApiUploadIcon = (payload: PostUploadIconReq) => {
  return request<StringResp>({
    url: '/admin/v1/site/upload_icon',
    method: 'post',
    data: payload,
  });
};

export interface PostDeleteSiteReq {
  id: number;
}

const postApiDeleteSite = (payload: PostDeleteSiteReq) => {
  return request<NullResp>({
    url: '/admin/v1/site/delete',
    method: 'post',
    data: payload,
  });
};
export interface PostUpdateSortSiteReq {
  ids: Array<number>;
  orders: Array<number>;
}

const postApiUpdateSortSite = (payload: PostUpdateSortSiteReq) => {
  return request<NullResp>({
    url: '/admin/v1/site/update_order',
    method: 'post',
    data: payload,
  });
};
export interface GetGetSortDataSiteRes {
  code: number;
  msg: string;
  data: Array<{
    update_time: string;
    name: string;
    edges: {};
    id: number;
    create_time: string;
  }>;
}
export interface GetGetSortDataSiteReq {
  pid?: number;
}

const getApiGetSortDataSite = (req: GetGetSortDataSiteReq) => {
  return request<GetGetSortDataSiteRes>({
    url: '/admin/v1/site/sort_data',
    method: 'get',
    params: req,
  });
};
export interface PostWriteStaticFileReq {
  filename: string;
  data: string;
}

const postApiWriteStaticFile = (payload: PostWriteStaticFileReq) => {
  return request<NullResp>({
    url: '/admin/v1/site_config/write_file',
    method: 'post',
    data: payload,
  });
};

export {
  postApiUpdateSite,
  postApiDeleteSite,
  postApiUploadIcon,
  getApiListSites,
  postApiCreateSite,
  getApiGetSiteMeta,
  postApiUpdateSortSite,
  getApiGetSortDataSite,
  postApiWriteStaticFile,
};
