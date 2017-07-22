import { routerRedux } from 'dva/router'
import { parse } from 'qs'
import { checkApiRs, cookie } from 'utils'
import { prefix } from 'config'
import * as service from '../services/app'

export default {
  namespace: 'app',
  state: {
    user: {},
    // menu: cookie.getCookie(`${prefix}menu`) ? JSON.parse(cookie.getCookie(`${prefix}menu`)) : [],
    menu: [],
    menuPopoverVisible: false,
    siderFold: localStorage.getItem(`${prefix}siderFold`) === 'true',
    darkTheme: localStorage.getItem(`${prefix}darkTheme`) === 'true',
    isNavbar: document.body.clientWidth < 769,
    navOpenKeys: JSON.parse(localStorage.getItem(`${prefix}navOpenKeys`)) || [],
  },
  subscriptions: {
    setup ({ dispatch }) {
      dispatch({ type: 'checkIsLogin' })
      let tid
      window.onresize = () => {
        clearTimeout(tid)
        tid = setTimeout(() => {
          dispatch({ type: 'changeNavbar' })
        }, 300)
      }
    },
  },
  effects: {
    * checkIsLogin ({ payload }, { call, put }) {
      const data = yield call(service.checkIsLoginService, parse(payload))
      if (data.success) {
        // cookie.setCookie(`${prefix}username`, data.r.UserName)
        cookie.setCookie(`${prefix}menu`, JSON.stringify(data.r.Menus))
        yield put({
          type: 'querySuccess',
          payload: { username: data.r.UserName },
        })
        yield put({
          type: 'common',
          payload: { menu: data.r.Menus },
        })
        if (location.pathname === '/login') {
          yield put(routerRedux.push('/home'))
        }
      } else if (location.pathname !== '/login') {
        let from = location.pathname
        if (location.pathname === '/home') {
          from = '/home'
        }
        window.location = `${location.origin}/login?from=${from}`
      }
    },
    * logout ({ payload }, { call }) {
      const data = yield call(service.logoutService, parse(payload))
      if (data.success) {
        // cookie.delCookie(`${prefix}username`, data.r.UserName)
        cookie.delCookie(`${prefix}token`, data.r.Token)
        cookie.delCookie(`${prefix}menu`, data.r.Menus)
        if (location.pathname !== '/login') {
          let from = location.pathname
          if (location.pathname === '/home') {
            from = '/home'
          }
          window.location = `${location.origin}/login?from=${from}`
        }
      } else {
        checkApiRs(data)
      }
    },
    * changeNavbar ({ payload }, { put, select }) {
      const { app } = yield (select(_ => _))
      const isNavbar = document.body.clientWidth < 769
      if (isNavbar !== app.isNavbar) {
        yield put({ type: 'handleNavbar', payload: isNavbar })
      }
    },
  },
  reducers: {
    querySuccess (state, { payload: user }) {
      return {
        ...state,
        user,
      }
    },

    switchSider (state) {
      localStorage.setItem(`${prefix}siderFold`, !state.siderFold)
      return {
        ...state,
        siderFold: !state.siderFold,
      }
    },

    switchTheme (state) {
      localStorage.setItem(`${prefix}darkTheme`, !state.darkTheme)
      return {
        ...state,
        darkTheme: !state.darkTheme,
      }
    },

    switchMenuPopver (state) {
      return {
        ...state,
        menuPopoverVisible: !state.menuPopoverVisible,
      }
    },

    handleNavbar (state, { payload }) {
      return {
        ...state,
        isNavbar: payload,
      }
    },

    handleNavOpenKeys (state, { payload: navOpenKeys }) {
      return {
        ...state,
        ...navOpenKeys,
      }
    },

    common (state, { payload }) {
      return {
        ...state,
        ...payload,
      }
    },
  },
}
