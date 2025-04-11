import React, {useEffect, useState} from 'react'
import {Button, Form, Input, Table, Popconfirm, Modal, message, Flex, Select, DatePicker} from 'antd'
import {channelCreate, channelFind, channelUpdate, channelDelete, channelTest, channelSend} from "../../api/channel";
import './channel.css'
import {useSelector} from "react-redux";
import {typeBotSendersFind, typeChannelTypeFind, typeSendTypeFind} from "../../api/type";
import {SenderStatusNotSupported} from "./senderStatus";
import {ColorButtonProvider} from "../../theme/button";
import dayjs from "dayjs";

const Channel = () => {
  const mode = useSelector(state => state.mode)
  const [searchKeyword, setSearchKeyword] = useState({
    search: ''
  })
  const [sendTypeData, setSendTypeData] = useState([])
  const [channelTypeData, setChannelTypeData] = useState([])
  const [botSendersData, setBotSendersData] = useState([])
  const [tableData, setTableData] = useState([])
  // Create
  const [inputChannelDataAction, setInputChannelDataAction] = useState('close');
  const [form] = Form.useForm()
  // Send
  const [inputSendDataAction, setInputSendDataAction] = useState('close');
  const [formSend] = Form.useForm()

  useEffect(() => {
    typeChannelTypeFind().then(({data}) => {
      if (data.code === 0) {
        let options = []
        data.data.forEach((item) => {
          options.push({
            value: item.type,
            label: item.name
          })
        })
        setChannelTypeData(options)
      }
    })
  }, [])

  useEffect(() => {
    typeSendTypeFind().then(({data}) => {
      if (data.code === 0) {
        let options = []
        data.data.forEach((item) => {
          options.push({
            value: item.type,
            label: item.name
          })
        })
        setSendTypeData(options)
      }
    })
  }, [])

  useEffect(() => {
    typeBotSendersFind().then(({data}) => {
      if (data.code === 0) {
        let options = []
        data.data.forEach((item) => {
          let opt = {
            value: item.key,
            label: item.name
          }
          let supportedSender = ""
          let supportedReceiver = ""
          item.status.forEach((stat) => {
            if (stat.sender_status !== SenderStatusNotSupported) {
              supportedSender += stat.channel_name + "/"
            }
            if (stat.receiver_status !== SenderStatusNotSupported) {
              supportedReceiver += stat.channel_name + "/"
            }
          })
          if (supportedSender.length > 0) {
            supportedSender = supportedSender.slice(0, -1)
            opt.label += " S[" + supportedSender + "]"
          }
          if (supportedReceiver.length > 0) {
            supportedReceiver = supportedReceiver.slice(0, -1)
            opt.label += " R[" + supportedReceiver + "]"
          }
          options.push(opt)
        })
        setBotSendersData(options)
      }
    })
  }, [])

  const updateTable = (search) => {
    channelFind(search).then(({data}) => {
      if (data.code === 0) {
        setTableData(data.data)
      }
    })
  }

  const handleDeleteChannel = (id) => {
    channelDelete({id: id}).then(({data}) => {
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
  const handleTestChannel = (id) => {
    channelTest({id: id}).then(({data}) => {
      if (data.code !== 0) {
        message.open({
          type: 'warning',
          content: '测试失败: ' + data.msg
        })
      } else {
        message.open({
          type: 'success',
          content: '测试成功'
        })
      }
    })
  }
  const handleSearchChannel = ({keyword}) => {
    setSearchKeyword({
      search: keyword
    })
  }
  const handleInputChannelDataShow = (action, data) => {
    if (data) {
      const cloneData = JSON.parse(JSON.stringify(data))
      form.setFieldsValue(cloneData)
    }
    setInputChannelDataAction(action)
  }

  const handleInputChannelDataOk = () => {
    form.validateFields().then((input) => {
      if (inputChannelDataAction === 'create') {
        channelCreate(input).then(({data}) => {
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
          setInputChannelDataAction('close')
          form.resetFields()
        })
      } else if (inputChannelDataAction === 'update') {
        channelUpdate(input).then(({data}) => {
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
          setInputChannelDataAction('close')
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

  const handleInputSendDataShow = (action, data) => {
    if (data) {
      const cloneData = JSON.parse(JSON.stringify(data))
      cloneData.at_date = dayjs()
      cloneData.at_time = dayjs()
      formSend.setFieldsValue(cloneData)
    }
    setInputSendDataAction(action)
  }

  const handleInputSendDataOk = () => {
    formSend.validateFields().then((input) => {
      input.at = dayjs(input.at_date.format('YYYY-MM-DD') + ' ' + input.at_time.format('HH:mm:ss'))
      if (inputSendDataAction === 'create') {
        input.at = dayjs(input.at).unix()
        delete input.at_date
        delete input.at_time
        channelSend(input).then(({data}) => {
          if (data.code !== 0) {
            message.open({
              type: 'warning',
              content: '发送失败'
            })
            return
          }
          message.open({
            type: 'success',
            content: '发送成功'
          })
          setInputSendDataAction('close')
          // formSend.resetFields()
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
      title: 'ID',
      dataIndex: 'id',
    }, {
      title: '名称',
      dataIndex: 'name',
    }, {
      title: '标记',
      dataIndex: 'sign'
    }, {
      title: '类型',
      dataIndex: 'type',
      render: (type) => {
        let cType = channelTypeData.find(item => item.value === type)
        if (cType) {
          return cType.label
        }
        return type
      }
    }, {
      title: 'Bot',
      dataIndex: 'bot',
      render: (type) => {
        let cType = botSendersData.find(item => item.value === type)
        if (cType) {
          return cType.label
        }
        return type
      }
    }, {
      title: '目标',
      dataIndex: 'target'
    }, {
      title: '操作',
      render: (rowData) => {
        return (
          <div>
            <ColorButtonProvider danger={false} color="green">
              <Button type="primary" style={{marginRight: '5px'}}
                      onClick={() => handleInputSendDataShow('create', {id: rowData.id})}>发送</Button>
            </ColorButtonProvider>
            <Button type="primary" style={{marginRight: '5px'}}
                    onClick={() => handleTestChannel(rowData.id)}>测试</Button>
            <Button type="primary" style={{marginRight: '5px'}}
                    onClick={() => handleInputChannelDataShow('update', rowData)}>编辑</Button>
            <Popconfirm
              title={'删除通知渠道'}
              description={"你确定删除" + rowData.name + "?"}
              onConfirm={() => handleDeleteChannel(rowData.id)}
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

  return (
    <div>
      {
        mode.isWide ? (
          <Flex style={{width: '100%', marginBottom: '15px'}} justify='space-between' align='center'>
            <Button type="primary" onClick={() => handleInputChannelDataShow('create')}>创建</Button>
            <Form
              layout="inline"
              onFinish={handleSearchChannel}
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
              <Button type="primary" onClick={() => handleInputChannelDataShow('create')}>创建</Button>
            </div>
            <Form
              layout="inline"
              onFinish={handleSearchChannel}
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
        title={inputChannelDataAction === 'create' ? '创建' : inputChannelDataAction === 'update' ? '编辑' : 'Unknown'}
        open={inputChannelDataAction !== 'close'}
        onOk={handleInputChannelDataOk}
        onCancel={() => {
          setInputChannelDataAction('close');
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
        >
          {
            inputChannelDataAction === 'update' &&
            <Form.Item
              name="id"
              hidden
            >
              <Input/>
            </Form.Item>
          }
          <Form.Item
            label="标记"
            name="sign"
            rules={[
              {
                required: true,
                message: '请输入标记'
              }
            ]}
          >
            <Input placeholder={'请输入标记'}/>
          </Form.Item>
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
          <Form.Item
            label="Bot"
            name="bot"
            rules={[
              {
                required: true,
                message: '请输入Bot'
              }
            ]}
          >
            <Select
              placeholder="请输入Bot"
              onChange={() => {
              }}
              allowClear
              options={botSendersData}
            />
          </Form.Item>
          <Form.Item
            label="类型"
            name="type"
            rules={[
              {
                required: true,
                message: '请输入类型'
              }
            ]}
          >
            <Select
              placeholder="类型"
              onChange={() => {
              }}
              allowClear
              options={channelTypeData}
            />
          </Form.Item>
          <Form.Item
            label="目标"
            name="target"
            rules={[
              {
                required: true,
                message: '请输入目标'
              }
            ]}
          >
            <Input placeholder={'请输入目标'}/>
          </Form.Item>
        </Form>
      </Modal>
      <Modal
        title={inputSendDataAction === 'create' ? '发送' : 'Unknown'}
        open={inputSendDataAction !== 'close'}
        onOk={handleInputSendDataOk}
        onCancel={() => {
          setInputSendDataAction('close');
          // formSend.resetFields()
        }}
        okText="确定"
        cancelText="取消"
      >
        <Form
          form={formSend}
          labelCol={{
            span: 6
          }}
          wrapperCol={{
            span: 18
          }}
        >
          <Form.Item
            name="id"
            hidden
          >
            <Input/>
          </Form.Item>
          <Form.Item
            label="类型"
            name="type"
            rules={[
              {
                required: true,
                message: '请输入消息类型'
              }
            ]}
          >
            <Select
              placeholder="请输入消息类型"
              onChange={() => {
              }}
              allowClear
              options={sendTypeData}
            />
          </Form.Item>
          <Form.Item
            label="标题"
            name="title"
            rules={[
              {
                required: false,
                message: '请输入标题'
              }
            ]}
          >
            <Input placeholder={'请输入标题'}/>
          </Form.Item>
          <Form.Item
            label="内容"
            name="msg"
            rules={[
              {
                required: false,
                message: '请输入内容'
              }
            ]}
          >
            <Input placeholder={'请输入内容'}/>
          </Form.Item>
          <Form.Item
            label="发送日期"
            name="at_date"
            rules={[
              {
                required: false,
                message: '请输入发送日期'
              }
            ]}
          >
            <DatePicker picker='date' inputReadOnly style={{width: '100%'}} placeholder={'请输入发送日期'}/>
          </Form.Item>
          <Form.Item
            label="发送时间"
            name="at_time"
            rules={[
              {
                required: false,
                message: '请输入发送时间'
              }
            ]}
          >
            <DatePicker picker='time' inputReadOnly style={{width: '100%'}} placeholder={'请输入发送时间'}/>
          </Form.Item>
        </Form>
      </Modal>
    </div>
  )
}

export default Channel
