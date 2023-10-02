// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  typescript: { strict: true },
  devtools: { enabled: true },
  modules: ["@nuxtjs/tailwindcss", "nuxt-icons", "@pinia/nuxt"],
  devServer: {
    port: 8080,
  },
});
