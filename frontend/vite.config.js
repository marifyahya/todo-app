import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    host: true, // Listen on all addresses, including LAN and public addresses
    strictPort: true,
    port: 5173,
    watch: {
      usePolling: true, // Required for hot reload to work reliably in Docker
    },
  },
})
