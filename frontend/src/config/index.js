const config = [
  {
    path: '/home',
    name: 'home',
    label: '首页',
    role: 'viewer',
    icon: 'CalendarOutlined',
    url: '/home/index'
  },
  {
    path: '/channel',
    name: 'channel',
    label: '发送渠道',
    role: 'viewer',
    icon: 'SendOutlined'
  },
  {
    path: '/send',
    name: 'send',
    label: '消息历史',
    role: 'viewer',
    icon: 'HistoryOutlined'
  },
  {
    path: '/admin',
    label: '管理',
    role: 'viewer',
    icon: 'SettingOutlined',
    children: [
      {
        path: '/admin/user-center',
        name: 'user-center',
        label: '用户中心',
        role: 'viewer',
        icon: 'UserOutlined'
      },
      {
        path: '/admin/user',
        name: 'user',
        label: '用户',
        role: 'admin',
        icon: 'UsergroupAddOutlined'
      },
      {
        path: '/admin/access_token',
        name: 'access_token',
        label: '访问密钥',
        role: 'user',
        icon: 'KeyOutlined'
      }
    ]
  }
]

const AdminMenu = JSON.parse(JSON.stringify(config)).filter(() => {
  return true
})

const UserMenu = JSON.parse(JSON.stringify(config)).filter(val => val.role !== 'admin').map(item => {
  if (item.children) {
    item.children = item.children.filter(val => {
      return val.role !== 'admin'
    })
  }
  return item
})

const ViewerMenu = JSON.parse(JSON.stringify(config)).filter(val => val.role !== 'admin' && val.role !== 'user').map(item => {
  if (item.children) {
    item.children = item.children.filter(val => {
      return val.role !== 'admin' && val.role !== 'user'
    })
  }
  return item
})

export {AdminMenu, UserMenu, ViewerMenu}
