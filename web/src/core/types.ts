export const EmptyFunc = () => {};
export const isProd = import.meta.env.PROD || false;
export const API_URL = import.meta.env.VITE_API_URL;
export const ICON_PROXY = import.meta.env.VITE_ICON_PROXY;
export const ICON_PIPE = `${API_URL}/v1/site/favicon?key=`;
