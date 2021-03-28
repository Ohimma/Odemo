import { createRouter, createWebHashHistory } from 'vue-router' 


import Public from './modules/public.js'
import Users from './modules/users.js'
import Elements from './modules/element.js'

// 方式一：框架内导入
let _routerArray = []
_routerArray = _routerArray.concat(Public.routers)
_routerArray = _routerArray.concat(Users.routers)
_routerArray = _routerArray.concat(Elements.routers)

// 方式二：非框架内导入
// routes = routes.concat(Public.routers)

let routes = [
  {
    path: '/',
    redirect: '/home',
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/layout/login1'),
    meta: { title: '登录', requiresAuth: false},
  },
  {
    path: '/lay',
    name: 'layout',
    component: () => import('@/layout/admin1.vue'),
    meta: { title: '后台',requiresAuth: true},
    children: _routerArray,
  }
]


const router = createRouter({
  history: createWebHashHistory(),
  routes: routes,
})

import store from '@/store/index.js'

router.beforeEach((to, from, next) => {
  console.log("enter route beforeEach ......")
  if (to.meta.requiresAuth) {
      document.title = to.meta.title
      console.log("document.title", document.title)

      if (store.state.layout.token) {
          next()
      } else {
          next({path: '/login'})
      }
  } else {
    next()
  }
})

export default router
