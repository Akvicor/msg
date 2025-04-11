import React, {useEffect, useState} from 'react'
import {Button, Form, Input, Table, Popconfirm, Modal, message, Flex, Select, DatePicker, InputNumber} from 'antd'
import {
  scheduleCreate,
  scheduleFind,
  scheduleUpdate,
  scheduleUpdateNext,
  scheduleDisable,
  scheduleEnable,
  scheduleDelete, scheduleUpdateSequence
} from "../../api/schedule";
import './schedule.css'
import {useSelector} from "react-redux";
import {ColorButtonProvider} from "../../theme/button";
import dayjs from "dayjs";
import {typePeriodTypeFind, typeSendTypeFind} from "../../api/type";
import {channelFind} from "../../api/channel";
import {
  PeriodTypeDaily, PeriodTypeDayInterval,
  PeriodTypeHour, PeriodTypeHourInterval,
  PeriodTypeMinute, PeriodTypeMinuteInterval, PeriodTypeMonthInterval,
  PeriodTypeMonthly, PeriodTypeQuarterInterval, PeriodTypeQuarterly,
  PeriodTypeSecond, PeriodTypeSecondInterval, PeriodTypeWeekInterval,
  PeriodTypeWeekly, PeriodTypeYearInterval, PeriodTypeYearly
} from "./periodType";

