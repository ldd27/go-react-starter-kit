import React from 'react'
import PropTypes from 'prop-types'
import { connect } from 'dva'
import { Input } from 'antd'
import Table from './table'

const Search = Input.Search
const SysLog = ({ dispatch, sysLog, loading }) => {
  const { data, pagination, search } = sysLog
  const { pageSize } = pagination

  const tableProps = {
    dataSource: data,
    pagination,
    loading: loading.effects['sysLog/getPage'],
    onChange (page) {
      dispatch({ type: 'sysLog/getPage', payload: { current: page.current, pageSize: page.pageSize, search } })
    },
  }

  const onSearch = (title) => {
    dispatch({ type: 'sysLog/getPage', payload: { current: 1, pageSize, search: { title } } })
  }

  return (
    <div className='content-inner'>
      <Search placeholder='标题' style={{ width: 200, marginBottom: 16 }} onSearch={onSearch} />
      <Table {...tableProps} />
    </div>
  )
}

SysLog.propTypes = {
  sysLog: PropTypes.shape,
  dispatch: PropTypes.func,
  loading: PropTypes.shape,
}

export default connect(({ sysLog, loading }) => ({ sysLog, loading }))(SysLog)
