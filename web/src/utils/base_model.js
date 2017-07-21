import modelExtend from 'dva-model-extend'
import { message } from 'antd'

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
      showSizeChanger: true,
      showQuickJumper: true,
      showTotal: total => `总共 ${total} 条数据`,
      current: 1,
      pageSize: 10,
      total: 0,
    },
  },
  reducers: {
    success (state, { payload }) {
      const { data, pagination } = payload
      return {
        ...state,
        data,
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
          setup(dispatch, history)
        }
      },
    },
    effects: {
      * get ({ payload = {} }, { call, put }) {
        const data = yield call(service.getSvc, { ...payload.search })
        if (data) {
          yield put({ type: 'success',
            payload: { data: data.data,
              search: payload.search,
            },
          })
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
          setup(dispatch, history)
        }
      },
    },
    effects: {
      * getPage ({ payload = { current: 1, pageSize: 10 } }, { call, put }) {
        const data = yield call(service.getPageSvc, { page: payload.current, size: payload.pageSize, ...payload.search })
        if (data) {
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
        }
      },
      ...effects,
    },
    reducers,
  })
}

function genCRUDComModel (namespace, service, state = {}, effects, setup, reducers) {
  return modelExtend(genComModel(namespace, service, state, effects, setup, reducers), {
    effects: {
      * create ({ payload }, { call, put }) {
        const data = yield call(service.addSvc, payload.data)
        if (data) {
          message.success('保存成功', 3)
          yield put({ type: 'hideModal' })
          yield put({ type: 'get' })
        }
      },
      * update ({ payload }, { call, put }) {
        const data = yield call(service.uptSvc, payload.data)
        if (data) {
          message.success('保存成功', 3)
          yield put({ type: 'hideModal' })
          yield put({ type: 'get' })
        }
      },
      * remove ({ payload }, { call, put }) {
        const data = yield call(service.delSvc, payload)
        if (data) {
          message.success('删除成功', 3)
          yield put({ type: 'get' })
        }
      },
    },
    reducers,
  })
}

function genCRUDComPageModel (namespace, service, state = {}, effects, setup, reducers) {
  return modelExtend(genComPageModel(namespace, service, state, effects, setup, reducers), {
    effects: {
      * create ({ payload }, { call, put }) {
        const data = yield call(service.addSvc, payload.data)
        if (data) {
          message.success('保存成功', 3)
          yield put({ type: 'hideModal' })
          yield put({ type: 'getPage' })
        }
      },
      * update ({ payload }, { call, put }) {
        const data = yield call(service.uptSvc, payload.data)
        if (data) {
          message.success('保存成功', 3)
          yield put({ type: 'hideModal' })
          yield put({ type: 'getPage' })
        }
      },
      * remove ({ payload }, { call, put }) {
        const data = yield call(service.delSvc, payload)
        if (data) {
          message.success('删除成功', 3)
          yield put({ type: 'getPage' })
        }
      },
    },
    reducers,
  })
}

export { model, comModel, comPageModel, genComModel, genComPageModel, genCRUDComModel, genCRUDComPageModel }
