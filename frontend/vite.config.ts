import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import vueDevTools from 'vite-plugin-vue-devtools'
import wails from '@wailsio/runtime/plugins/vite'
import tailwindcss from '@tailwindcss/vite'

// https://vite.dev/config/
export default defineConfig(({ mode }) => {
  const isDev = mode === 'development'

  return {
    plugins: [
      vue(),
      vueJsx(),
      // 【生产构建安全】DevTools 仅在开发模式下加载，生产包完全剥离
      ...(isDev
        ? [
            vueDevTools({
              launchEditor: 'C:\\Program Files\\JetBrains\\GoLand 2024.3.5\\bin\\goland64.exe',
            }),
          ]
        : []),
      wails('./bindings'),
      tailwindcss(),
    ],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url)),
        '#': fileURLToPath(new URL('./bindings', import.meta.url)),
      },
    },
  }
})
