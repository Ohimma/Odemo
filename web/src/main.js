import { createApp } from 'vue'
import App from './App.vue'
import router from './router/index.js'
import store from './store/index.js'

// 导入公共样式css + js
import '@/assets/css/reset.css'
import '@/assets/css/base.css'

const app = createApp(App)


// 挂载原型
// 引入 elment
import ElementPlus from 'element-plus';
import 'element-plus/lib/theme-chalk/index.css';
app.use(ElementPlus)

// 引入 echarts
import * as echarts from 'echarts'
app.config.globalProperties.$echarts = echarts


// 引入http
import http from '@/utils/http.js'
app.config.globalProperties.$http = http
import utils from '@/utils/utils.js' 
app.config.globalProperties.$utils = utils
import Conver from '@/utils/conver.js' 
app.config.globalProperties.$Conver = Conver

app.use(store)
app.use(router)

app.mount("#app")


// console.log("app = ", app)
