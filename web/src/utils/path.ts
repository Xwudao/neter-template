import { isProd } from '@/core/types.ts';
import parseUrl from 'parse-url';

const buildPublicPath = (p: string) => {
  console.log('import.meta.env', import.meta.env);
  return isProd ? `.${p}` : p;
};

const cleanUrl = (url: string) => {
  if (url.trim() === '') return '';
  const data = parseUrl(url);

  return `${data.protocol}://${data.resource}${data.pathname}`;
};

// get hostname
const getHostname = (url: string) => {
  if (url.trim() === '') return '';
  const data = parseUrl(url);

  return data.resource;
};

export { buildPublicPath, cleanUrl, getHostname };
