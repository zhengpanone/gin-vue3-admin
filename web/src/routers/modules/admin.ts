import { RouteRecordRaw, RouterView } from 'vue-router'
const routes: RouteRecordRaw =
  {
    path: 'system',
    name: 'system',
    component: RouterView,
    meta: {
      title: '系统管理',
      requireAuth: true,
    },
    children: [
      {
        path: 'user-list',
        name: 'user-list',
        component: () => import('@/views/admin/index.vue'),
        meta: {
          title: '用户列表'
        }
      }
    ]
  }

export default routes
