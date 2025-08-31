import { request } from '@/api/request.ts';
import { NullResp } from '@/api/types.ts';

export interface DataList {
  id: number;
  create_time: string;
  update_time: string;
  label: string;
  kind: string;
  key: string;
  value: string;
}

export interface PostCreateDataListReq {
  value: string;
  key: string;
  label: string;
  kind: string;
}

export interface PostCreateDataListRes {
  code: number;
  msg: string;
  data: DataList;
}

const postApiCreateDataList = (payload: PostCreateDataListReq) => {
  return request<PostCreateDataListRes>({
    url: '/admin/v1/data_list/create',
    method: 'post',
    data: payload,
  });
};
export interface PostUpdateDataListReq {
  id: number;
  key: string;
  value: string;
}

export interface PostUpdateDataListRes {
  msg: string;
  data: DataList;
  code: number;
}

const postApiUpdateDataList = (payload: PostUpdateDataListReq) => {
  return request<PostUpdateDataListRes>({
    url: '/admin/v1/data_list/update',
    method: 'post',
    data: payload,
  });
};
export interface GetListDataListByKindQuery {
  kind: string;
  page: number;
  size: number;
}

export interface GetListDataListByKindRes {
  data: {
    list: Array<DataList>;
    total: number;
  };
  code: number;
  msg: string;
}

const getAdminApiListDataListByKind = (query: GetListDataListByKindQuery) => {
  return request<GetListDataListByKindRes>({
    url: '/admin/v1/data_list/list',
    method: 'get',
    params: query,
  });
};

const getApiListDataListByKind = (query: GetListDataListByKindQuery) => {
  return request<GetListDataListByKindRes>({
    url: '/v1/data_list/list',
    method: 'get',
    params: query,
  });
};

export interface PostDeleteDataListReq {
  id: number;
}

const postApiDeleteDataList = (payload: PostDeleteDataListReq) => {
  return request<NullResp>({
    url: '/admin/v1/data_list/delete',
    method: 'post',
    data: payload,
  });
};
export interface GetGetDataListSortDataQuery {
  kinds: string[];
}

export interface GetGetDataListSortDataRes {
  code: number;
  msg: string;
  data: Array<DataList>;
}

const getAdminApiGetDataListSortData = (query: GetGetDataListSortDataQuery) => {
  return request<GetGetDataListSortDataRes>({
    url: '/admin/v1/data_list/sort_data',
    method: 'get',
    params: query,
  });
};
const getApiGetDataListSortData = (query: GetGetDataListSortDataQuery) => {
  return request<GetGetDataListSortDataRes>({
    url: '/v1/data_list/sort_data',
    method: 'get',
    params: query,
  });
};

export interface PostUpdateSortDataListReq {
  orders: Array<number>;
  ids: Array<number>;
}

const postApiUpdateSortDataList = (payload: PostUpdateSortDataListReq) => {
  return request<NullResp>({
    url: '/admin/v1/data_list/update_order',
    method: 'post',
    data: payload,
  });
};

export {
  postApiUpdateSortDataList,
  postApiCreateDataList,
  getAdminApiGetDataListSortData,
  getApiGetDataListSortData,
  postApiDeleteDataList,
  postApiUpdateDataList,
  getApiListDataListByKind,
  getAdminApiListDataListByKind,
};
