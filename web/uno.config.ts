import presetIcons from '@unocss/preset-icons'
import { defineConfig, presetWind3 } from 'unocss'

export default defineConfig({
  presets: [
    presetWind3({
      preflight: false,
    }),
    presetIcons({
      scale: 1,
      prefix: 'i-',
      extraProperties: {
        color: 'unset',
        display: 'inline-block',
        'vertical-align': 'middle',
      },
      customizations: {
        customize(props) {
          props.width = '1em'
          props.height = '1em'
          return props
        },
      },
    }),
  ],
})
