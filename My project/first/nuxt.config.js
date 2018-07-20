module.exports = {
  /*
  ** Headers of the page
  */
  head: {
    title: 'project1',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: 'Nuxt.js project' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      { rel: 'stylesheet', href: 'https://stackpath.bootstrapcdn.com/bootstrap/4.1.0/css/bootstrap.min.css' }
    ]
  },
  scripts: [
    { src: 'https://code.jquery.com/jquery-3.2.1.slim.min.js' },
    { src: 'https://use.fontawesome.com/releases/v5.0.13/js/all.js' },
    { src: 'https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js' },
    { src: 'https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js' },
    // { src: 'https://cdn.jsdelivr.net/npm/vuetify/dist/vuetify.js'}
  ],
  // plugins: [
  //   {src: '~/plugins/mixins.js', ssr:true}
  // ],
  /*
  ** Customize the progress bar color
  */
  loading: { color: '#3B8070' },
  /*
  ** Build configuration
  */
  build: {
    /*
    ** Run ESLint on save
    */
    extend(config, { isDev, isClient }) {
      if (isDev && isClient) {
        config.module.rules.push({
          enforce: 'pre',
          test: /\.(js|vue)$/,
          loader: 'eslint-loader',
          exclude: /(node_modules)/
        })
      }
    }
  }
}
