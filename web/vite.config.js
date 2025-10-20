import vue from '@vitejs/plugin-vue'
import { fileURLToPath, URL } from 'node:url'
import { defineConfig } from 'vite'
import eslint from 'vite-plugin-eslint'
import vueDevTools from 'vite-plugin-vue-devtools'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueDevTools(),
    eslint({
      fix: true,
      cache: false,
      lintOnStart: true,
      include: ['src/**/*.js', 'src/**/*.vue']
    })
  ],
  base: './',
  build: {
    emptyOutDir: true,
    outDir: '../docs'
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  }
})
