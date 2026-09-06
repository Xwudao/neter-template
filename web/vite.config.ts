import { fileURLToPath } from 'node:url'
import * as path from 'path'

import babel from '@rolldown/plugin-babel'
import { tanstackRouter } from '@tanstack/router-plugin/vite'
import react, { reactCompilerPreset } from '@vitejs/plugin-react'
import UnoCSS from 'unocss/vite'
import AutoImport from 'unplugin-auto-import/vite'
import IconsResolver from 'unplugin-icons/resolver'
import Icons from 'unplugin-icons/vite'
import { defineConfig } from 'vite'

const __dirname = path.dirname(fileURLToPath(import.meta.url))
const resolve = (p: string) => path.resolve(__dirname, p)

// https://vite.dev/config/
export default defineConfig(({ isSsrBuild }) => {
  const outDir = isSsrBuild ? resolve('../assets/ssr') : resolve('../assets/dist')

  return {
  // The Go SPA middleware serves `assets/dist` at the site root. Vite's
  // emitted file names already start with `assets/`, so adding `/assets/` as
  // a base produced `/assets/assets/...` URLs and made client-only /admin
  // receive HTML instead of JavaScript.
  base: '/',
  resolve: {
    alias: {
      '@': resolve('src'),
    },
  },
  plugins: [
    tanstackRouter({
      target: 'react',
      // The browser keeps route chunks. The synchronous gotossr runtime needs
      // one complete SSR graph because renderToString cannot await Suspense.
      autoCodeSplitting: !isSsrBuild,
    }),
    react(),
    babel({ presets: [reactCompilerPreset()] }),
    UnoCSS({
      configFile: resolve('./uno.config.ts'),
    }),
    AutoImport({
      dts: resolve('./src/auto-imports.d.ts'),
      imports: ['react'],
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
  define: {
    __BUILD_TIME__: JSON.stringify(new Date().toISOString()),
  },
  ssr: {
    // gotossr evaluates the server artifact in its embedded runtime, so React,
    // TanStack Router, and every application dependency must be in this file.
    noExternal: true,
  },
  build: {
    outDir,
    emptyOutDir: true,
    manifest: !isSsrBuild,
    // Vite intentionally leaves SSR output readable by default. gotossr loads
    // this artifact in production, so minify it just like the client bundle.
    minify: isSsrBuild ? 'oxc' : undefined,
    rollupOptions: isSsrBuild
      ? {
          output: {
            format: 'iife',
            inlineDynamicImports: true,
            entryFileNames: 'entry-server.js',
          },
        }
      : undefined,
  },
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
  }
})
