/**
 * 公共基础接口
 */
import request from '@/utils/request'
import type{ ICaptchaInfo, ILoginResponse } from './types/common'
export const getCaptcha = () => {
  return request<ICaptchaInfo>({
    method: 'GET',
    url: '/api/captcha',
    params: { stamp: Date.now() },
    responseType: 'json'
  })
}

export const login = (data: {
  username: string,
  password: string,
  captcha: string,
  captchaId: string
}) => {
  return request<ILoginResponse>({
    method: 'POST',
    url: '/api/admin/login',
    data
  })
}

export const logout = (token:string) => {
  return request({
    method: 'POST',
    url: '/api/jwt/jwtInBlacklist',
    headers: {'Authorization': token}
  })
}
