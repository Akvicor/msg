import React, {useEffect, useState} from 'react'
import {Button, Form, Input, Table, Popconfirm, message, Flex, Tag} from 'antd'
import {sendFind, sendCancel} from "../../api/send";
import './send.css'
import {useSelector} from "react-redux";
import {typeBotSendersFind, typeSendTypeFind, typeChannelTypeFind} from "../../api/type";
import dayjs from "dayjs";
import {channelFind} from "../../api/channel";

const Send = () => {
  const mode = useSelector(state => state.mode)
  const [searchKeyword, setSearchKeyword] = useState({
    search: '',
    sent: [],
    channel_ids: [],
    types: []
  })
  const [sendTypeData, setSendTypeData] = useState([])
  const [channelTypeData, setChannelTypeData] = useState([])
  const [botSendersData, setBotSendersData] = useState([])
  const [tableData, setTableData] = useState([])
  const [tableLoadingData, setTableLoadingData] = useState(false)
  const [tablePaginationData, setTablePaginationData] = useState({current: 1, pageSize: 10})
  const [tablePaginationRespData, setTablePaginationRespData] = useState({total: 0})
  const [tableFilterChannelData, setTableFilterChannelData] = useState([])
  const [tableFilterTypesData, setTableFilterTypesData] = useState([])

  useEffect(() => {
    typeSendTypeFind().then(({data}) => {
      if (data.code === 0) {
        let options = []
        let filters = []
        data.data.forEach((item) => {
          options.push({
            value: item.type,
            label: item.name
          })
          filters.push({
            value: item.type,
            text: item.name
          })
        })
        setSendTypeData(options)
        setTableFilterTypesData(filters)
      }
    })
  }, [])

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
    channelFind().then(({data}) => {
      if (data.code === 0) {
        let filters = []
        data.data.forEach((item) => {
          filters.push({
            value: item.id,
            text: item.name
          })
        })
        setTableFilterChannelData(filters)
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
          options.push(opt)
        })
        setBotSendersData(options)
      }
    })
  }, [])

  const updateTable = () => {
    setTableLoadingData(true)
    sendFind({
      search: searchKeyword.search,
      sent: searchKeyword.sent,
      channel_ids: searchKeyword.channel_ids,
      types: searchKeyword.types,
      index: tablePaginationData.current,
      limit: tablePaginationData.pageSize,
    }).then(({data}) => {
      if (data.code === 0) {
        setTableData(data.data)
        setTableLoadingData(false)
        setTablePaginationRespData({
          total: data.page.total,
        })
      }
    })
  }

  const handleCancelSend = (id) => {
    sendCancel({id: id}).then(({data}) => {
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
      updateTable()
    })
  }
  const handleSearchSend = ({keyword}) => {
    const cloneData = JSON.parse(JSON.stringify(searchKeyword))
    cloneData.search = keyword
    setSearchKeyword(cloneData)
  }
  const truncateWithEllipsis = (str, length) => {
    const truncated = Array.from(str).slice(0, length).join('');
    return str.length > length ? truncated + '...' : truncated;
  }

  const columns = [
    {
      title: '渠道',
      key: 'channel',
      dataIndex: 'channel',
      render: (channel) => {
        return channel.name
      },
      filters: tableFilterChannelData,
      onFilter: (value, record) => true
    }, {
      title: 'Bot',
      key: 'bot',
      dataIndex: 'channel',
      render: (channel) => {
        let res = ""
        let cBot = botSendersData.find(item => item.value === channel.bot)
        if (cBot) {
          res += cBot.label
        } else {
          res += channel.bot
        }
        return res
      }
    }, {
      title: '类型',
      key: 'channel_type',
      dataIndex: 'channel',
      render: (channel) => {
        let res = ""
        let cType = channelTypeData.find(item => item.value === channel.type)
        if (cType) {
          res += cType.label
        } else {
          res += channel.type
        }
        return res
      }
    }, {
      title: '类型',
      dataIndex: 'type',
      render: (type) => {
        let cType = sendTypeData.find(item => item.value === type)
        if (cType) {
          return cType.label
        }
        return type
      },
      filters: tableFilterTypesData,
      onFilter: (value, record) => true
    }, {
      title: '标题',
      key: 'title',
      dataIndex: 'title'
    }, {
      title: '内容',
      key: 'msg',
      dataIndex: 'msg',
      render: (msg) => {
        if (msg.length === 0) {
          return '-'
        }
        return truncateWithEllipsis(msg, 16)
      }
    }, {
      title: '调用IP',
      dataIndex: 'ip',
      render: (ip) => {
        if (ip.length === 0) {
          return '-'
        }
        return ip
      }
    }, {
      title: '创建时间',
      dataIndex: 'ready_at',
      render: (ready_at) => {
        if (ready_at === 0) {
          return '-'
        }
        return dayjs(ready_at * 1000).format('YYYY-MM-DD HH:mm:ss')
      }
    }, {
      title: '发送时间',
      dataIndex: 'send_at',
      render: (send_at) => {
        if (send_at === 0) {
          return '-'
        }
        return dayjs(send_at * 1000).format('YYYY-MM-DD HH:mm:ss')
      }
    }, {
      title: '送达时间',
      dataIndex: 'sent_at',
      render: (sent_at) => {
        let color = 'green'
        let status = 'success'
        if (sent_at < 0) {
          sent_at = -sent_at
          color = 'red'
          status = 'cancel'
        } else if (sent_at === 0) {
          color = 'yellow'
          status = 'waiting'
          return (
            <div>
              <Tag color={color} key={sent_at}>
                {status}
              </Tag>
            </div>
          );
        }
        return (
          <div>
            <Tag color={color} key={sent_at}>
              {status}
            </Tag>
            {dayjs(sent_at * 1000).format('YYYY-MM-DD HH:mm:ss')}
          </div>
        );
      },
      filters: [
        {value: 1, text: 'success'},
        {value: 0, text: 'waiting'},
        {value: -1, text: 'cancel'}
      ],
      onFilter: (value, record) => true
    }, {
      title: '操作',
      render: (rowData) => {
        if (rowData.sent_at !== 0) {
          return (
            <div></div>
          )
        }
        return (
          <div>
            <Popconfirm
              title={'取消发送'}
              description={"你确定取消发送?"}
              onConfirm={() => handleCancelSend(rowData.id)}
              onCancel={() => {
              }}
              okText="确认"
              cancelText="取消"
            >
              <Button danger type="primary">取消</Button>
            </Popconfirm>
          </div>
        );
      }
    }
  ]

  useEffect(updateTable, [tablePaginationData, searchKeyword]);

  return (
    <div>
      {
        mode.isWide ? (
          <Flex style={{width: '100%', marginBottom: '15px'}} justify='space-between' align='center'>
            <Button type="primary" onClick={() => updateTable(searchKeyword)}>刷新</Button>
            <Form
              layout="inline"
              onFinish={handleSearchSend}
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
              <Button type="primary" onClick={() => updateTable(searchKeyword)}>刷新</Button>
            </div>
            <Form
              layout="inline"
              onFinish={handleSearchSend}
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
          total: tablePaginationRespData.total,
          current: tablePaginationData.current,
          pageSize: tablePaginationData.pageSize,
          loading: tableLoadingData,
          showQuickJumper: true,
          showSizeChanger: true,
          pageSizeOptions: [5, 10, 15, 20, 50, 100],
          responsive: true,
          onChange: (pageNumber, pageSize) => {
            let data = JSON.parse(JSON.stringify(tablePaginationData))
            data.current = pageNumber
            data.pageSize = pageSize
            setTablePaginationData(data)
          },
          showTotal: (total) => {
            return 'Total ' + total + ' items'
          }
        }}
        expandable={{
          expandedRowRender: (record) => {
            let title = ''
            let msg = ''
            let err_msg = ''
            if (record.title.length !== 0) {
              title = <p style={{margin: 0}}><strong>{record.title}</strong></p>
            }
            if (record.msg.length !== 0) {
              msg = <p style={{margin: 0}}>{record.msg}</p>
            }
            if (record.err_msg.length !== 0) {
              err_msg = <p style={{margin: 0, color: "red"}}><strong>{record.err_msg}</strong></p>
            }
            return (
              <div style={{margin: 0}}>
                {title}
                {msg}
                {err_msg}
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
        onChange={
          (pagination, filters, sorter, extra) => {
            const cloneData = JSON.parse(JSON.stringify(searchKeyword))
            cloneData.sent = filters.sent_at
            cloneData.channel_ids = filters.channel
            cloneData.types = filters.type
            setSearchKeyword(cloneData)
          }
        }
        scroll={{
          x: 'max-content',
        }}
        rowKey={'id'}
      />
    </div>
  )
}

export default Send
