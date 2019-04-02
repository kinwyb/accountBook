'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')
const api_root = '"http://localhost:8080/v1"'

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  API_URL: api_root,
  Mock:true, //mock数据
})
