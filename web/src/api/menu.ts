import request from '@/utils/request'

// @Summary 用户登录 获取动态路由
// @Produce  application/json
// @Param 可以什么都不填 调一下即可
// @Router /api/menu/getMenu [post]
export const getBaseMenuTree = () => {
  return request({
    method: 'POST',
    url: '/api/menu/getMenu'
  })
}
