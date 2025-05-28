import { defineConfig, mergeConfig } from 'vite';
import { analyzer } from 'vite-bundle-analyzer';
import { commonConfig } from './config/vite.common';
import { productionConfig } from './config/vite.production';
import { libConfig } from './config/vite.lib';

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  console.log('mode:', mode);

  if (mode === 'lib') {
    return mergeConfig(commonConfig, libConfig);
  }

  if (mode === 'production') {
    return mergeConfig(commonConfig, productionConfig);
  }

  if (mode === 'analyze') {
    commonConfig.plugins?.push(analyzer());
  }

  return commonConfig;
});
