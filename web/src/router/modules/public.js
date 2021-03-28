
const routes = [
  {
    path: '/home',
    name: 'Home',
    component: () => import('@/views/public/home.vue'),
    meta: {title: '首页', icon: 'home', requiresAuth: true},
  },
  {
    path: '/404',
    name: '404',
    component: () => import('@/views/public/404.vue'),
    meta: {title: '404', icon: '404', requiresAuth: true},
  },
  {
    path: '/:catchAll(.*)',
    hidden: false,
    component: () => import('@/views/public/404.vue'),
  },
]  

export default {
  routers: routes
}