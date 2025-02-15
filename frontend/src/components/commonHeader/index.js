import React, {useEffect, useRef, useState} from 'react'
import {Layout, theme, Button, Avatar, Dropdown, Space, Tooltip, Flex} from "antd";
import './index.css';
import {MenuFoldOutlined} from "@ant-design/icons";
import {useDispatch} from "react-redux";
import {collapseMenu} from "../../store/reducers/tab";
import {switchMode} from "../../store/reducers/mode";
import {useNavigate} from "react-router-dom";
import {userInfo, userLogout} from "../../api/user";

const {Header} = Layout;

const CommonHeader = ({mode}) => {
  const navigate = useNavigate()
  const navigation = useRef(navigate);
  const {
    token: {colorBgContainer},
  } = theme.useToken();

  const dispath = useDispatch()
  const setCollapsed = () => {
    dispath(collapseMenu())
  }
  const setMode = () => {
    dispath(switchMode())
  }

  const logout = () => {
    userLogout()
    localStorage.clear()
    navigate('/login')
  }
  const [user, setUserData] = useState(JSON.parse(localStorage.getItem('user')))
  useEffect(() => {
    userInfo().then(({data}) => {
      if (data.code === 0) {
        localStorage.setItem('role', data.data.role)
        localStorage.setItem('user', JSON.stringify(data.data))
        setUserData(data.data)
      } else {
        userLogout()
        localStorage.clear()
        navigation.current.navigate('/login')
      }
    })
  }, [navigation])
  const items = [
    {
      key: '1',
      label: (
        <div onClick={() => {
          navigate('/admin/user-center')
        }}>
          个人中心
        </div>
      ),
    },
    {
      key: '2',
      label: (
        <div onClick={setMode}>
          {mode.display}
        </div>
      ),
    },
    {
      key: '3',
      label: (
        <div onClick={logout}>
          退出
        </div>
      ),
    }
  ];
  return (
    <Header
      style={{
        background: colorBgContainer,
        position: 'fixed',
        top: 0,
        zIndex: 10,
        width: '100%',
        height: 64,
        display: 'flex',
        alignItems: 'center',
      }}
    >
      <Flex style={{width: '100%'}} justify='space-between' align='center'>
        <Button
          type="text"
          icon={<MenuFoldOutlined/>}
          style={{
            fontSize: '16px',
            backgroundColor: colorBgContainer
          }}
          onClick={() => {
            setCollapsed()
          }}
        />
        <Space>
          <Tooltip>
            <span>{user.nickname}</span>
          </Tooltip>
          <Dropdown
            menu={{items}}
          >
            <Avatar size="large"
                    src={user.avatar ? user.avatar : ''}/>
          </Dropdown>
        </Space>
      </Flex>
    </Header>
  )
}

export default CommonHeader
