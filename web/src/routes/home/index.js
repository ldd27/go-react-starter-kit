import React from 'react'
import PropTypes from 'prop-types'
import { connect } from 'dva'
import Table from '../../components/table'

const Home = ({ sysLog }) => {
  const columns = [
    {
      title: 'Id',
      dataIndex: 'Id',
      key: 'Id',
    },
    {
      title: 'Type',
      dataIndex: 'Type',
      key: 'Type',
    },
    {
      title: 'Title',
      dataIndex: 'Title',
      key: 'Title',
    },
  ]

  const tableProps = {
    isMotion: true,
    columns,
    rowKey: record => record.Id,
    loading: false,
    pagination: {
      total: sysLog.total,
      pageSize: 10,
      defaultCurrent: 1,
      current: 1,
    },
    dataSource: sysLog.dataSource,
  }
  return (
    <div className='content-inner'>
      <Table {...tableProps} />
    </div>
  )
}

Home.propTypes = {
  sysLog: PropTypes.shape,
}

export default connect(({ sysLog }) => ({ sysLog }))(Home)
