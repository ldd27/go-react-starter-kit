import React from 'react'
import DDTable from '../../components/table'

const Table = ({ ...tableProps }) => {
  const columns = [

  ]

  return (
    <DDTable
      {...tableProps}
      columns
      rowKey={record => record.Id}
    />
  )
}

export default Table
