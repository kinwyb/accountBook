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

const Tree = function () {
  return api({
    method: 'GET',
    url: '/web/receiptType/tree'
  })
}

const Add = function (data) {
  return api({
    method: 'POST',
    url: '/web/receiptType/add',
    data: data
  })
}

export default {
  List,
  ListByLevel,
  Tree,
  Add
}
