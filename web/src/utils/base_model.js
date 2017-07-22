import modelExtend from 'dva-model-extend'
import { message } from 'antd'
import cookie from './cookie'
import config from './config'

const { prefix } = config
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

const model = {
  reducers: {
    updateState (state, { payload }) {
      return { ...state, ...payload }
    },
  },
}

const comModel = modelExtend(model, {
  state: {
    data: [],
    search: {},
    currentItem: {},
    modalVisible: false,
    modalType: 'create',
  },
  reducers: {
    success (state, { payload }) {
      return {
        ...state,
        ...payload,
      }
    },
    showModal (state, { payload }) {
      return { ...state, ...payload, modalVisible: true }
    },
    hideModal (state, { payload }) {
      return { ...state, ...payload, modalVisible: false }
    },
  },
})

const comPageModel = modelExtend(comModel, {
  state: {
    pagination: {
      showSizeChanger: false,
      showQuickJumper: false,
      showTotal: total => `总共 ${total} 条数据`,
      current: 1,
      pageSize: 10,
      total: 0,
    },
  },
  reducers: {
    success (state, { payload }) {
      const { pagination } = payload
      return {
        ...state,
        ...payload,
        pagination: {
          ...state.pagination,
          ...pagination,
        },
      }
    },
  },
})

function genComModel (namespace, service, state = {}, effects, setup, reducers) {
  return modelExtend(comModel, {
    namespace,
    state,
    subscriptions: {
      setup ({ dispatch, history }) {
        if (!setup) {
          history.listen((location) => {
            if (location.pathname === `/${namespace}`) {
              dispatch({ type: 'get' })
            }
          })
        } else {
          setup({ dispatch, history })
        }
      },
    },
    effects: {
      * get ({ payload = {} }, { call, put }) {
        const data = yield call(service.getSvc, { ...payload.search })
        if (data.success) {
          yield put({ type: 'success',
            payload: { data: data.r,
              search: payload.search,
            },
          })
        } else {
          checkApiRs(data)
        }
      },
      ...effects,
    },
    reducers,
  })
}

function genComPageModel (namespace, service, state = {}, effects, setup, reducers) {
  return modelExtend(comPageModel, {
    namespace,
    state,
    subscriptions: {
      setup ({ dispatch, history }) {
        if (!setup) {
          history.listen((location) => {
            if (location.pathname === `/${namespace}`) {
              dispatch({ type: 'getPage' })
            }
          })
        } else {
          setup({ dispatch, history })
        }
      },
    },
    effects: {
      * getPage ({ payload = { current: 1, pageSize: 10 } }, { call, put }) {
        const data = yield call(service.getPageSvc, { page: payload.current, size: payload.pageSize, ...payload.search })
        if (data.success) {
          yield put({ type: 'success',
            payload: { data: data.r.data,
              search: payload.search,
              pagination: {
                total: data.r.total,
                current: payload.current,
                pageSize: payload.pageSize,
              },
            },
          })
        } else {
          checkApiRs(data)
        }
      },
      ...effects,
    },
    reducers,
  })
}

function genCRUDComModel (namespace, service, state = {}, effects, setup, reducers) {
  return modelExtend(genComModel(namespace, service, state, effects, setup, reducers), {
    namespace,
    effects: {
      * create ({ payload }, { call, put }) {
        const data = yield call(service.addSvc, payload.data)
        if (data.success) {
          message.success('保存成功', 3)
          yield put({ type: 'hideModal' })
          yield put({ type: 'get' })
        } else {
          checkApiRs(data)
        }
      },
      * update ({ payload }, { call, put }) {
        const data = yield call(service.uptSvc, payload.data)
        if (data.success) {
          message.success('保存成功', 3)
          yield put({ type: 'hideModal' })
          yield put({ type: 'get' })
        } else {
          checkApiRs(data)
        }
      },
      * remove ({ payload }, { call, put }) {
        const data = yield call(service.delSvc, payload)
        if (data.success) {
          message.success('删除成功', 3)
          yield put({ type: 'get' })
        } else {
          checkApiRs(data)
        }
      },
    },
    reducers,
  })
}

function genCRUDComPageModel (namespace, service, state = {}, effects, setup, reducers) {
  return modelExtend(genComPageModel(namespace, service, state, effects, setup, reducers), {
    namespace,
    effects: {
      * create ({ payload }, { call, put }) {
        const data = yield call(service.addSvc, payload.data)
        if (data.success) {
          message.success('保存成功', 3)
          yield put({ type: 'hideModal' })
          yield put({ type: 'getPage' })
        } else {
          checkApiRs(data)
        }
      },
      * update ({ payload }, { call, put }) {
        const data = yield call(service.uptSvc, payload.data)
        if (data.success) {
          message.success('保存成功', 3)
          yield put({ type: 'hideModal' })
          yield put({ type: 'getPage' })
        } else {
          checkApiRs(data)
        }
      },
      * remove ({ payload }, { call, put }) {
        const data = yield call(service.delSvc, payload)
        if (data.success) {
          message.success('删除成功', 3)
          yield put({ type: 'getPage' })
        } else {
          checkApiRs(data)
        }
      },
    },
    reducers,
  })
}

export { model, comModel, comPageModel, genComModel, genComPageModel, genCRUDComModel, genCRUDComPageModel }
