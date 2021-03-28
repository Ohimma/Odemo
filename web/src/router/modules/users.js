
const routes = [
  {
    path: '/user',
    component: () => import('@/views/users/index.vue'),
    name: 'user',
    meta: { title: '用户管理', icon: 'user', requiresAuth: true},
    children: [
      {
        path: 'user',
        name: 'user',
        component: () => import('@/views/users/user.vue'),
        meta: { title: '用户', icon: 'user', requiresAuth: true}
      },
      {
        path: 'role',
        name: 'role',
        component: () => import('@/views/users/role.vue'),
        meta: { title: '角色', icon: 'role', requiresAuth: true}
      },
      {
        path: 'auth',
        name: 'auth',
        component: () => import('@/views/users/auth.vue'),
        meta: { title: '权限因子', icon: 'auth', requiresAuth: true}
      },
    ]
  }
]  

export default {
  routers: routes
}