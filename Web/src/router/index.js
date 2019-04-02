import Vue from 'vue'
import Router from 'vue-router'
import Layout from '@/views/layout/index'

const _import = require('./autoImport')

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      redirect: '/receipt/list/0'
    },
    {
      path: '/receipt',
      name: '收支数据',
      component: Layout,
      redirect: '/receipt/list',
      children: [{
        path: 'list/0',
        name: '收支列表',
        component: _import('receipt/list')
      }, {
        path: 'list/1',
        name: '收入列表',
        component: _import('receipt/list')
      }, {
        path: 'list/2',
        name: '支出列表',
        component: _import('receipt/list')
      }]
    }
  ]
})
