import { request } from '@/api/request.ts';

export interface PostUserLoginReq {
  username: string;
  password: string;
}

export interface User {
  username: string;
  role: string;
  id: number;
  create_time: string;
  update_time: string;

  token?: string;
}
export interface PostUserLoginRes {
  code: number;
  msg: string;
  data: {
    token: string;
    user: User;
  };
}

const postApiUserLogin = (payload: PostUserLoginReq) => {
  return request<PostUserLoginRes>({
    url: '/v1/user/login',
    method: 'post',
    data: payload,
  });
};

export interface GetUserInfoRes {
  code: number;
  msg: string;
  data: User;
}

const getApiUserInfo = () => {
  return request<GetUserInfoRes>({
    url: '/auth/v1/user/info',
    method: 'get',
  });
};

export { postApiUserLogin, getApiUserInfo };
