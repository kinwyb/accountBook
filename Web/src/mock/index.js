import Mock from 'mockjs'

import receipt from './receipt'

// 请求的url最好用正则表达式,以防止匹配失败

Mock.mock(/\/v1\/micro\/index\/data/, receipt.List)

export default Mock
