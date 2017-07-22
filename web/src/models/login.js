import { routerRedux } from 'dva/router'
import { queryURL, cookie, config, checkApiRs } from 'utils'
import * as service from '../services/login'

const { prefix } = config

export default {
  namespace: 'login',
  state: {
    loginLoading: false,
  },

  effects: {
    * login ({
      payload,
    }, { put, call }) {
      yield put({ type: 'showLoginLoading' })
      const data = yield call(service.loginService, payload)
      yield put({ type: 'hideLoginLoading' })
      if (data.success) {
        // cookie.setCookie(`${prefix}username`, data.r.UserName)
        cookie.setCookie(`${prefix}token`, data.r.Token)
        cookie.setCookie(`${prefix}menu`, JSON.stringify(data.r.Menus))
        const from = queryURL('from')
        yield put({
          type: 'app/common',
          payload: { menu: data.r.Menus },
        })
        yield put({
          type: 'app/querySuccess',
          payload: { username: data.r.UserName },
        })

        if (from) {
          yield put(routerRedux.push(from))
        } else {
          yield put(routerRedux.push('/home'))
        }
      } else {
        checkApiRs(data)
      }
    },
  },
  reducers: {
    showLoginLoading (state) {
      return {
        ...state,
        loginLoading: true,
      }
    },
    hideLoginLoading (state) {
      return {
        ...state,
        loginLoading: false,
      }
    },
  },
}
