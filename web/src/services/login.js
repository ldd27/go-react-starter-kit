import { post } from 'utils'

export async function loginService (data) {
  return post('sysUser/login', data)
}
