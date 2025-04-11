import http from "./axios";

export const scheduleFind = (params) => {
  return http.request({
    url: '/schedule/find',
    method: 'get',
    params
  })
}

export const scheduleCreate = (data) => {
  return http.request({
    url: '/schedule/create',
    method: 'post',
    data
  })
}

export const scheduleUpdate = (data) => {
  return http.request({
    url: '/schedule/update',
    method: 'post',
    data
  })
}

export const scheduleUpdateNext = (data) => {
  return http.request({
    url: '/schedule/update/next',
    method: 'post',
    data
  })
}

export const scheduleUpdateSequence = (data) => {
  return http.request({
    url: '/schedule/update/sequence',
    method: 'post',
    data
  })
}

export const scheduleDisable = (data) => {
  return http.request({
    url: '/schedule/disable',
    method: 'post',
    data
  })
}

export const scheduleEnable = (data) => {
  return http.request({
    url: '/schedule/enable',
    method: 'post',
    data
  })
}

export const scheduleDelete = (data) => {
  return http.request({
    url: '/schedule/delete',
    method: 'post',
    data
  })
}
