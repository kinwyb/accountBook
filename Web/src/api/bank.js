import api from '@/api/request.js'

const BankList = function (params) {
  return api({
    method: 'GET',
    url: '/web/bank/list',
    params: params
  })
}

export default {
  BankList
}
