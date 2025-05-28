import { defineConfig } from 'vite';

const isProd = process.env.NODE_ENV === 'production';

export const productionConfig = defineConfig({
  // base: "/",
  esbuild: {
    drop: isProd ? ['debugger', 'console'] : [],
  },
  build: {
    rollupOptions: {
      // output: {
      //   manualChunks: (id) => {
      //     const uiLibs = ["@douyinfe/semi-ui/"];
      //     if (uiLibs.some((lib) => id.includes(lib))) {
      //       return "ui";
      //     }
      //     const libs = [
      //       "lodash",
      //       "recharts",
      //       "js-base64",
      //       "@dnd-kit",
      //       "ace-builds",
      //       "md5",
      //       "mitt",
      //       "dayjs",
      //       "emotion",
      //       "axios",
      //       "draggable",
      //       "crypto-js",
      //     ];
      //     if (libs.some((lib) => id.includes(lib))) {
      //       return "libs";
      //     }
      //     return undefined;
      //   },
      // },
    },
  },
});
