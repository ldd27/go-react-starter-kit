import { get, post } from '../utils'

export async function checkIsLoginService (params) {
  return get('sysUser/checkIsLogin')
}

export async function loginService (params) {
  return post('sysUser/login')
}

export async function logoutService (params) {
  return post('sysUser/logout')
}