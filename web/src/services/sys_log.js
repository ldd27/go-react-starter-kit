import { get } from 'utils'

export async function getPageSvc (params) {
  return get('sysLog/page', params)
}
