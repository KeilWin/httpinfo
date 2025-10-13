import { defineConfig } from 'vite'
import solid from 'vite-plugin-solid'

export default defineConfig({
  plugins: [solid()],
  resolve: {
    extensions: ['.tsx', '.ts', '.jsx', '.js'] // Порядок разрешения
  },
  css: {
    modules: {
      localsConvention: "camelCaseOnly",
    },
},
})
