import React, {useEffect, useState} from 'react'
import {Button, Form, Input, Table, Tag, Popconfirm, Modal, Select, message, Avatar, Flex} from 'antd'
import {userCreate, userDisable, userEnable, userFind, userUpdate} from "../../api/admin";
import './user.css'
import {ColorButtonProvider} from "../../theme/button";
import dayjs from "dayjs";
import {useSelector} from "react-redux";

const User = () => {
  const mode = useSelector(state => state.mode)
  const [searchKeyword, setSearchKeyword] = useState({
    search: ''
  })
  const [tableData, setTableData] = useState([])
  const [inputUserDataAction, setInputUserDataAction] = useState('close');
  const [form] = Form.useForm()

  const updateTable = (search) => {
    userFind(search).then(({data}) => {
      if (data.code === 0) {
        setTableData(data.data)
      }
    })
  }

  const handleDisableUser = (disable, id) => {
    if (disable) {
      userDisable({id: id}).then(({data}) => {
        if (data.code !== 0) {
          message.open({
            type: 'warning',
            content: '停用失败: ' + data.msg
          })
          return
        }
        message.open({
          type: 'success',
          content: '停用成功'
        })
        updateTable(searchKeyword)
      })
    } else {
      userEnable({id: id}).then(({data}) => {
        if (data.code !== 0) {
          message.open({
            type: 'warning',
            content: '启用失败' + data.msg
          })
          return
        }
        message.open({
          type: 'success',
          content: '启用成功'
        })
        updateTable(searchKeyword)
      })
    }
  }
  const handleSearchUser = ({keyword}) => {
    setSearchKeyword({
      search: keyword
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
      if (inputUserDataAction === 'create') {
        userCreate(input).then(({data}) => {
          if (data.code !== 0) {
            message.open({
              type: 'warning',
              content: '创建失败'
            })
            return
          }
          message.open({
            type: 'success',
            content: '创建成功'
          })
          updateTable(searchKeyword)
          setInputUserDataAction('close')
          form.resetFields()
        })
      } else if (inputUserDataAction === 'update') {
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
          updateTable(searchKeyword)
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
      title: '状态',
      dataIndex: 'disabled',
      render: (val) => {
        let color
        let status
        if (val === 0) {
          color = 'green';
          status = 'active'
        } else {
          color = 'red';
          status = 'disabled'
        }
        return (
          <>
            <Tag color={color} key={val}>
              {status}
            </Tag>
            {
              val !== 0 &&
              <Tag color={color} key={'time' + val}>
                {dayjs(val * 1000).format('YYYY-MM-DD HH:mm:ss')}
              </Tag>
            }
          </>
        );
      }
    }, {
      title: '操作',
      render: (rowData) => {
        let disableStatus = rowData.disabled === 0
        let disableMsg = disableStatus ? '停用' : '启用'
        return (
          <div>
            <Button type="primary" style={{marginRight: '5px'}}
                    onClick={() => handleInputUserDataShow('update', rowData)}>编辑</Button>
            <ColorButtonProvider danger={disableStatus} color="green">
              <Popconfirm
                title={disableMsg + '用户'}
                description={"你确定" + disableMsg + "用户?"}
                onConfirm={() => handleDisableUser(disableStatus, rowData.id)}
                onCancel={() => {
                }}
                okText="确认"
                cancelText="取消"
              >
                <Button danger={disableStatus} type="primary">{disableMsg}</Button>
              </Popconfirm>
            </ColorButtonProvider>
          </div>
        );
      }
    }
  ]

  useEffect(() => {
    updateTable(searchKeyword)
  }, [searchKeyword]);

  const customizeRequiredMark = (label, {required}) => (
    <>
      {required ? <Tag color="error">required</Tag> : <Tag color="warning">optional</Tag>}
      {label}
    </>
  );
  return (
    <div>
      {
        mode.isWide ? (
          <Flex style={{width: '100%', marginBottom: '15px'}} justify='space-between' align='center'>
            <Button type="primary" onClick={() => handleInputUserDataShow('create')}>创建</Button>
            <Form
              layout="inline"
              onFinish={handleSearchUser}
            >
              <Form.Item name="keyword">
                <Input placeholder='请输入关键词'/>
              </Form.Item>
              <Form.Item>
                <Button htmlType='submit' type='primary'>搜索</Button>
              </Form.Item>
            </Form>
          </Flex>
        ) : (
          <div style={{marginBottom: 15}}>
            <div style={{marginBottom: 6}}>
              <Button type="primary" onClick={() => handleInputUserDataShow('create')}>创建</Button>
            </div>
            <Form
              layout="inline"
              onFinish={handleSearchUser}
            >
              <Form.Item name="keyword">
                <Input placeholder='请输入关键词'/>
              </Form.Item>
              <Form.Item>
                <Button htmlType='submit' type='primary'>搜索</Button>
              </Form.Item>
            </Form>
          </div>
        )
      }
      <Table
          columns={columns}
          scroll={{
            x: 'max-content',
          }}
          dataSource={tableData}
          pagination={{
            pageSizeOptions: [10, 15, 20, 50, 100],
            responsive: true,
            showQuickJumper: true,
            showSizeChanger: true
          }}
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
          <Form.Item
            label="角色"
            name="role"
            rules={[
              {
                required: true,
                message: '请输入角色'
              }
            ]}
          >
            <Select
              placeholder="请选择角色"
              onChange={() => {
              }}
              allowClear
              options={[
                {value: 'user', label: 'User'},
                {value: 'viewer', label: 'Viewer'},
                {value: 'admin', label: 'Admin'}
              ]}
            />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  )
}

export default User
