import React from 'react'
import { formatDate } from 'utils'
import { DDTable } from 'components'

const Table = ({ ...tableProps }) => {
  const columns = [
    {
      title: '字典编码',
      dataIndex: 'ItemCode',
      key: 'ItemCode',
    },
    {
      title: '字典名称',
      dataIndex: 'ItemName',
      key: 'ItemName',
    },
    {
      title: '系统字典',
      dataIndex: 'IsSys',
      key: 'IsSys',
    },
    {
      title: '状态',
      dataIndex: 'Status',
      key: 'Status',
      render: text => formatDate(text),
    },
  ]

  tableProps = {
    ...tableProps,
    isMotion: true,
    columns,
    rowKey: record => record.Id,
  }

  return (
    <DDTable
      {...tableProps}
    />
  )
}

export default Table
