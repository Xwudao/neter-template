import react from "@vitejs/plugin-react-swc";
import { resolve } from "path";
import UnoCSS from "unocss/vite";
import AutoImport from "unplugin-auto-import/vite";
import { FileSystemIconLoader } from "unplugin-icons/loaders";
import IconsResolver from "unplugin-icons/resolver";
import Icons from "unplugin-icons/vite";
import { defineConfig, mergeConfig } from "vite";
import { analyzer } from "vite-bundle-analyzer";

const isProd = process.env.NODE_ENV === "production";

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  console.log("mode:", mode);

  const productionConfig = defineConfig({
    base: "/app",
    esbuild: {
      drop: isProd ? ["debugger", "console"] : [],
    },
    build: {
      rollupOptions: {
        output: {
          manualChunks: (id) => {
            const uiLibs = ["@douyinfe/semi-ui/"];
            if (uiLibs.some((lib) => id.includes(lib))) {
              return "ui";
            }
            const libs = [
              "lodash",
              "recharts",
              "js-base64",
              "@dnd-kit",
              "ace-builds",
              "md5",
              "mitt",
              "dayjs",
              "emotion",
              "axios",
              "draggable",
              "crypto-js",
            ];
            if (libs.some((lib) => id.includes(lib))) {
              return "libs";
            }
            return undefined;
          },
        },
      },
    },
  });

  const cmnConfig = defineConfig({
    resolve: {
      alias: { "@": resolve("src") },
      extensions: [".tsx", ".js", ".ts"],
    },
    css: {
      modules: {
        localsConvention: "camelCase",
        generateScopedName: "[local]_[hash:base64:5]",
      },
      preprocessorOptions: {
        scss: {
          additionalData: `@import "@/assets/styles/app-mixins.scss";\n`,
        },
      },
    },
    plugins: [
      react(),
      AutoImport({
        imports: ["react"],
        resolvers: [
          IconsResolver({
            prefix: "Icon",
            extension: "jsx",
          }),
        ],
      }),
      Icons({
        autoInstall: true,
        jsx: "react",
        compiler: "jsx", // or 'solid'
        customCollections: {
          iconfont: FileSystemIconLoader("./src/assets/icons", (svg) =>
            svg.replace(/^<svg /, '<svg fill="currentColor" '),
          ),
        },
        iconCustomizer(collection, icon, props) {
          props.width = "1.2em";
          props.height = "1.2em";
        },
      }),
      UnoCSS(resolve("./uno.config.ts")),
    ],
  });

  const libConfig = defineConfig({
    build: {
      outDir: "static",
      emptyOutDir: false,
      minify: "esbuild",
      lib: {
        formats: ["iife"],
        entry: resolve(__dirname, "src/lib-entry.ts"),
        name: "main",
        fileName: () => `main.js`,
      },
      rollupOptions: {
        external: ["react", "react-dom"],
        output: {
          assetFileNames: function (chunkInfo) {
            if (chunkInfo?.name?.endsWith(".css")) {
              return "main.css";
            }
            return "";
          },
          globals: {
            react: "React",
            "react-dom": "ReactDOM",
          },
        },
      },
    },
    define: { "process.env.NODE_ENV": '"production"' },
  });

  if (mode === "lib") {
    return mergeConfig(cmnConfig, libConfig);
  }

  if (mode === "production") {
    return mergeConfig(cmnConfig, productionConfig);
  }

  if (mode === "analyze") {
    cmnConfig.plugins?.push(analyzer());
  }

  return cmnConfig;
});
