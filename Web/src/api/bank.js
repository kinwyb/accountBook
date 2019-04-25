import api from '@/api/request.js'

const BankList = function (params) {
  return api({
    method: 'GET',
    url: '/web/bank/list',
    params: params
  })
}

const BankListCompate = function (params) {
  return api({
    method: 'GET',
    url: '/web/bank/list/compute/day',
    params: params
  })
}

export default {
  BankList,
  BankListCompate
}
