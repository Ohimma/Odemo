const routes = [
  {
    path: '/product',
    component: () => import('@/views/product'),
    meta: { title: '产品管理', icon: 'product'},
    children: [
      {
        path: 'test',
        name: 'test',
        component: () => import('@/views/product/test.vue'),
        meta: { title: '测试', icon: 'test' }
      },
    
    ]
  }
]

export default {
  routers: routes
}