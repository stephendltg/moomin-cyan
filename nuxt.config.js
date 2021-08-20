const pkg = require('./package.json')
const capitalize = (s) => s[0].toUpperCase() + s.substr(1)

export default {
  target: 'static',
  router: {
    base: '/dist/'
  },
  env: {
    dev: process.env.NODE_ENV !== 'production',
    baseUrl: process.env.BASE_URL || 'http://localhost:3000',
    title: pkg.name.replace('-', ' '),
    debug: '*'
  },
  generate: {
    dir: 'dist'
  },
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: capitalize(pkg.name).replace('-', ' '),
    htmlAttrs: {
      lang: 'en'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
    '@/plugins/debug.js'
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/eslint
    '@nuxtjs/eslint-module',
    // https://go.nuxtjs.dev/stylelint
    '@nuxtjs/stylelint-module',
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
  ],

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
  }
}
