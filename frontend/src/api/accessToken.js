import http from "./axios";

export const accessTokenFind = (params) => {
  return http.request({
    url: '/user/access_token/find',
    method: 'get',
    params
  })
}

export const accessTokenCreate = (data) => {
  return http.request({
    url: '/user/access_token/create',
    method: 'post',
    data
  })
}

export const accessTokenUpdate = (data) => {
  return http.request({
    url: '/user/access_token/update',
    method: 'post',
    data
  })
}

export const accessTokenDelete = (data) => {
  return http.request({
    url: '/user/access_token/delete',
    method: 'post',
    data
  })
}
