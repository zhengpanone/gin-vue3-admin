import axios, { AxiosRequestConfig } from 'axios'

const request = axios.create({
  // baseURL: import.meta.env.VITE_API_BASEURL

})
// 请求拦截器
request.interceptors.request.use(function (config) {
  // 统一设置用户身份 token
  return config
}, function (error) {
  return Promise.reject(error)
})
// 响应拦截器
request.interceptors.response.use(function (response) {
  // 统一处理接口响应错误 如 token过期,服务端异常
  return response
}, function (error) {
  return Promise.reject(error)
})

interface ResponseData<T=any>{
  code: number,
  msg: string,
  data: T,
  meta: string,
  errors: [{ key: string, error: string }],
}
// request 不支持泛型
// request.get post put支持响应数据泛型
// 后端封装了一层数据,导致访问比较麻烦,所以自己封装了一个request
export default <T=any>(config: AxiosRequestConfig) => {
  return request(config).then(res => {
    return res.data as ResponseData<T>
  })
}
