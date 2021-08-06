const path = require('path')

// 优化一：使用cdn
// const cdn = {
//   // 忽略打包的第三方库
//   externals: {
//     vue: 'Vue',
//     vuex: 'Vuex',
//     'vue-router': 'VueRouter',
//     axios: 'axios'
//   },
//   // 通过cdn方式使用
//   js: [
//     'https://cdn.bootcss.com/vue/2.6.11/vue.runtime.min.js',
//     'https://cdn.bootcss.com/vue-router/3.1.2/vue-router.min.js',
//     'https://cdn.bootcss.com/vuex/3.1.2/vuex.min.js',
//     'https://cdn.bootcss.com/axios/0.19.2/axios.min.js',
//     'https://cdn.bootcss.com/moment.js/2.24.0/moment.min.js',
//     'https://cdn.bootcss.com/echarts/3.7.1/echarts.min.js'
//   ],
//   css: []
// }

module.exports = {
    publicPath: './',
    outputDir: './dist',
    assetsDir: './static', // 放置静态资源的目录
    indexPath: 'index.html', // index.html 的输出路径 (相对于 outputDir)
    filenameHashing: true, // 文件名hash 以便更好的控制缓存

    lintOnSave: true, // 是否在保存的时候eslint检查
    runtimeCompiler: false, // 是否使用包含运行时编译器的 Vue 构建版本
    productionSourceMap: false, // 生产环境是否生成 sourceMap 文件
    integrity: false,

    chainWebpack: (config) => {
        // // 移除prefetch插件，避免加载多余的资源
        // config.plugins.delete('prefetch')
        //     // 压缩图片
        // const imagesRule = config.module.rule('images')
        // imagesRule.uses.clear()
        // imagesRule.use('file-loader')
        //     .loader('url-loader')
        //     .options({
        //         limit: 10240,
        //         fallback: {
        //             loader: 'file-loader',
        //             options: {
        //                 outputPath: 'static/images'
        //             }
        //         }
        //     })

        // // 压缩响应的app.json返回的代码压缩
        // config.optimization.minimize(true)

    },
    configureWebpack: (config) => {
        if (process.env.NODE_ENV === 'production') {
            config.mode = 'production';
            config["performance"] = { //打包文件大小配置
                    "maxEntrypointSize": 10000000,
                    "maxAssetSize": 30000000
                }
                // config.plugins.push(
                //     new CompressionWebpackPlugin({
                //         filename: info => {
                //             return `${info.path}.gz${info.query}`
                //         },
                //         algorithm: 'gzip',
                //         threshold: 10240, // 只有大小大于该值的资源会被处理 10240
                //         test: new RegExp('\\.(' + ['js'].join('|') + ')$'),
                //         minRatio: 0.8, // 只有压缩率小于这个值的资源才会被处理
                //         deleteOriginalAssets: false // 删除原文件
                //     })
                // )
        } else {
            config.mode = 'development'
        }
    },
    css: {
        extract: true, // 是否使用css分离插件 ExtractTextPlugin
        sourceMap: false, // 开启 CSS source maps?
        loaderOptions: { // css预设器配置项
        },
        requireModuleExtension: true // 启用 CSS modules for all css / pre-processor files.
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
    pluginOptions: {}
}