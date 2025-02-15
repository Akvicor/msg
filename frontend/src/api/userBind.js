import http from "./axios";

export const userBindHomeTipsFind = (data) => {
  return http.request({
    url: '/user/bind/home/tips/find',
    method: 'get',
    data
  })
}

export const userBindHomeTipsSave = (data) => {
  return http.request({
    url: '/user/bind/home/tips/save',
    method: 'post',
    data
  })
}
