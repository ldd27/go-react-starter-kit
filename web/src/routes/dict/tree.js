import React from 'react'
import { Tree } from 'antd'
import PropTypes from 'prop-types'

const TreeNode = Tree.TreeNode
const CusTree = ({ selected, treeData, onSelect }) => {
  const genTreeNodes = data => data.map((item) => {
    let temp = item.key.split('-')
    let isSys = true
    if (temp.length === 2) {
      isSys = temp[1] === 'y'
    }
    if (item.children) {
      return (
        <TreeNode title={`${item.title}${isSys ? '' : '(可添加)'}`} disabled={item.type === 'type'} key={`${item.key}`}>
          {genTreeNodes(item.children)}
        </TreeNode>
      )
    }
    return (<TreeNode title={`${item.title}${isSys ? '' : '(可添加)'}`} disabled={item.type === 'type'} key={`${item.key}`} />)
  })

  const treeProps = {
    defaultExpandAll: true,
    selectedKeys: [selected],
    showLine: true,
    onSelect (keys) {
      if (!keys || keys.length === 0) {
      } else {
        let temp = keys[0].split('-')
        onSelect(keys[0], temp[0], temp[1])
      }
    }
  }

  return (
    <div>
      { treeData.length === 0 ? <span>暂无数据</span> : <Tree {...treeProps}>{genTreeNodes(treeData)}</Tree> }
    </div>
  )
}

CusTree.propTypes = {
  selected: PropTypes.string,
  treeData: PropTypes.shape,
  onSelect: PropTypes.func,
}

export default CusTree
