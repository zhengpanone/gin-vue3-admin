import { RouteRecordRaw, RouterView } from 'vue-router'
const routes: RouteRecordRaw =
  {
    path: 'product',
    name: 'product',
    component: RouterView,
    meta: {
      title: '商品',
      requireAuth: true,
    },
    children: [
      {
        path: 'product-list',
        name: 'product-list',
        component: () => import('@/views/product/list/index.vue'),
        meta: {
          title: '商品列表'
        }
      },
      {
        path: 'product-attr',
        name: 'product-attr',
        component: () => import('@/views/product/attr/index.vue'),
        meta: {
          title: '商品规格'
        }
      },
      {
        path: 'product-classify',
        name: 'product-classify',
        component: () => import('@/views/product/classify/index.vue'),
        meta: {
          title: '商品类别'
        }
      },
      {
        path: 'product-reply',
        name: 'product-reply',
        component: () => import('@/views/product/reply/index.vue'),
        meta: {
          title: '商品评价'
        }
      }
    ]
  }

export default routes
