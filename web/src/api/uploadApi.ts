import { API_URL } from '@/core/types.ts';

export const UploadS3Api = API_URL + '/admin/v1/upload/s3';

export interface UploadS3Res {
  code: number;
  msg: string;
  data: {
    key: string;
    url: string;
  };
}
