// 公有state
const comState = {
  table: {
    loading: false,
    pagination: {
      showSizeChanger: true,
      showQuickJumper: true,
      showTotal: total => `总共 ${total} 条数据`,
      current: 1,
      pageSize: 10,
      total: 0,
    },
    datas: [],
  },
  search: {},
  currentItem: {},
  modal: {
    visible: false,
    type: 'create',
    loading: false,
  },
}
// 公用reducer
const comReducer = {
  showLoading (state, action) {
    return { ...state, ...action.payload, 'table.loading': true }
  },
  success (state, action) {
    return { ...state, ...action.payload, 'table.loading': false }
  },
  fail (state, action) {
    return { ...state, ...action.payload, 'table.loading': false }
  },
  showModal (state, action) {
    return { ...state, ...action.payload, 'modal.visible': true }
  },
  hideModal (state, action) {
    return { ...state, ...action.payload, 'modal.visible': false }
  },
  update (state, action) {
    return { ...state, ...action.payload }
  },
}


/**
 * base model
 * @param {any} namespace
 * @param {any} state
 * @param {any} effects
 * @param {any} setup
 * @param {any} reducers
 * @returns
 */
function comModel (namespace, state, effects, setup, reducers) {
  return {
    namespace,
    state: {
      ...comState,
      ...state,
    },
    subscriptions: {
      setup ({ dispatch, history }) {
        if (!setup) {
          history.listen((location) => {
            if (location.pathname === `/${namespace}`) {
              dispatch({ type: 'getPaging', payload: { pageIndex: 1, pageSize: 10 } })
            }
          })
        } else {
          setup({ dispatch, history })
        }
      },
    },
    effects: {
      ...effects,
    },
    reducers: {
      ...comReducer,
      ...reducers,
    },
  }
}

function comCRUDModel (namespace, state, service, effects, setup, reducers) {

}