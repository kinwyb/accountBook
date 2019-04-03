import Mock from 'mockjs'

// import receipt from './receipt'
import bank from './bank'

// 请求的url最好用正则表达式,以防止匹配失败

// Mock.mock(/\/v1\/micro\/index\/data/, receipt.List)
Mock.mock(/\/v1\/web\/bank\/list/, bank.List)

export default Mock
