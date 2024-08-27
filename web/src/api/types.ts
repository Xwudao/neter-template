export interface IBasicResp<T> {
  code: number;
  msg: string;
  data: T;
}

export const OK_CODE = 200;
export const PASS_ERR_CODE = 400;
export const Invalid_CODE = 401;
export const ERR_CODE = 0;
export const AUTH_CODE = 403;
export const LOGIN_CODE = 405;
export const WECHAT_PWD_CODE = 406;
// export const COIN_NOT_ENOUGH_CODE = 408;
export const NEED_VIP_CODE = 409;
export const NEED_BUY_CODE = 410;

export interface IUserInfo {
  email: string;
  nickname: string;
  role: string;
  token: string;
  username: string;
  id: number;
}

export interface IListResp<T> {
  code: number;
  msg: string;
  data: {
    list: T[];
    total: number;
  };
}

export interface StringResp extends IBasicResp<string> {}
export interface NullResp extends IBasicResp<null> {}
export interface NumResp extends IBasicResp<number> {}

export interface DeleteIDPayload {
  id: number;
}
