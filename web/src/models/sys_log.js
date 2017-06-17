import * as service from '../services/sys_log'
import { checkApiRs } from '../utils'

export default {
  namespace: 'sysLog',
  state: {
    dataSource: [],
    loading: false,
    total: 0,
    pageIndex: 1,
    pageSize: 10,
  },

  subscriptions: {
    setup ({ dispatch, history }) {
        dispatch({
          type: 'query',
        })
    },
  },

  effects: {
    * query ({ payload }, { put, call }) {
      yield put({ type: 'showLoading' })
      const data = yield call(service.getPagingService, payload)
      yield put({ type: 'hideLoading' })
      if (data.success) {
        yield put({
          type: 'querySuccess',
          payload: { dataSource: data.r.data, total: data.r.total },
        })
      } else {
        checkApiRs(data)
      }
    },
  },
  reducers: {
    querySuccess (state, { payload }) {
      return {
        ...state,
        ...payload,
      }
    },
    showLoading (state) {
      return {
        ...state,
        loadingg: true,
      }
    },
    hideLoading (state) {
      return {
        ...state,
        loadingg: false,
      }
    },
  },
}
