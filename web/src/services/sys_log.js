import { get } from '../utils'

export async function getPagingService (params) {
  return get('sysLog/page', params)
}
