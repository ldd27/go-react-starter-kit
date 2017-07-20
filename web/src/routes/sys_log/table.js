import React from 'react'
import DDTable from '../../components/table'
import { formatDate } from '../../utils'

const Table = ({ ...tableProps }) => {
  const columns = [
    {
      title: 'ID',
      dataIndex: 'Id',
      key: 'Id',
    },
    {
      title: '标题',
      dataIndex: 'Title',
      key: 'Title',
    },
    {
      title: '操作人',
      dataIndex: 'OpUserName',
      key: 'OpUserName',
    },
    {
      title: '操作时间',
      dataIndex: 'OpTime',
      key: 'OpTime',
      render: text => formatDate(text),
    },
  ]

  tableProps = {
    ...tableProps,
    isMotion: true,
    columns,
    rowKey: record => record.Id,
    expandedRowRender: record => record.Info,
    pagination: {
      total: tableProps.total,
      pageSize: tableProps.pageSize,
      defaultCurrent: 1,
      current: tableProps.pageIndex,
    },
  }

  return (
    <DDTable
      {...tableProps}
    />
  )
}

export default Table
