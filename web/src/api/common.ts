/**
 * 公共基础接口
 */
import request from '@/utils/request'
import type {ICaptchaInfo, ILoginResponse} from './types/common'
import {Md5} from 'ts-md5'

export const getCaptcha = () => {
    return request<ICaptchaInfo>({
        method: 'GET',
        url: '/api/captcha',
        params: {stamp: Date.now()},
        responseType: 'json'
    })
}

export const login = (data: {
    username: string,
    password: string,
    captcha: string,
    captchaId: string
}) => {

        const md5 = new Md5()
        md5.appendStr(data.password)
        const passwordMD5 = md5.end() as string

    return request<ILoginResponse>({
        method: 'POST',
        url: '/api/admin/login',
        data: {
            ...data,
            password: passwordMD5
        }
    })
}

export const logout = (token: string | undefined) => {
    return request({
        method: 'POST',
        url: '/api/jwt/jwtInBlacklist',
        headers: token ? {Authorization: token} : undefined
    })
}
