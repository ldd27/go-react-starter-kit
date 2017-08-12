import React from 'react'
import PropTypes from 'prop-types'
import { connect } from 'dva'
import { Button, Form, Input } from 'antd'
import { logo, name } from 'config'
import styles from './index.less'

const FormItem = Form.Item

const Login = ({
  login,
  dispatch,
  form: {
    getFieldDecorator,
    validateFieldsAndScroll,
  },
}) => {
  const { loginLoading } = login

  function handleOk () {
    validateFieldsAndScroll((errors, values) => {
      if (errors) {
        return
      }
      dispatch({ type: 'login/login', payload: values })
    })
  }

  return (
    <div className={styles.bg}>
      <div className={styles.header}>
        <img alt={'logo'} src={logo} />
        <br />
        <br />
        <h1 className={styles.title}>{name}</h1>
        <span className={styles.desc}>{name}</span>
      </div>
      <div className={styles.form}>
        <Form>
          <FormItem hasFeedback>
            {getFieldDecorator('LoginKey', {
              rules: [
                {
                  required: true,
                },
              ],
            })(<Input size='large' onPressEnter={handleOk} placeholder='用户名' prefix={<Icon type='user' />} />)}
          </FormItem>
          <FormItem hasFeedback>
            {getFieldDecorator('Password', {
              rules: [
                {
                  required: true,
                },
              ],
            })(<Input size='large' type='password' onPressEnter={handleOk} placeholder='密码' prefix={<Icon type='lock' />} />)}
          </FormItem>
          <FormItem>
            <Button type='primary' size='large' style={{ width: '100%' }} onClick={handleOk} loading={loginLoading}>
              登录
            </Button>
          </FormItem>
        </Form>
      </div>
    </div>
  )
}

Login.propTypes = {
  form: PropTypes.shape,
  login: PropTypes.shape,
  dispatch: PropTypes.func,
}

export default connect(({ login }) => ({ login }))(Form.create()(Login))
