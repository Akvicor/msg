import axios from 'axios'

let baseUrl

if (process.env.NODE_ENV === 'development') {
  baseUrl = 'http://127.0.0.1:3000/api'
} else if (process.env.NODE_ENV === 'production') {
  baseUrl = '/api'
} else {
  baseUrl = '/api'
}

const getToken = () => {
  const token = localStorage.getItem('login-token')
  if (token) {
    return token
  }
  return ''
}

class HttpRequest {
  constructor(baseUrl) {
    this.baseURL = baseUrl
  }

  getInsideConfig() {
    return {
      baseURL: this.baseURL,
      headers: {
        'X-Auth-Token': getToken()
      }
    }
  }

  interception(instance) {
    // 请求拦截器
    instance.interceptors.request.use(function (config) {
      // 在发请求之前做些什么
      return config
    }, function (error) {
      // 对请求错误做些什么
      return Promise.reject(error)
    });
    // 响应拦截器
    instance.interceptors.response.use(function (response) {
      // 2xx 范围内的状态码都会触发该函数
      // 对响应数据做些什么
      return response
    }, function (error) {
      // 超出2xx范围的状态码都会触发该函数
      // 对响应错误做些什么
      return Promise.reject(error)
    });

  }

  request(options) {
    options = {...this.getInsideConfig(), ...options}
    // 创建axios的实例
    const instance = axios.create()
    // 实例拦截器的绑定
    this.interception(instance)
    return instance(options)
  }
}

const Axios = new HttpRequest(baseUrl)

export default Axios
