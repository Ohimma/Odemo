// const NotFoundPage = resolve => require(['../../views/errors/404.vue'], resolve)
// const NotPowerPage = resolve => require(['../../views/errors/401.vue'], resolve)

const Errors = {}

Errors.routers = [
  {
    path: '/404',
    component: () => import('@/views/errors/404.vue'),
    meta: { title: '404', icon: 'role', requiresAuth: false}
  },
  {
    path: '/401',
    component: () => import('@/views/errors/401.vue'),
    hidden: true,
    meta: { title: '401', icon: 'role', requiresAuth: false}
  },
  {
    path: '/:catchAll(.*)',
    hidden: false,
    redirect: { path: '/404' },
  }
]

export default Errors
