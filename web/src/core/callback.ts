import { ERR_CODE, IBasicResp, OK_CODE } from '@/api/types.ts';
import { Toast } from '@douyinfe/semi-ui';
export type ShowMsg = 'data' | string | boolean;

const onSuccess = <T>(showMsg: ShowMsg = true, callback?: (rtn: T) => void) => {
  return (rtn: IBasicResp<T>) => {
    switch (rtn.code) {
      case OK_CODE:
        if (showMsg === 'data' && rtn.data && typeof rtn.data === 'string') {
          Toast.success(rtn.data);
        } else if (showMsg === true) {
          Toast.success(rtn.msg);
        } else if (typeof showMsg === 'string') {
          Toast.success(showMsg);
        }
        callback?.(rtn.data);
        break;
      case ERR_CODE:
        Toast.error(rtn.msg);
    }
  };
};

// const onError = (e: any) => {
//   let msg = String(e);
//   if (msg) {
//     Toast.error(`发生错误了: ${msg}`);
//   }
// };

const onError = (callback?: () => void) => {
  return (e: any) => {
    try {
      const data = JSON.parse(e);
      const code = data['code'];
      if (code !== OK_CODE) {
        if (data['msg']) {
          Toast.error(String(data['msg']));
        }
      }
    } catch (e) {
      const msg = String(e);
      if (msg) {
        Toast.error(`发生错误了: ${msg}`);
      }
    }

    callback?.();
  };
};

export { onError, onSuccess };
