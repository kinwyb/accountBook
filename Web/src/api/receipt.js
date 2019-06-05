import api from '@/api/request.js'

const List = function (data, params) {
  return api({
    method: 'POST',
    url: '/web/receipt/list',
    data: data,
    params: params
  })
}

const LastNo = function () {
  return api({
    method: 'GET',
    url: '/web/receipt/nextNo'
  })
}

const Add = function (data) {
  return api({
    method: 'POST',
    url: '/web/receipt/add',
    data: data
  })
}

export default {
  List,
  LastNo,
  Add
}
