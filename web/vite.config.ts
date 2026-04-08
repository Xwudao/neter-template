import react, { reactCompilerPreset } from '@vitejs/plugin-react'
import { tanstackRouter } from '@tanstack/router-plugin/vite'
import babel from '@rolldown/plugin-babel'
import * as path from 'path'
import { fileURLToPath } from 'node:url'
import AutoImport from 'unplugin-auto-import/vite'
import IconsResolver from 'unplugin-icons/resolver'
import Icons from 'unplugin-icons/vite'
import { defineConfig } from 'vite'
import checker from 'vite-plugin-checker'
import UnoCSS from 'unocss/vite'

const __dirname = path.dirname(fileURLToPath(import.meta.url))
const resolve = (p: string) => path.resolve(__dirname, p)

// https://vite.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      '@': resolve('src'),
    },
  },
  build: {
    manifest: true,
  },
  plugins: [
    tanstackRouter({
      target: 'react',
      autoCodeSplitting: true,
    }),
    react(),
    babel({ presets: [reactCompilerPreset()] }),
    checker({
      oxlint: true,
    }),
    UnoCSS({
      configFile: resolve('./uno.config.ts'),
    }),
    AutoImport({
      dts: resolve('./src/auto-imports.d.ts'),
      imports: [
        'react',
        {
          '@tanstack/react-router': [
            'Link',
            'Outlet',
            'useNavigate',
            'useRouter',
            'useRouterState',
          ],
        },
      ],
      resolvers: [
        IconsResolver({
          prefix: 'Icon',
        }),
      ],
    }),
    Icons({
      autoInstall: true,
      jsx: 'react',
      compiler: 'jsx',
      iconCustomizer(_collection, _icon, props) {
        props.width = '1em'
        props.height = '1em'
      },
    }),
  ],
  css: {
    modules: {
      localsConvention: 'camelCase',
    },
    preprocessorOptions: {
      scss: {
        additionalData: '@use "@/styles/mixins" as *;',
      },
    },
  },
})
