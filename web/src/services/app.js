import { get, post } from 'utils'

export async function checkIsLoginService (params) {
  return get('sysUser/checkIsLogin', params)
}

export async function loginService (params) {
  return post('sysUser/login', params)
}

export async function logoutService (params) {
  return post('sysUser/logout', params)
}
