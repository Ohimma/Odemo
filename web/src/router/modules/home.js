
const routes = [
  {
    path: '/home',
    name: 'Home',
    component: () => import('@/views/home.vue'),
    meta: {title: '首页', icon: 'home', requiresAuth: true},
  }
]  

export default {
  routers: routes
}