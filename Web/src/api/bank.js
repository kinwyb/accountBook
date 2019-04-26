import api from '@/api/request.js'

// 银行列表
const BankList = function (params) {
  return api({
    method: 'GET',
    url: '/web/bank/list',
    params: params
  })
}

// 银行余额计算列表
const BankListCompate = function (params) {
  return api({
    method: 'GET',
    url: '/web/bank/list/compute/day',
    params: params
  })
}

// 新增银行
const BankAdd = function (data) {
  return api({
    method: 'POST',
    url: '/web/bank/add',
    data: data
  })
}

export default {
  BankList,
  BankListCompate,
  BankAdd
}
