import { request } from '@/api/request.ts';
import { NullResp } from '@/api/types.ts';

export interface Category {
  description: string;
  parent_id: number;
  edges: {};
  id: number;
  create_time: string;
  update_time: string;
  name: string;
}
export interface PostCreateCategoryReq {
  pid: number;
  name: string;
  item_order: number;
  description: string;
}

export interface PostCreateCategoryRes {
  code: number;
  msg: string;
  data: Category;
}

const postApiCreateCategory = (payload: PostCreateCategoryReq) => {
  return request<PostCreateCategoryRes>({
    url: '/admin/v1/category/create',
    method: 'post',
    data: payload,
  });
};
export interface PostUpdateCategoryReq {
  id: number;
  pid?: number;
  item_order?: number;
  name: string;
  description: string;
}

export interface PostUpdateCategoryRes {
  code: number;
  msg: string;
  data: Category;
}

const postApiUpdateCategory = (payload: PostUpdateCategoryReq) => {
  return request<PostUpdateCategoryRes>({
    url: '/admin/v1/category/update',
    method: 'post',
    data: payload,
  });
};
export interface GetListAllCategoryRes {
  data: Array<Category>;
  code: number;
  msg: string;
}
export interface GetListAllCategoryReq {
  by_order?: string;
}

const getApiListAllCategory = (req: GetListAllCategoryReq) => {
  return request<GetListAllCategoryRes>({
    url: '/v1/category/all',
    method: 'get',
    params: req,
  });
};
export interface GetListCategoryQuery {
  page: number;
  pid?: number;
  size: number;
}

export interface GetListCategoryRes {
  code: number;
  msg: string;
  data: {
    list: Array<Category>;
    total: number;
  };
}

const getApiListCategory = (query: GetListCategoryQuery) => {
  return request<GetListCategoryRes>({
    url: '/admin/v1/category/list',
    method: 'get',
    params: query,
  });
};
export interface PostDeleteCategoryReq {
  id: number;
}

const postApiDeleteCategory = (payload: PostDeleteCategoryReq) => {
  return request<NullResp>({
    url: '/admin/v1/category/delete',
    method: 'post',
    data: payload,
  });
};
export interface PostUpdateSortCategoryReq {
  ids: Array<number>;
  orders: Array<number>;
}

const postApiUpdateSortCategory = (payload: PostUpdateSortCategoryReq) => {
  return request<NullResp>({
    url: '/admin/v1/category/update_order',
    method: 'post',
    data: payload,
  });
};

export interface GetGetSortDataCategoryRes {
  code: number;
  msg: string;
  data: Array<{
    id: number;
    create_time: string;
    update_time: string;
    name: string;
    edges: {};
  }>;
}

const getApiGetSortDataCategory = () => {
  return request<GetGetSortDataCategoryRes>({
    url: '/admin/v1/category/sort_data',
    method: 'get',
  });
};

export {
  postApiCreateCategory,
  getApiListCategory,
  postApiDeleteCategory,
  getApiListAllCategory,
  postApiUpdateCategory,
  postApiUpdateSortCategory,
  getApiGetSortDataCategory,
};
