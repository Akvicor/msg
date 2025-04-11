import React, {useEffect, useState} from 'react'
import {Button, Form, Input, Table, Popconfirm, Modal, message, Flex, Tag} from 'antd'
import {accessTokenFind, accessTokenCreate, accessTokenDelete, accessTokenUpdate} from "../../api/accessToken";
import './accessToken.css'
import {useSelector} from "react-redux";
import dayjs from "dayjs";

const AccessToken = () => {
  const mode = useSelector(state => state.mode)
  const [searchKeyword, setSearchKeyword] = useState({
    search: ''
  })
  const [tableData, setTableData] = useState([])
  const [inputAccessTokenDataAction, setInputAccessTokenDataAction] = useState('close');
  const [form] = Form.useForm()

  const updateTable = (search) => {
    accessTokenFind(search).then(({data}) => {
      if (data.code === 0) {
        setTableData(data.data)
      }
    })
  }

  const handleDeleteAccessToken = (id) => {
    accessTokenDelete({id: id}).then(({data}) => {
      if (data.code !== 0) {
        message.open({
          type: 'warning',
          content: '删除失败: ' + data.msg
        })
      } else {
        message.open({
          type: 'success',
          content: '删除成功'
        })
      }
      updateTable(searchKeyword)
    })
  }
  const handleSearchAccessToken = ({keyword}) => {
    setSearchKeyword({
      search: keyword
    })
  }
  const handleInputAccessTokenDataShow = (action, data) => {
    if (data) {
      const cloneData = JSON.parse(JSON.stringify(data))
      form.setFieldsValue(cloneData)
    }
    setInputAccessTokenDataAction(action)
  }

  const handleInputAccessTokenDataOk = () => {
    form.validateFields().then((input) => {
      if (inputAccessTokenDataAction === 'create') {
        accessTokenCreate(input).then(({data}) => {
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
          setInputAccessTokenDataAction('close')
          form.resetFields()
        })
      } else if (inputAccessTokenDataAction === 'update') {
        accessTokenUpdate(input).then(({data}) => {
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
          setInputAccessTokenDataAction('close')
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
      title: '名称',
      dataIndex: 'name',
    }, {
      title: 'Token',
      dataIndex: 'token'
    }, {
      title: '上次使用',
      dataIndex: 'last_used',
      render: (last_used) => {
        if (last_used === 0) {
          return '-'
        }
        return dayjs(last_used * 1000).format('YYYY-MM-DD HH:mm:ss')
      }
    }, {
      title: '操作',
      render: (rowData) => {
        return (
          <div>
            <Button type="primary" style={{marginRight: '5px'}}
                    onClick={() => handleInputAccessTokenDataShow('update', rowData)}>编辑</Button>
            <Popconfirm
              title={'删除Token'}
              description={"你确定删除" + rowData.name + "?"}
              onConfirm={() => handleDeleteAccessToken(rowData.id)}
              onCancel={() => {
              }}
              okText="确认"
              cancelText="取消"
            >
              <Button danger type="primary">删除</Button>
            </Popconfirm>
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
            <Button type="primary" onClick={() => handleInputAccessTokenDataShow('create')}>创建</Button>
            <Form
              layout="inline"
              onFinish={handleSearchAccessToken}
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
              <Button type="primary" onClick={() => handleInputAccessTokenDataShow('create')}>创建</Button>
            </div>
            <Form
              layout="inline"
              onFinish={handleSearchAccessToken}
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
        dataSource={tableData}
        pagination={{
          pageSizeOptions: [10, 15, 20, 50, 100],
          responsive: true,
          showQuickJumper: true,
          showSizeChanger: true
        }}
        scroll={{
          x: 'max-content',
        }}
        rowKey={'id'}
      />
      <Modal
        title={inputAccessTokenDataAction === 'create' ? '创建' : inputAccessTokenDataAction === 'update' ? '编辑' : 'Unknown'}
        open={inputAccessTokenDataAction !== 'close'}
        onOk={handleInputAccessTokenDataOk}
        onCancel={() => {
          setInputAccessTokenDataAction('close');
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
            inputAccessTokenDataAction === 'update' &&
            <Form.Item
              name="id"
              hidden
            >
              <Input/>
            </Form.Item>
          }
          <Form.Item
            label="名称"
            name="name"
            rules={[
              {
                required: true,
                message: '请输入名称'
              }
            ]}
          >
            <Input placeholder={'请输入名称'}/>
          </Form.Item>
        </Form>
      </Modal>
    </div>
  )
}

export default AccessToken
