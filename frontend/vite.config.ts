import { defineConfig } from 'vitest/config'
import { loadEnv } from 'vite'
import react from '@vitejs/plugin-react'
import { resolve } from 'node:path'

export default defineConfig(({ mode }) => {
  const rootEnvDir = resolve(import.meta.dirname, '..')
  const env = loadEnv(mode, rootEnvDir, '')

  return {
    plugins: [react()],

    envDir: rootEnvDir,

    server: {
      port: Number(env.FRONTEND_PORT || 5173),

      proxy: {
        '/api': {
          target: env.BACKEND_API_URL || 'http://localhost:8080',
          changeOrigin: true,
        },
      },
    },

    test: {
      environment: 'jsdom',
      globals: true,
      restoreMocks: true,
    },
  }
})


