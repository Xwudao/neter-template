import react from '@vitejs/plugin-react-swc';
import { resolve } from 'path';
import UnoCSS from 'unocss/vite';
import AutoImport from 'unplugin-auto-import/vite';
import { FileSystemIconLoader } from 'unplugin-icons/loaders';
import IconsResolver from 'unplugin-icons/resolver';
import Icons from 'unplugin-icons/vite';
import { defineConfig } from 'vite';
import { tanstackRouter } from '@tanstack/router-plugin/vite';

export const commonConfig = defineConfig({
  resolve: {
    alias: { '@': resolve('src') },
    extensions: ['.tsx', '.js', '.ts'],
  },
  css: {
    modules: {
      localsConvention: 'camelCase',
      generateScopedName: '[local]_[hash:base64:5]',
    },
    preprocessorOptions: {
      scss: {
        additionalData: `@use "@/assets/styles/app-mixins.scss";\n`,
      },
    },
  },
  plugins: [
    tanstackRouter({
      target: 'react',
      quoteStyle: 'single',
      autoCodeSplitting: true,
    }),
    react(),
    AutoImport({
      imports: ['react'],
      resolvers: [
        IconsResolver({
          prefix: 'Icon',
          extension: 'jsx',
        }),
      ],
    }),
    Icons({
      autoInstall: true,
      jsx: 'react',
      compiler: 'jsx',
      customCollections: {
        iconfont: FileSystemIconLoader('./src/assets/icons', (svg) =>
          svg.replace(/^<svg /, '<svg fill="currentColor" '),
        ),
      },
      iconCustomizer(collection, icon, props) {
        props.width = '1.2em';
        props.height = '1.2em';
      },
    }),
    UnoCSS(resolve('./uno.config.ts')),
  ],
});
