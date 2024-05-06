import axios, {AxiosRequestConfig} from 'axios'
import {ElMessage, ElMessageBox} from 'element-plus'
import {getItem, removeItem} from "@/utils/storage";
import {IUserInfo} from "@/api/types/common";
import {USER} from "@/utils/constants";
import router from "@/routers";

const request = axios.create({
  baseURL: import.meta.env.VITE_API_BASEURL

})
// 控制登录过期的锁
let isRefreshing = false
// 请求拦截器
request.interceptors.request.use((config) => {
  // 统一设置用户token
  const userInfo = getItem<{ token: string } & IUserInfo>(USER)
  if (userInfo && userInfo.token) {
    if (!config.headers) {
      config.headers = {}
    }
    config.headers.Authorization = `Bear ${userInfo.token}`
  }
  // 统一设置用户身份 token
  return config
}, (error) => {
  return Promise.reject(error)
})
// 响应拦截器
request.interceptors.response.use((response) => {

  if(!response.data.code || response.data.code===200){
    return response
  }
  // 统一处理接口响应错误 如 token过期,服务端异常
  // token失效
  if (response.data.code == -2) {
    if(isRefreshing) return Promise.reject(response)
    isRefreshing = true
    // 清除本地过期token
    // 跳转到登录页面
    // 抛出异常
    ElMessageBox.confirm('您的登录已经过期，您可以取消停留在此页面，或确认重新登录', '登录过期', {
      confirmButtonText: '确认',
      cancelButtonClass: '取消'
    }).then(() => {
      removeItem(USER)
      router.push({
        name: 'login',
        query: {
          redirect: router.currentRoute.value.fullPath
        }
      })
    }).finally(() => {
      isRefreshing = false
    })
    return Promise.reject(response)

  }
  // 其他错误情况
  ElMessage.error(response.data.msg || '请求失败，请稍后重试')
  // 手动返回promise异常
  return Promise.reject(response.data)


}, (error) => {
  return Promise.reject(error)
})

interface ResponseData<T = any> {
  code: number,
  msg: string,
  data: T,
  meta: string,
  errors: [{ key: string, error: string }],
}
// request 不支持泛型
// request.get post put支持响应数据泛型
// 后端封装了一层数据,导致访问比较麻烦,所以自己封装了一个request
export default <T = any>(config: AxiosRequestConfig) => {
  return request(config).then(res => {
    return res.data as ResponseData<T>
  })
}
