const Provider = {}

Provider.routers = [
  {
    path: '/provider',
    component: () => import('@/views/provider'),
    meta: { title: '供应商管理', icon: 'provider'},
    children: [
      {
        path: 'define',
        name: 'define',
        component: () => import('@/views/provider/define.vue'),
        meta: { title: '商户管理', icon: 'define' }
      },
      {
        path: 'account',
        name: 'account',
        component: () => import('@/views/provider/account.vue'),
        meta: { title: '账户管理', icon: 'account' }
      },
    ]
  }
]

export default Provider
