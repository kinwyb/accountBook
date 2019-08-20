import api from '@/api/request.js'

// 银行列表
const LogList = function (body, params) {
  return api({
    method: 'POST',
    url: '/web/system/logList',
    data: body,
    params: params
  })
}

export default {
  LogList
}
