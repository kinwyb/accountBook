import axios from 'axios'

const api = axios.create({
  baseURL: process.env.API_URL,
  timeout: 180000
})

// 请求拦截
api.interceptors.request.use(
  config => {
    return config
  }, error => {
    Promise.reject(error)
  })

// 返回结果拦截
api.interceptors.response.use(
  response => {
    return response.data
  },
  error => {
    console.log('err:' + JSON.stringify(error)) // debug
  }
)

export default api
