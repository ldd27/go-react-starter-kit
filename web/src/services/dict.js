import { get, post, put, remove } from 'utils'

export async function getIndexTreeSvc (data) {
  return get('dictIndex/tree', data)
}

export async function getSvc (data) {
  return get('dictItem', data)
}

export async function addSvc (data) {
  return post('dictItem', data)
}

export async function uptSvc (data) {
  return put('dictItem', data)
}

export async function delSvc (data) {
  return remove('dictItem', data)
}
