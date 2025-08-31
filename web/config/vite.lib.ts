import { resolve } from 'path';
import { defineConfig } from 'vite';

export const libConfig = defineConfig({
  build: {
    outDir: 'static',
    emptyOutDir: false,
    minify: 'esbuild',
    lib: {
      formats: ['iife'],
      entry: resolve(__dirname, '../src/lib-entry.ts'),
      name: 'main',
      fileName: () => `main.js`,
    },
  },
  define: { 'process.env.NODE_ENV': '"production"' },
});
