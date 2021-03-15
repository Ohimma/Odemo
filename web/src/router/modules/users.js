
const routes = [
  {
    path: '/user',
    component: () => import('@/views/users/index.vue'),
    name: 'user',
    meta: { title: '用户管理', icon: 'user', requiresAuth: true},
    children: [
      {
        path: 'role',
        name: 'role',
        component: () => import('@/views/users/role.vue'),
        meta: { title: '角色', icon: 'role', requiresAuth: true}
      },
      {
        path: 'user',
        name: 'user',
        component: () => import('@/views/users/user.vue'),
        meta: { title: '用户', icon: 'user', requiresAuth: true}
      },
      {
        path: 'acl',
        name: 'acl',
        component: () => import('@/views/users/acl.vue'),
        meta: { title: '资源', icon: 'acl', requiresAuth: true}
      },
    ]
  }
]  

export default {
  routers: routes
}