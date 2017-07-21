import React from 'react'
import PropTypes from 'prop-types'
import { Table } from 'antd'
import classnames from 'classnames'
import AnimTableBody from './AnimTableBody'
import styles from './index.less'

const DDTable = ({ page, isMotion, ...tableProps }) => {

  const getBodyWrapperProps = {
    page,
    current: tableProps.pagination.current,
  }

  const getBodyWrapper = (body) => { return isMotion ? <AnimTableBody {...getBodyWrapperProps} body={body} /> : body }

  return (
    <Table
      {...tableProps}
      className={classnames({ [styles.table]: true, [styles.motion]: isMotion })}
      bordered
      scroll={{ x: 1024 }}
      simple
      size={'middle'}
      getBodyWrapper={getBodyWrapper}
    />
  )
}

DDTable.propTypes = {
  isMotion: PropTypes.bool,
  page: PropTypes.number,
  tableProps: PropTypes.shape,
}

export default DDTable
