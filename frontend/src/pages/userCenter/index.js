import React, {useEffect, useState} from 'react'
import {Button, Form, Input, Table, Tag, Modal, message, Avatar} from 'antd'
import {userInfo, userUpdate} from "../../api/user";
import './userCenter.css'

const UserCenter = () => {
  const [tableData, setTableData] = useState([])
  const [inputUserDataAction, setInputUserDataAction] = useState('close');
  const [form] = Form.useForm()

  const updateTable = () => {
    userInfo().then(({data}) => {
      if (data.code === 0) {
        setTableData([data.data])
      }
    })
  }

  const handleInputUserDataShow = (action, data) => {
    if (data) {
      const cloneData = JSON.parse(JSON.stringify(data))
      form.setFieldsValue(cloneData)
    }
    setInputUserDataAction(action)
  }

  const handleInputUserDataOk = () => {
    form.validateFields().then((input) => {
      if (inputUserDataAction === 'update') {
        userUpdate(input).then(({data}) => {
          if (data.code !== 0) {
            message.open({
              type: 'warning',
              content: '更新失败'
            })
            return
          }
          message.open({
            type: 'success',
            content: '更新成功'
          })
          updateTable()
          setInputUserDataAction('close')
          form.resetFields()
        })
      }
    }).catch(() => {
      message.open({
        type: 'warning',
        content: '请检查输入'
      })
    })
  }

  const columns = [
    {
      title: '头像',
      dataIndex: 'avatar',
      render: (avatar) => {
        return <Avatar size="large" src={avatar ? avatar : ''}/>
      }
    }, {
      title: '用户名',
      dataIndex: 'username',
    }, {
      title: '昵称',
      dataIndex: 'nickname'
    }, {
      title: '邮箱',
      dataIndex: 'mail'
    }, {
      title: '手机',
      dataIndex: 'phone'
    }, {
      title: '角色',
      dataIndex: 'role',
      render: (role) => {
        let color = 'yellow';
        if (role === 'admin') {
          color = 'geekblue';
        } else if (role === 'user') {
          color = 'green'
        }
        return (
          <Tag color={color} key={role}>
            {role}
          </Tag>
        );
      }
    }, {
      title: '操作',
      render: (rowData) => {
        return (
          <div>
            <Button type="primary" style={{marginRight: '5px'}}
                    onClick={() => handleInputUserDataShow('update', rowData)}>编辑</Button>
          </div>
        );
      }
    }
  ]

  useEffect(() => {
    updateTable()
  }, []);

  const customizeRequiredMark = (label, {required}) => (
    <>
      {required ? <Tag color="error">required</Tag> : <Tag color="warning">optional</Tag>}
      {label}
    </>
  );
  return (
    <div>
      <Table
          columns={columns}
          scroll={{
            x: 'max-content',
          }}
          dataSource={tableData}
          rowKey={'id'}
      />
      <Modal
        title={inputUserDataAction === 'create' ? '创建' : inputUserDataAction === 'update' ? '编辑' : 'Unknown'}
        open={inputUserDataAction !== 'close'}
        onOk={handleInputUserDataOk}
        onCancel={() => {
          setInputUserDataAction('close');
          form.resetFields()
        }}
        okText="确定"
        cancelText="取消"
      >
        <Form
          form={form}
          labelCol={{
            span: 6
          }}
          wrapperCol={{
            span: 18
          }}
          requiredMark={customizeRequiredMark}
        >
          {
            inputUserDataAction === 'update' &&
            <Form.Item
              name="id"
              hidden
            >
              <Input/>
            </Form.Item>
          }
          <Form.Item
            label="用户名"
            name="username"
            rules={[
              {
                required: true,
                message: '请输入用户名'
              }
            ]}
          >
            <Input placeholder={'请输入用户名'}/>
          </Form.Item>
          <Form.Item
            label="密码"
            name="password"
          >
            <Input.Password placeholder={'请输入密码'}/>
          </Form.Item>
          <Form.Item
            label="昵称"
            name="nickname"
          >
            <Input placeholder={'请输入昵称'}/>
          </Form.Item>
          <Form.Item
            label="头像"
            name="avatar"
          >
            <Input placeholder={'请输入头像URL'}/>
          </Form.Item>
          <Form.Item
            label="邮箱"
            name="mail"
            rules={[
              {
                type: 'email',
                message: '请输入正确的邮箱地址'
              }
            ]}
          >
            <Input placeholder={'请输入邮箱'}/>
          </Form.Item>
          <Form.Item
            label="手机号"
            name="phone"
          >
            <Input placeholder={'请输入手机号'}/>
          </Form.Item>
        </Form>
      </Modal>
    </div>
  )
}

export default UserCenter
