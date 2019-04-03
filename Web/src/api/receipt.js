import api from '@/api/request.js'

const List = function (data, params) {
  return api({
    method: 'POST',
    url: '/web/receipt/list',
    data: data,
    params: params
  })
}

export default {
  List
}