const Schedule = () => {
  const mode = useSelector(state => state.mode)
  const [searchKeyword, setSearchKeyword] = useState({
    search: ''
  })
  const [updateOrder, setUpdateOrder] = useState(false)
  const [showDisabled, setShowDisabled] = useState(true)
  const [sendTypeData, setSendTypeData] = useState([])
  const [periodTypeData, setPeriodTypeData] = useState([])
  const [selectedPeriodType, setSelectedPeriodType] = useState(0)
  const [channelData, setChannelData] = useState([])
  const [tableData, setTableData] = useState([])
  const [tableFilterData, setTableFilterData] = useState([])
  // Create
  const [inputScheduleDataAction, setInputScheduleDataAction] = useState('close');
  const [form] = Form.useForm()

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
    typePeriodTypeFind().then(({data}) => {
      if (data.code === 0) {
        let options = []
        data.data.forEach((item) => {
          options.push({
            value: item.type,
            label: item.name
          })
        })
        setPeriodTypeData(options)
      }
    })
  }, [])

  useEffect(() => {
    channelFind().then(({data}) => {
      if (data.code === 0) {
        let options = []
        data.data.forEach((item) => {
          options.push({
            value: item.id,
            label: item.name + '(' + item.bot + '/' + item.type + ')'
          })
        })
        setChannelData(options)
      }
    })
  }, [])

  const updateTable = ({search}) => {
    scheduleFind({search: search, all: true}).then(({data}) => {
      if (data.code === 0) {
        setTableData(data.data)
        setTableFilterData(JSON.parse(JSON.stringify(data.data)).filter(item => item.disabled === 0))
      }
    })
  }
  const getTableData = () => {
    if (showDisabled) {
      return tableData
    } else {
      return tableFilterData
    }
  }

  const handleUpdateNext = (id) => {
    scheduleUpdateNext({id: id}).then(({data}) => {
      if (data.code !== 0) {
        message.open({
          type: 'warning',
          content: '更新失败: ' + data.msg
        })
      } else {
        message.open({
          type: 'success',
          content: '更新成功'
        })
      }
      updateTable(searchKeyword)
    })
  }
  const updateScheduleSequence = (id, target) => {
    scheduleUpdateSequence({id: id, target: target}).then(({data}) => {
      if (data.code === 0) {
        updateTable(searchKeyword)
      }
    })
  }

  const handleDeleteSchedule = (id) => {
    scheduleDelete({id: id}).then(({data}) => {
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

  const handleDisableSchedule = (disable, id) => {
    if (disable) {
      scheduleDisable({id: id}).then(({data}) => {
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
      scheduleEnable({id: id}).then(({data}) => {
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

  const truncateWithEllipsis = (str, length) => {
    const truncated = Array.from(str).slice(0, length).join('');
    return str.length > length ? truncated + '...' : truncated;
  }
  const handleSearchSchedule = ({keyword}) => {
    setSearchKeyword({
      search: keyword
    })
  }
  const handleInputScheduleDataShow = (action, data) => {
    if (data) {
      const cloneData = JSON.parse(JSON.stringify(data))
      cloneData.expiration_date_date = dayjs(cloneData.expiration_date * 1000)
      cloneData.expiration_date_time = dayjs(cloneData.expiration_date * 1000)
      cloneData.start_at_date = dayjs(cloneData.start_at * 1000)
      cloneData.start_at_time = dayjs(cloneData.start_at * 1000)
      setSelectedPeriodType(cloneData.period_type)
      form.setFieldsValue(cloneData)
    } else {
      const cloneData = {}
      cloneData.expiration_date_date = dayjs('2100-01-01 00:00:00')
      cloneData.expiration_date_time = dayjs('2100-01-01 00:00:00')
      cloneData.start_at_date = dayjs()
      cloneData.start_at_time = dayjs()
      cloneData.expiration_times = -1
      form.setFieldsValue(cloneData)
    }
    setInputScheduleDataAction(action)
  }

  const handleInputScheduleDataOk = () => {
    form.validateFields().then((input) => {
      input.expiration_date = dayjs(input.expiration_date_date.format('YYYY-MM-DD') + ' ' + input.expiration_date_time.format('HH:mm:ss')).unix()
      input.start_at = dayjs(input.start_at_date.format('YYYY-MM-DD') + ' ' + input.start_at_time.format('HH:mm:ss')).unix()
      delete input.expiration_date_date
      delete input.expiration_date_time
      delete input.start_at_date
      delete input.start_at_time
      if (inputScheduleDataAction === 'create') {
        scheduleCreate(input).then(({data}) => {
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
          setInputScheduleDataAction('close')
          form.resetFields()
        })
      } else if (inputScheduleDataAction === 'update') {
        scheduleUpdate(input).then(({data}) => {
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
          setInputScheduleDataAction('close')
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
      title: '标题',
      dataIndex: 'title'
    }, {
      title: '分类',
      dataIndex: 'category',
    }, {
      title: '信息类型',
      dataIndex: 'type',
      render: (type) => {
        let cType = sendTypeData.find(item => item.value === type)
        if (cType) {
          return cType.label
        }
        return type
      }
    }, {
      title: '内容',
      key: 'message',
      dataIndex: 'message',
      render: (msg) => {
        if (msg.length === 0) {
          return '-'
        }
        return truncateWithEllipsis(msg, 16)
      }
    }, {
      title: '渠道',
      key: 'channel',
      dataIndex: 'channel',
      render: (channel) => {
        return channel.name + '(' + channel.bot + '/' + channel.type + ')'
      }
    }, {
      title: '开始时间',
      dataIndex: 'start_at',
      render: (start_at) => {
        if (start_at === 0) {
          return '-'
        }
        return dayjs(start_at * 1000).format('YYYY-MM-DD HH:mm:ss')
      }
    }, {
      title: '下次发送',
      dataIndex: 'next_of_period',
      render: (next_of_period) => {
        if (next_of_period === 0) {
          return '-'
        }
        return dayjs(next_of_period * 1000).format('YYYY-MM-DD HH:mm:ss')
      }
    }, {
      title: '过期时间',
      dataIndex: 'expiration_date',
      render: (expiration_date) => {
        if (expiration_date === 0) {
          return '-'
        }
        return dayjs(expiration_date * 1000).format('YYYY-MM-DD HH:mm:ss')
      }
    }, {
      title: '剩余提醒次数',
      dataIndex: 'expiration_times',
      render: (expiration_times) => {
        if (expiration_times < 0) {
          return '-'
        }
        return expiration_times
      }
    }, {
      title: '周期',
      key: 'period_type',
      render: (rowData) => {
        let msg = ""
        let cType = periodTypeData.find(item => item.value === rowData.period_type)
        if (cType) {
          msg = cType.label
        }
        msg += " | "
        switch (rowData.period_type) {
          case PeriodTypeSecond:
            msg += "-"
            break
          case PeriodTypeMinute:
            msg += rowData.second + "秒"
            break
          case PeriodTypeHour:
            msg += rowData.minute + "分 " + rowData.second + "秒"
            break
          case PeriodTypeDaily:
            msg += rowData.hour + "时 " + rowData.minute + "分 " + rowData.second + "秒"
            break
          case PeriodTypeWeekly:
            msg += rowData.day + "日 " + rowData.hour + "时 " + rowData.minute + "分 " + rowData.second + "秒"
            break
          case PeriodTypeMonthly:
            msg += rowData.day + "日 " + rowData.hour + "时 " + rowData.minute + "分 " + rowData.second + "秒"
            break
          case PeriodTypeQuarterly:
            msg += rowData.month + "月 " + rowData.day + "日 " + rowData.hour + "时 " + rowData.minute + "分 " + rowData.second + "秒"
            break
          case PeriodTypeYearly:
            msg += rowData.month + "月 " + rowData.day + "日 " + rowData.hour + "时 " + rowData.minute + "分 " + rowData.second + "秒"
            break
          case PeriodTypeSecondInterval:
            msg += rowData.second + "秒"
            break
          case PeriodTypeMinuteInterval:
            msg += rowData.minute + "分 " + rowData.second + "秒"
            break
          case PeriodTypeHourInterval:
            msg += rowData.hour + "时 " + rowData.minute + "分 " + rowData.second + "秒"
            break
          case PeriodTypeDayInterval:
            msg += rowData.day + "日 " + rowData.hour + "时 " + rowData.minute + "分 " + rowData.second + "秒"
            break
          case PeriodTypeWeekInterval:
            msg += rowData.week + "周 " + rowData.day + "日 " + rowData.hour + "时 " + rowData.minute + "分 " + rowData.second + "秒"
            break
          case PeriodTypeMonthInterval:
            msg += rowData.month + "月 " + rowData.day + "日 " + rowData.hour + "时 " + rowData.minute + "分 " + rowData.second + "秒"
            break
          case PeriodTypeQuarterInterval:
            msg += rowData.quarter + "季 " + rowData.month + "月 " + rowData.day + "日 " + rowData.hour + "时 " + rowData.minute + "分 " + rowData.second + "秒"
            break
          case PeriodTypeYearInterval:
            msg += rowData.year + "年 " + rowData.month + "月 " + rowData.day + "日 " + rowData.hour + "时 " + rowData.minute + "分 " + rowData.second + "秒"
            break
          default:
            break
        }
        return msg
      }
    }, {
      title: '操作',
      render: (rowData) => {
        let disableStatus = rowData.disabled === 0
        let disableMsg = disableStatus ? '停用' : '启用'
        return (
          <div>
            <ColorButtonProvider color="green">
              <Popconfirm
                title={'更新日程提醒'}
                description={"你确定更新日程提醒?"}
                onConfirm={() => handleUpdateNext(rowData.id)}
                onCancel={() => {
                }}
                okText="确认"
                cancelText="取消"
              >
                <Button type="primary" style={{marginRight: '5px'}}>更新</Button>
              </Popconfirm>
            </ColorButtonProvider>
            <Button type="primary" style={{marginRight: '5px'}}
                    onClick={() => handleInputScheduleDataShow('update', rowData)}>编辑</Button>
            <ColorButtonProvider danger={disableStatus} color="green">
              <Popconfirm
                title={disableMsg + '日程提醒'}
                description={"你确定" + disableMsg + "日程提醒?"}
                onConfirm={() => handleDisableSchedule(disableStatus, rowData.id)}
                onCancel={() => {
                }}
                okText="确认"
                cancelText="取消"
              >
                <Button danger={disableStatus} type="primary"
                        style={{marginRight: '5px'}}>{disableMsg}</Button>
              </Popconfirm>
            </ColorButtonProvider>
            <Popconfirm
              title={'删除日程提醒'}
              description={"你确定删除" + rowData.name + "?"}
              onConfirm={() => handleDeleteSchedule(rowData.id)}
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
  if (updateOrder) {
    columns.push({
      title: '排序',
      render: (rowData) => {
        return (
          <Flex style={{width: '100%'}} justify='flex-start' align='flex-start'>
            <Button type="default" style={{marginRight: '5px'}}>{rowData.sequence}</Button>
            <Button type="default" style={{marginRight: '5px'}}
                    onClick={() => updateScheduleSequence(rowData.id, rowData.sequence - 1)}>上</Button>
            <Button type="default"
                    onClick={() => updateScheduleSequence(rowData.id, rowData.sequence + 1)}>下</Button>
          </Flex>
        );
      }
    })
  }

  useEffect(() => {
    updateTable(searchKeyword)
  }, [searchKeyword]);

  return (
    <div>
      {
        mode.isWide ? (
          <Flex style={{width: '100%', marginBottom: '15px'}} justify='space-between' align='center'>
            <div style={{marginBottom: 6}}>
              <Button type="primary" onClick={() => handleInputScheduleDataShow('create')}
                      style={{marginRight: 6}}>创建</Button>
              <Button type="default" onClick={() => setUpdateOrder(!updateOrder)}
                      style={{marginRight: 6}}>排序</Button>
              <Button type="default" onClick={() => setShowDisabled(!showDisabled)}>隐藏</Button>
            </div>
            <Form
              layout="inline"
              onFinish={handleSearchSchedule}
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
              <Button type="primary" onClick={() => handleInputScheduleDataShow('create')}>创建</Button>
              <Button type="default" onClick={() => setUpdateOrder(!updateOrder)}
                      style={{marginRight: 6}}>排序</Button>
              <Button type="default" onClick={() => setShowDisabled(!showDisabled)}>隐藏</Button>
            </div>
            <Form
              layout="inline"
              onFinish={handleSearchSchedule}
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
        dataSource={getTableData()}
        pagination={{
          pageSizeOptions: [10, 15, 20, 50, 100],
          responsive: true,
          showQuickJumper: true,
          showSizeChanger: true
        }}
        expandable={{
          expandedRowRender: (record) => {
            let title = ''
            let msg = ''
            if (record.title.length !== 0) {
              title = <p style={{margin: 0}}><strong>{record.title}</strong></p>
            }
            if (record.message.length !== 0) {
              msg = <p style={{margin: 0}}>{record.message}</p>
            }
            return (
              <div style={{margin: 0}}>
                {title}
                {msg}
              </div>
            )
          },
          expandIcon: ({expanded, onExpand, record}) => {
            return expanded ? (
                <Button type="default" onClick={e => onExpand(record, e)}>收起</Button>
              ) :
              (
                <Button type="default" onClick={e => onExpand(record, e)}>展开</Button>
              )
          },
          rowExpandable: (record) => true,
        }}
        scroll={{
          x: 'max-content',
        }}
        rowKey={'id'}
      />
      <Modal
        title={inputScheduleDataAction === 'create' ? '创建' : inputScheduleDataAction === 'update' ? '编辑' : 'Unknown'}
        open={inputScheduleDataAction !== 'close'}
        onOk={handleInputScheduleDataOk}
        onCancel={() => {
          setInputScheduleDataAction('close');
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
            inputScheduleDataAction === 'update' &&
            <Form.Item
              name="id"
              hidden
            >
              <Input/>
            </Form.Item>
          }
          <Form.Item
            label="分类"
            name="category"
            rules={[
              {
                required: true,
                message: '请输入分类'
              }
            ]}
          >
            <Input placeholder={'请输入分类'}/>
          </Form.Item>
          <Form.Item
            label="通知类型"
            name="type"
            rules={[
              {
                required: true,
                message: '请输入通知类型'
              }
            ]}
          >
            <Select
              placeholder="通知类型"
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
                required: true,
                message: '请输入标题'
              }
            ]}
          >
            <Input placeholder={'请输入标题'}/>
          </Form.Item>
          <Form.Item
            label="信息"
            name="message"
            rules={[
              {
                required: true,
                message: '请输入信息'
              }
            ]}
          >
            <Input placeholder={'请输入信息'}/>
          </Form.Item>
          <Form.Item
            label="通知渠道"
            name="channel_id"
            rules={[
              {
                required: true,
                message: '请输入通知渠道'
              }
            ]}
          >
            <Select
              placeholder="通知渠道"
              onChange={() => {
              }}
              allowClear
              options={channelData}
            />
          </Form.Item>
          <Form.Item
            label="过期日期"
            name="expiration_date_date"
            rules={[
              {
                required: true,
                message: '请输入过期日期'
              }
            ]}
          >
            <DatePicker picker='date' inputReadOnly style={{width: '100%'}} placeholder={'请输入过期日期'}/>
          </Form.Item>
          <Form.Item
            label="过期时间"
            name="expiration_date_time"
            rules={[
              {
                required: true,
                message: '请输入过期时间'
              }
            ]}
          >
            <DatePicker picker='time' inputReadOnly style={{width: '100%'}} placeholder={'请输入过期时间'}/>
          </Form.Item>
          <Form.Item
            label="截止次数"
            name="expiration_times"
            rules={[
              {
                required: true,
                message: '请输入截止次数'
              }
            ]}
          >
            <InputNumber placeholder={'-1代表无限制'} style={{width: '100%'}}/>
          </Form.Item>
          <Form.Item
            label="周期类型"
            name="period_type"
            rules={[
              {
                required: true,
                message: '请输入周期类型'
              }
            ]}
          >
            <Select
              placeholder="周期类型"
              onChange={(t) => {
                setSelectedPeriodType(t)
              }}
              allowClear
              options={periodTypeData}
            />
          </Form.Item>
          <Form.Item
            label="开始日期"
            name="start_at_date"
            rules={[
              {
                required: true,
                message: '请输入开始日期'
              }
            ]}
          >
            <DatePicker picker='date' inputReadOnly style={{width: '100%'}} placeholder={'请输入开始日期'}/>
          </Form.Item>
          <Form.Item
            label="开始时间"
            name="start_at_time"
            rules={[
              {
                required: true,
                message: '请输入开始时间'
              }
            ]}
          >
            <DatePicker picker='time' inputReadOnly style={{width: '100%'}} placeholder={'请输入开始时间'}/>
          </Form.Item>
          {
            (
              selectedPeriodType === PeriodTypeYearInterval
            ) &&
            <Form.Item
              label="年"
              name="year"
              rules={[
                {
                  required: true,
                  message: '请输入年'
                }
              ]}
            >
              <InputNumber placeholder={'负数表示倒数'} style={{width: '100%'}}/>
            </Form.Item>
          }
          {
            (
              selectedPeriodType === PeriodTypeQuarterInterval
            ) &&
            <Form.Item
              label="季"
              name="quarter"
              rules={[
                {
                  required: true,
                  message: '请输入季'
                }
              ]}
            >
              <InputNumber placeholder={'负数表示倒数'} style={{width: '100%'}}/>
            </Form.Item>
          }
          {
            (
              selectedPeriodType === PeriodTypeQuarterly ||
              selectedPeriodType === PeriodTypeYearly ||
              selectedPeriodType === PeriodTypeMonthInterval ||
              selectedPeriodType === PeriodTypeQuarterInterval ||
              selectedPeriodType === PeriodTypeYearInterval
            ) &&
            <Form.Item
              label="月"
              name="month"
              rules={[
                {
                  required: true,
                  message: '请输入月'
                }
              ]}
            >
              <InputNumber placeholder={'负数表示倒数'} style={{width: '100%'}}/>
            </Form.Item>
          }
          {
            (
              selectedPeriodType === PeriodTypeWeekInterval
            ) &&
            <Form.Item
              label="周"
              name="week"
              rules={[
                {
                  required: true,
                  message: '请输入周'
                }
              ]}
            >
              <InputNumber placeholder={'负数表示倒数'} style={{width: '100%'}}/>
            </Form.Item>
          }
          {
            (
              selectedPeriodType === PeriodTypeWeekly ||
              selectedPeriodType === PeriodTypeMonthly ||
              selectedPeriodType === PeriodTypeQuarterly ||
              selectedPeriodType === PeriodTypeYearly ||
              selectedPeriodType === PeriodTypeDayInterval ||
              selectedPeriodType === PeriodTypeWeekInterval ||
              selectedPeriodType === PeriodTypeMonthInterval ||
              selectedPeriodType === PeriodTypeQuarterInterval ||
              selectedPeriodType === PeriodTypeYearInterval
            ) &&
            <Form.Item
              label="日"
              name="day"
              rules={[
                {
                  required: true,
                  message: '请输入日'
                }
              ]}
            >
              <InputNumber placeholder={'负数表示倒数'} style={{width: '100%'}}/>
            </Form.Item>
          }
          {
            (
              selectedPeriodType === PeriodTypeDaily ||
              selectedPeriodType === PeriodTypeWeekly ||
              selectedPeriodType === PeriodTypeMonthly ||
              selectedPeriodType === PeriodTypeQuarterly ||
              selectedPeriodType === PeriodTypeYearly ||
              selectedPeriodType === PeriodTypeHourInterval ||
              selectedPeriodType === PeriodTypeDayInterval ||
              selectedPeriodType === PeriodTypeWeekInterval ||
              selectedPeriodType === PeriodTypeMonthInterval ||
              selectedPeriodType === PeriodTypeQuarterInterval ||
              selectedPeriodType === PeriodTypeYearInterval
            ) &&
            <Form.Item
              label="时"
              name="hour"
              rules={[
                {
                  required: true,
                  message: '请输入时'
                }
              ]}
            >
              <InputNumber placeholder={'负数表示倒数'} style={{width: '100%'}}/>
            </Form.Item>
          }
          {
            (
              selectedPeriodType === PeriodTypeHour ||
              selectedPeriodType === PeriodTypeDaily ||
              selectedPeriodType === PeriodTypeWeekly ||
              selectedPeriodType === PeriodTypeMonthly ||
              selectedPeriodType === PeriodTypeQuarterly ||
              selectedPeriodType === PeriodTypeYearly ||
              selectedPeriodType === PeriodTypeMinuteInterval ||
              selectedPeriodType === PeriodTypeHourInterval ||
              selectedPeriodType === PeriodTypeDayInterval ||
              selectedPeriodType === PeriodTypeWeekInterval ||
              selectedPeriodType === PeriodTypeMonthInterval ||
              selectedPeriodType === PeriodTypeQuarterInterval ||
              selectedPeriodType === PeriodTypeYearInterval
            ) &&
            <Form.Item
              label="分"
              name="minute"
              rules={[
                {
                  required: true,
                  message: '请输入分'
                }
              ]}
            >
              <InputNumber placeholder={'负数表示倒数'} style={{width: '100%'}}/>
            </Form.Item>
          }
          {
            (
              selectedPeriodType === PeriodTypeMinute ||
              selectedPeriodType === PeriodTypeHour ||
              selectedPeriodType === PeriodTypeDaily ||
              selectedPeriodType === PeriodTypeWeekly ||
              selectedPeriodType === PeriodTypeMonthly ||
              selectedPeriodType === PeriodTypeQuarterly ||
              selectedPeriodType === PeriodTypeYearly ||
              selectedPeriodType === PeriodTypeSecondInterval ||
              selectedPeriodType === PeriodTypeMinuteInterval ||
              selectedPeriodType === PeriodTypeHourInterval ||
              selectedPeriodType === PeriodTypeDayInterval ||
              selectedPeriodType === PeriodTypeWeekInterval ||
              selectedPeriodType === PeriodTypeMonthInterval ||
              selectedPeriodType === PeriodTypeQuarterInterval ||
              selectedPeriodType === PeriodTypeYearInterval
            ) &&
            <Form.Item
              label="秒"
              name="second"
              rules={[
                {
                  required: true,
                  message: '请输入秒'
                }
              ]}
            >
              <InputNumber placeholder={'负数表示倒数'} style={{width: '100%'}}/>
            </Form.Item>
          }
        </Form>
      </Modal>
    </div>
  )
}

export default Schedule
