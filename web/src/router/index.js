import { createRouter, createWebHashHistory } from 'vue-router' 

// 导入一级路由
import Errors from './modules/errors.js'

// 导入带有UI框架的路由
import Home from './modules/home.js'
import Users from './modules/users.js'
import Product from './modules/product.js'
import Provider from './modules/provider.js'

let _routerArray = []
_routerArray = _routerArray.concat(Home.routers)
_routerArray = _routerArray.concat(Users.routers)
_routerArray = _routerArray.concat(Product.routers)
_routerArray = _routerArray.concat(Provider.routers)


let routes = [
  {
    path: '/test',
    component: () => import('@/views/home.vue'),
    meta: { title: '测试', icon: 'test', requiresAuth: false, },
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/layout/login1'),
    meta: { title: '登录', icon: 'login', requiresAuth: false},
  },
  {
    path: '/',
    redirect: 'Home',
  },
  {
    path: '/lay',
    name: 'layout',
    component: () => import('@/layout/admin1.vue'),
    meta: { title: '后台', icon: 'layout', requiresAuth: true},
    children: _routerArray,
  }
]

routes = routes.concat(Errors.routers)


const router = createRouter({
  history: createWebHashHistory(),
  routes: routes,
})

import store from '@/store/index.js'

router.beforeEach((to, from, next) => {
  console.log("enter route beforeEach to", to)
  if (to.meta.requiresAuth) {
      document.title = to.meta.title
      console.log("document.title", document.title)

      if (store.state.login.token) {
          next()
      } else {
          next({path: '/login'})
      }
  } else {
    next()
  }
})

export default router
