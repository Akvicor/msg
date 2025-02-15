import http from "./axios";


export const sendSend = (data) => {
  return http.request({
    url: '/send',
    method: 'post',
    data
  })
}

export const sendFind = (data) => {
  return http.request({
    url: '/send/find',
    method: 'post',
    data
  })
}

export const sendCancel = (data) => {
  return http.request({
    url: '/send/cancel',
    method: 'post',
    data
  })
}

export const sendStatus = (data) => {
  return http.request({
    url: '/send/status',
    method: 'post',
    data
  })
}