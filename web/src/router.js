import React from 'react'
import PropTypes from 'prop-types'
import { Router } from 'dva/router'
import App from './routes/app'

const registerModel = (app, model) => {
  if (!(app._models.filter(m => m.namespace === model.namespace).length === 1)) {
    app.model(model)
  }
}

const Routers = function ({ history, app }) {
  const routes = [
    {
      path: '/',
      component: App,
      getIndexRoute (nextState, cb) {
        require.ensure([], (require) => {
          registerModel(app, require('./models/sys_log'))
          cb(null, { component: require('./routes/home') })
        }, 'home')
      },
      childRoutes: [
        {
          path: 'home',
          getComponent (nextState, cb) {
            require.ensure([], (require) => {
              registerModel(app, require('./models/sys_log'))
              cb(null, require('./routes/home'))
            }, 'home')
          },
        },
        {
          path: 'login',
          getComponent (nextState, cb) {
            require.ensure([], (require) => {
              registerModel(app, require('./models/login'))
              cb(null, require('./routes/login'))
            }, 'login')
          },
        },
        {
          path: 'sysLog',
          getComponent (nextState, cb) {
            require.ensure([], (require) => {
              registerModel(app, require('./models/sys_log'))
              cb(null, require('./routes/sys_log'))
            }, 'sysLog')
          },
        },
        {
          path: 'dict',
          getComponent (nextState, cb) {
            require.ensure([], (require) => {
              registerModel(app, require('./models/dict'))
              cb(null, require('./routes/dict'))
            }, 'dict')
          },
        },
        {
          path: '*',
          getComponent (nextState, cb) {
            require.ensure([], (require) => {
              cb(null, require('./routes/error/'))
            }, 'error')
          },
        },
      ],
    },
  ]

  return <Router history={history} routes={routes} />
}

Routers.propTypes = {
  history: PropTypes.shape,
  app: PropTypes.shape,
}

export default Routers
