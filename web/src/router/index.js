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
  if (store.state.layout.token) {
      next()
  } else {
      if (to.path !== '/login') {
          next({path: '/login'})
      } else {
          next()
      }
  }

  document.title = to.meta.title
})

export default router
