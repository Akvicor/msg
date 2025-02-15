import http from "./axios";


export const channelCreate = (data) => {
  return http.request({
    url: '/channel/create',
    method: 'post',
    data
  })
}

export const channelFind = (params) => {
  return http.request({
    url: '/channel/find',
    method: 'get',
    params
  })
}

export const channelUpdate = (data) => {
  return http.request({
    url: '/channel/update',
    method: 'post',
    data
  })
}

export const channelDelete = (data) => {
  return http.request({
    url: '/channel/delete',
    method: 'post',
    data
  })
}

export const channelTest = (params) => {
  return http.request({
    url: '/channel/test',
    method: 'get',
    params
  })
}

export const channelSend = (data) => {
  return http.request({
    url: '/channel/send',
    method: 'post',
    data
  })
}
