import http from "./axios";

export const typeRoleTypeFind = (params) => {
  return http.request({
    url: '/type/role/type',
    method: 'get',
    params
  })
}

export const typeChannelTypeFind = (params) => {
  return http.request({
    url: '/type/channel/type',
    method: 'get',
    params
  })
}

export const typeSendTypeFind = (params) => {
  return http.request({
    url: '/type/send/type',
    method: 'get',
    params
  })
}

export const typePeriodTypeFind = (params) => {
  return http.request({
    url: '/type/period/type',
    method: 'get',
    params
  })
}

export const typeBotSendersFind = (params) => {
  return http.request({
    url: '/type/bot/senders',
    method: 'get',
    params
  })
}