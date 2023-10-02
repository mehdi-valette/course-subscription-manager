// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  typescript: { strict: true },
  devtools: { enabled: true },
  modules: ["@nuxtjs/tailwindcss", "nuxt-icons"],
  devServer: {
    port: 8080,
  },
});