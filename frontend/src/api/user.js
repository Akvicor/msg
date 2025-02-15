import http from "./axios";

export const userInfo = () => {
  return http.request({
    url: '/user/info',
    method: 'get'
  })
}

export const userUpdate = (data) => {
  return http.request({
    url: '/user/update',
    method: 'post',
    data
  })
}

export const userLogout = (data) => {
  return http.request({
    url: '/user/logout',
    method: 'post',
    data
  })
}
