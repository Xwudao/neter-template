import { defineConfig } from 'vite';
import { compression } from 'vite-plugin-compression2';

const isProd = process.env.NODE_ENV === 'production';

export const productionConfig = defineConfig({
  // base: "/",
  esbuild: {
    drop: isProd ? ['debugger', 'console'] : [],
  },
  build: {
    manifest: true,
  },
  plugins: [compression()],
});
