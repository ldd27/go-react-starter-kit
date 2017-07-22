import React from 'react'
import { connect } from 'dva'
import { Row, Col } from 'antd'
import PropTypes from 'prop-types'
import Tree from './tree'
import Table from './table'

const Index = ({ dispatch, dict, loading }) => {
  const { tree, data, selectTree } = dict
  const treeProps = {
    selected: selectTree,
    treeData: tree,
    onSelect (key, DictCode, IsSys) {
      dispatch({ type: 'dict/updateState', payload: { selectTree: key, search: { DictCode, IsSys } } })
      dispatch({ type: 'dict/get', payload: { search: { DictCode } } })
    },
  }

  const tableProps = {
    dataSource: data,
    loading: loading.effects['dict/get'],
  }
  return (
    <div className='content-inner'>
      <Row>
        <Col span={4}><Tree {...treeProps} /></Col>
        <Col span={18}>
          <Table {...tableProps} />
        </Col>
      </Row>
    </div>
  )
}

Index.propTypes = {
  dict: PropTypes.shape,
  dispatch: PropTypes.func,
  loading: PropTypes.shape,
}

export default connect(({ dict, loading }) => ({ dict, loading }))(Index)
