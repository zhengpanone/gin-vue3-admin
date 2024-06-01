import request from '@/utils/request'
import {Admin, AdminPostData, IListParams} from './types/admin'

export const getAdmins = (params: IListParams) => {
  return request<{
    count: number
    list: Admin[]
  }>({
    method: 'GET',
    url: '/admin/getUserList',
    params
  })
}

export const createAdmin = (data: AdminPostData) => {
  return request({
    method: 'POST',
    url: '/setting/admin',
    data
  })
}

export const updateAdmin = (id: number, data: AdminPostData) => {
  return request({
    method: 'PUT',
    url: `/setting/admin/${id}`,
    data
  })
}

export const deleteAdmin = (id: number) => {
  return request({
    method: 'DELETE',
    url: `/setting/admin/${id}`
  })
}

export const updateAdminStatus = (id: number,status: number) => {
  return request({
    method: 'PUT',
    url: `/setting/admin/${id}/${status}`,
  })
}