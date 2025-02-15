import http from "./axios";


export const userCreate = (data) => {
  return http.request({
    url: '/admin/user/create',
    method: 'post',
    data
  })
}

export const userFind = (params) => {
  return http.request({
    url: '/admin/user/find',
    method: 'get',
    params
  })
}

export const userUpdate = (data) => {
  return http.request({
    url: '/admin/user/update',
    method: 'post',
    data
  })
}

export const userDisable = (data) => {
  return http.request({
    url: '/admin/user/disable',
    method: 'post',
    data
  })
}

export const userEnable = (data) => {
  return http.request({
    url: '/admin/user/enable',
    method: 'post',
    data
  })
}
