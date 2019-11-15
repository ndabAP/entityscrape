module.exports = {
  outputDir: '../website',
  devServer: {
    proxy: {
      '/api': {
        target: 'http://localhost:3000/api',
        changeOrigin: true,
        pathRewrite: {
          '^/api': ''
        }
      }
    }
  },

  transpileDependencies: [
    /[/\\]node_modules[/\\]veui[/\\]/,
    /[/\\]node_modules[/\\]vue-awesome[/\\]/,
    /[/\\]node_modules[/\\]resize-detector[/\\]/
  ],

  chainWebpack: config => {
    config.module.rule('eslint').use('eslint-loader').options({ fix: true })

    const types = ['vue-modules', 'vue', 'normal-modules', 'normal']
    types.forEach(type => {
      config.module
        .rule('less')
        .oneOf(type)
        .use('less-loader')
        .tap(options => Object.assign({}, options, {
          javascriptEnabled: true
        }))
    })
  }
}
