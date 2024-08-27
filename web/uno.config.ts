import UnocssIcons from "@unocss/preset-icons";
import presetRemToPx from "@unocss/preset-rem-to-px";
import { defineConfig, presetUno } from "unocss";
import presetAnimations from "unocss-preset-animations";

export default defineConfig({
  cli: {
    entry: {
      patterns: ["front/**/*.tpl"],
      outFile: "static/utils.css",
    },
  },
  rules: [],
  shortcuts: {
    "line-center": "flex items-center",
    "block-center": "flex justify-center items-center",
    am: "align-middle",
    cp: "cursor-pointer",
    ovh: "overflow-hidden",
    "admin-toolbar": "flex gap-2 flex-wrap items-center",
  },
  theme: {},
  presets: [
    presetUno({ preflight: false }),
    presetRemToPx({ baseFontSize: 16 }),
    presetAnimations(),
    UnocssIcons({
      scale: 1,
      prefix: "i-",
      extraProperties: {
        color: "unset",
        "vertical-align": "middle",
        display: "inline-block",
      },
      customizations: {
        transform(svg: string) {
          return svg.replace(/#fff/, "currentColor");
        },
        customize(props: any) {
          props.width = "1.2em";
          props.height = "1.2em";
          return props;
        },
      },
    }),
  ],
});
