import React from 'react'
import PropTypes from 'prop-types'
import { connect } from 'dva'
import { Input } from 'antd'
import Table from './table'

const Search = Input.Search
const SysLog = ({ dispatch, sysLog }) => {
  const { total, pageIndex, pageSize, dataSource, loading, title } = sysLog

  const tableProps = {
    dataSource,
    loading,
    total,
    pageIndex,
    pageSize,
    onPageChange (page) {
      dispatch({ type: 'sysLog/query', payload: { pageIndex: page, pageSize, title } })
    },
  }

  const onSearch = (value) => {
    dispatch({ type: 'sysLog/query', payload: { pageIndex, pageSize, title: value } })
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
}

export default connect(({ sysLog }) => ({ sysLog }))(SysLog)
