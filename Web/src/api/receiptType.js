import api from '@/api/request.js'

const List = function (params) {
  return api({
    method: 'GET',
    url: '/web/receiptType/list',
    params: params
  })
}

const ListByLevel = function (params) {
  return api({
    method: 'GET',
    url: '/web/receiptType/list/level',
    params: params
  })
}

export default {
  List,
  ListByLevel
}
