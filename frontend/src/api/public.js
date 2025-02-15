import http from './axios'

export const getVersion = () => {
  return http.request({
    url: '/sys/info/version',
    method: 'get'
  })
}

export const getVersionFull = () => {
  return http.request({
    url: '/sys/info/version/full',
    method: 'get'
  })
}

export const userLogin = (data) => {
  return http.request({
    url: '/user/login',
    method: 'post',
    data
  })
}
