const path = require('path')

module.exports = {  
  // publicPath: './',  
  outputDir: './dist',
  assetsDir: 'static',   // 放置静态资源的目录
  indexPath: 'index.html', // index.html 的输出路径 (相对于 outputDir)
  filenameHashing: true,  // 文件名hash 以便更好的控制缓存

  lintOnSave: true,   // 是否在保存的时候eslint检查
  runtimeCompiler: false,  // 是否使用包含运行时编译器的 Vue 构建版本
  productionSourceMap: false,  // 生产环境是否生成 sourceMap 文件
  integrity: false,  

  chainWebpack: (config) => {
      // config.module
      //     .rule('svgIcon')
      //     .test(/\.svg$/)
      //     .include.add(resolve('src/common/icons'))
      //     .end()
      //     .use('svg-sprite-loader')
      //     .loader('svg-sprite-loader')
      //     .tap(options => {
      //       options = {
      //         symbolId: 'icon-[name]'
      //       }
      //       return options
      //     })
      // config.module
      //   .rule('svg')
      //   .exclude.add(resolve('src/common/icons'))
      //   .end()
  },
  configureWebpack: (config) => {    
  if (process.env.NODE_ENV === 'production') {      
      config.mode = 'production';
      config["performance"] = {//打包文件大小配置
        "maxEntrypointSize": 10000000,
        "maxAssetSize": 30000000
      }
    } else {      
      config.mode = 'development'
    }
  },  
  css: {    
    extract: true,   // 是否使用css分离插件 ExtractTextPlugin
    sourceMap: false,  // 开启 CSS source maps?
    loaderOptions: {  // css预设器配置项
    },
    requireModuleExtension: true  // 启用 CSS modules for all css / pre-processor files.
  },  

  // 是否使用 thread-loader
  parallel: require('os').cpus().length > 1, 

  // PWA 插件相关配置
  pwa: {}, 

  // webpack-dev-server 相关配置
  devServer: {
    open: true,
    host: 'localhost',
    port: 8080,
    https: false,
    proxy: {      
      '/api': {
        target: "http://127.0.0.1:8081",
        ws: true,
        changeOrigin: true,
        pathRewrite: {          
            '^/api': '/api'
        }
      },
      // secure: false 
    },
    before: (app) => {}
  }, 
  // 第三方插件配置
  pluginOptions: {
  }
}
