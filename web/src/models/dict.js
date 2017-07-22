import { genCRUDComModel, checkApiRs } from 'utils'
import * as svc from '../services/dict'

const state = {
  tree: [],
  selectTree: null,
}
const effects = {
  * getTree ({ payload }, { call, put }) {
    const data = yield call(svc.getIndexTreeSvc)
    if (data.success) {
      yield put({ type: 'success',
        payload: { tree: data.r },
      })
    } else {
      checkApiRs(data)
    }
  },
}
const setup = ({ dispatch, history }) => {
  history.listen((location) => {
    if (location.pathname === '/dict') {
      dispatch({ type: 'getTree' })
    }
  })
}
export default genCRUDComModel('dict', svc, state, effects, setup)
