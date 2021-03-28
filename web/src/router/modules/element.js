const element = {}

element.routers = [
  {
    path: '/element',
    component: () => import('@/views/element'),
    meta: { title: '饿了么'},
    children: [
      {
        path: 'images',
        name: 'images',
        component: () => import('@/views/element/images.vue'),
        meta: { title: '图片' }
      },
      {
        path: 'steps',
        name: 'steps',
        component: () => import('@/views/element/steps.vue'),
        meta: { title: '步骤条' }
      },
    ]
  }
]

export default element
