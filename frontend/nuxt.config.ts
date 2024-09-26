// https://nuxt.com/docs/api/configuration/nuxt-config

export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
  ssr: false, // Set to true if you want server-side rendering
  runtimeConfig: {
      public: {
          apiBase: process.env.API_BASE || 'http://localhost:8080/api',
      },
  },
  devServer: {
    port: 3333,
  },
  modules: ['@nuxtjs/tailwindcss'],
  extends: [
    '@nuxt/examples-ui'
  ],
  css: ['@/assets/css/tailwind.css']
})
