import React from "react"
import {Form, Input, Checkbox, Button, message} from 'antd'
import {useNavigate} from "react-router-dom";
import {getVersion, userLogin} from '../../api/public'
import {userInfo} from "../../api/user";

import './login.css'

const Login = () => {
  const navigate = useNavigate()
  if (localStorage.getItem('login-token')) {
    userInfo().then(({data}) => {
      if (data.code === 0) {
        localStorage.setItem('role', data.data.user.role)
        localStorage.setItem('user', JSON.stringify(data.data.user))
        navigate('/home')
      }
    })
  }
  getVersion().then(({data}) => {
    if (data.code !== 0) {
      localStorage.setItem('api-version', 'unknown')
    } else {
      localStorage.setItem('api-version', data.data)
    }
  })
  const handleSubmit = (input) => {
    if (!input.username || !input.password) {
      return message.open({
        type: 'warning',
        content: '请输入用户名和密码'
      })
    }
    userLogin(input).then(({data}) => {
      if (data.code !== 0) {
        message.open({
          type: 'warning',
          content: '登录失败'
        })
        return
      }
      message.open({
        type: 'success',
        content: '登录成功'
      })
      localStorage.setItem('login-token', data.data.token)
      localStorage.setItem('role', data.data.user.role)
      localStorage.setItem('user', JSON.stringify(data.data.user))
      navigate('/home')
    })
  }
  return (
    <Form className="login-container" onFinish={handleSubmit}>
      <div className="login-title">
        登录
      </div>
      <Form.Item
        label="账号"
        name="username"
      >
        <Input placeholder="请输入账号"/>
      </Form.Item>
      <Form.Item
        label="密码"
        name="password"
      >
        <Input.Password placeholder="请输入密码"/>
      </Form.Item>
      <Form.Item
        name="remember"
        valuePropName="checked"
      >
        <Checkbox>Remember me</Checkbox>
      </Form.Item>
      <Form.Item className="login-button">
        <Button type="primary" htmlType="submit">Sign in</Button>
      </Form.Item>
    </Form>
  )
}

export default Login
