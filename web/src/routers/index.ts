import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'
import AppLayout from '@/layout/AppLayout.vue'
import productRoutes from './modules/product'
import nprogress from 'nprogress'
import 'nprogress/nprogress.css'
import { indexStore } from '@/store/index'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: AppLayout,
    children: [
      {
        path: '', // 默认子路由
        name: 'home',
        component: () => import('../views/home/index.vue'),
        meta: {
          title: '首页',
          requireAuth: true,
        }
      },
      productRoutes
    ]

  },
  {
    path: '/login',
    name: 'login',
    component: () => import('../views/login/index.vue')
  }
]

const router = createRouter({
  history: createWebHashHistory(), // 路由模式
  routes// 路由规则
})
// 导航守卫

// 全局前置守卫
router.beforeEach((to, from) => {
  const store = indexStore()
  if (to.meta.requireAuth && !store.$state.user) {
    return {
      path: '/login',
      query: { redirect: to.fullPath }
    }
  }
  nprogress.start() // 页面加载进度条
})
// 全局后置守卫
router.afterEach(() => {
  nprogress.done()
})

export default router
