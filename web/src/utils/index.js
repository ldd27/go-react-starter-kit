import classnames from 'classnames'
import lodash from 'lodash'
import { message } from 'antd'
import moment from 'moment'
import cookie from './cookie'
import config from './config'
import request from './request'
import { model, comModel, comPageModel, genComModel, genComPageModel, genCRUDComModel, genCRUDComPageModel } from './base_model'

const { prefix } = config

/**
 * 数组内查询
 * @param   {array}      array
 * @param   {String}    id
 * @param   {String}    keyAlias
 * @return  {Array}
 */
const queryArray = (array, key, keyAlias = 'key') => {
  if (!(array instanceof Array)) {
    return null
  }
  const item = array.filter(_ => _[keyAlias] === key)
  if (item.length) {
    return item[0]
  }
  return null
}

/**
 * 数组格式转树状结构
 * @param   {array}     array
 * @param   {String}    id
 * @param   {String}    pid
 * @param   {String}    children
 * @return  {Array}
 */
const arrayToTree = (array, id = 'id', pid = 'pid', children = 'children') => {
  let data = lodash.cloneDeep(array)
  let result = []
  let hash = {}
  data.forEach((item, index) => {
    hash[data[index][id]] = data[index]
  })

  data.forEach((item) => {
    let hashVP = hash[item[pid]]
    if (hashVP) {
      !hashVP[children] && (hashVP[children] = [])
      hashVP[children].push(item)
    } else {
      result.push(item)
    }
  })
  return result
}

const get = (url, params) => {
  return request({
    url,
    method: 'get',
    data: params,
    headers: { Authorization: cookie.getCookie(`${prefix}token`) },
  })
}

const post = (url, params) => {
  return request({
    url,
    method: 'post',
    data: params,
    headers: { Authorization: cookie.getCookie(`${prefix}token`) },
  })
}

const put = (url, params) => {
  return request({
    url,
    method: 'put',
    data: params,
    headers: { Authorization: cookie.getCookie(`${prefix}token`) },
  })
}

const remove = (url, params) => {
  return request({
    url,
    method: 'delete',
    data: params,
    headers: { Authorization: cookie.getCookie(`${prefix}token`) },
  })
}

const checkApiRs = (rs) => {
  switch (rs.code) {
    case 1003:
      cookie.delCookie(`${prefix}username`)
      cookie.delCookie(`${prefix}token`)
      cookie.delCookie(`${prefix}menu`)
      if (location.pathname !== '/login') {
        let from = location.pathname
        if (location.pathname === '/home') {
          from = '/home'
        }
        window.location = `${location.origin}/login?from=${from}`
      }
      break
    case 1002:
      message.warn('非法参数，请重新输入', 3)
      break
    case 2001:
      message.warn('用户名密码错误，请重新输入', 3)
      break
    case 2002:
      message.warn('原密码错误，请重新输入', 3)
      break
    default:
      message.error('服务器繁忙', 3)
      break
  }
}

/**
 * @param   {String}
 * @return  {String}
 */

const queryURL = (name) => {
  let reg = new RegExp(`(^|&)${name}=([^&]*)(&|$)`, 'i')
  let r = window.location.search.substr(1).match(reg)
  if (r != null) return decodeURI(r[2])
  return null
}

/**
 * format
 * 时间转化
 * @export
 * @param {any} origin
 * @param {string} [format='YYYY-MM-DD HH:mm:ss']
 * @param {any} originFormat 原始时间格式
 * @returns
 */
export function formatDate (origin, format = 'YYYY-MM-DD HH:mm:ss', originFormat) {
  if (!origin || origin.indexOf('0001-01-01') !== -1) return ''
  let temp
  if (originFormat) temp = moment(origin, originFormat)
  else temp = moment(origin)
  if (temp.isValid()) return temp.format(format)
  return ''
}

module.exports = {
  config,
  request,
  classnames,
  queryArray,
  arrayToTree,
  get,
  post,
  put,
  remove,
  checkApiRs,
  cookie,
  queryURL,
  formatDate,
  model,
  comModel,
  comPageModel,
  genComModel,
  genComPageModel,
  genCRUDComModel,
  genCRUDComPageModel,
}
