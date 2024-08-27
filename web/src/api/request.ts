import { KEY_TOKEN } from '@/core/constants.ts';
import { API_URL } from '@/core/types.ts';
import { Toast } from '@douyinfe/semi-ui';
import axios, { Method } from 'axios';
import qs from 'query-string';

export const uploadStaticUrl = `${API_URL}/admin/v1/upload/static`;
export const cnBaseURL = import.meta.env.VITE_CN_API_URL;
const instance = axios.create({
  baseURL: API_URL,
  paramsSerializer: (params) => {
    return qs.stringify(params, { arrayFormat: 'none' }); // param=value1&param=value2
  },
});

//interceptors
instance.interceptors.request.use(
  (config) => {
    config.headers = config.headers || {};
    const token = localStorage.getItem(KEY_TOKEN);
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }

    // let time = new Date().getTime().toString();
    // let path = config.url || '';
    // config.headers['X-Time'] = time;
    // config.headers['X-Path'] = path;
    // if (window?.APP_CONFIG?.anti_key) {
    //   config.headers['X-Sign'] = encryptParams(
    //     hashCaptchaKey(window.APP_CONFIG.anti_key),
    //     [path, time],
    //   );
    //   config.headers['X-Body-Sign'] = encryptParams(
    //     hashCaptchaKey(window.APP_CONFIG.anti_key),
    //     [JSON.stringify(config.data)],
    //   );
    // }

    return config;
  },
  (error) => {
    console.log('network', error);
  },
);

instance.interceptors.response.use(
  (resp) => {
    return new Promise((resolve, reject) => {
      const status = resp.status || 0;
      const { code } = resp.data;
      if (status === 403 || (code && code === 403)) {
        Toast.error('请刷新页面登录');
        // location.href = '/login';
        reject(resp);
      }
      resolve(resp);
    });
  },
  (error) => {
    return new Promise((resolve, reject) => {
      const status = error.response.status || 0;
      if (status === 403) {
        // location.href = '/login';
        Toast.error('请刷新页面登录');
      }
      reject(error);
    });
  },
);
interface RequestConfig {
  method: Method;
  url: string;
  data?: { [key: string]: any };
  params?: { [key: string]: any };
  headers?: { [key: string]: string };
}

const request = <T>(config: RequestConfig) => {
  return new Promise<T>((resolve, reject) => {
    instance({
      method: config.method,
      url: config.url,
      data: config.data,
      params: config.params,
      headers: config.headers,
    })
      .then((res) => {
        resolve(res.data);
      })
      .catch((err) => {
        reject(err);
      });
  });
};
export { request };
export type { RequestConfig };
export default instance;
