import api from '@/api/request.js'

const List = function (data) {
  return api({
    method: 'GET',
    url: '/micro/index/data',
    data: data
  })
}

export default {
  List
}
