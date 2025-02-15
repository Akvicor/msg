import React, {useEffect, useState} from 'react'
import {AdminMenu, UserMenu, ViewerMenu} from '../../config'
import {Layout, Menu, theme} from "antd";
import * as Icon from '@ant-design/icons';
import {useDispatch} from "react-redux";
import {closeMenu, collapseMenu} from "../../store/reducers/tab";
import {useLocation, useNavigate} from "react-router-dom";


const {Sider} = Layout;
const iconToElement = (name) => React.createElement(Icon[name])

const CommonSider = ({collapsed, mode}) => {
  const navigate = useNavigate()
  const dispath = useDispatch()
  const {
    token: {colorBgContainer},
  } = theme.useToken();
  const [selectedMenu, setSelectedMenuData] = useState({
    OpenKeys: [],
    SelectedKeys: ['/home'],
  })
  const setCollapsed = () => {
    dispath(collapseMenu())
  }

  let pathname = useLocation().pathname
  useEffect(() => {
    const menuRank = pathname.split('/')
    let data = {
      OpenKeys: [],
      SelectedKeys: ['/home'],
    }
    switch (menuRank.length) {
      case 2:
        data = {
          OpenKeys: [],
          SelectedKeys: [pathname],
        }
        break
      case 3:
        data = {
          OpenKeys: [menuRank.slice(0, 2).join('/')],
          SelectedKeys: [pathname],
        }
        break
      case 4:
        data = {
          OpenKeys: [menuRank.slice(0, 2).join('/'), menuRank.slice(0, 3).join('/')],
          SelectedKeys: [pathname],
        }
        break
      default:
        break
    }
    setSelectedMenuData(data)
  }, [pathname])


  const role = localStorage.getItem('role')
  let menu
  if (role === 'admin') {
    menu = AdminMenu
  } else if (role === 'user') {
    menu = UserMenu
  } else {
    menu = ViewerMenu
  }

  const items = menu.map((item) => {
    const child = {
      key: item.path,
      icon: iconToElement(item.icon),
      label: item.label
    }
    if (item.children) {
      child.children = item.children.map(item => {
        return {
          key: item.path,
          icon: iconToElement(item.icon),
          label: item.label
        }
      })
    }
    return child
  })
  const selectMenu = (e) => {
    let data = {
      OpenKeys: [],
      SelectedKeys: ['/home'],
    }
    menu.forEach(item => {
      if (item.path === e.keyPath[e.keyPath.length - 1]) {
        if (e.keyPath.length > 1) {
          let result = item.children.find(child => {
            return child.path === e.key
          })
          data.OpenKeys = [item.path]
          data.SelectedKeys = [result.path]
        } else {
          data.SelectedKeys = [item.path]
        }
      }
    })
    setSelectedMenuData(data)
    navigate(e.key)
    if (!mode.isWide) {
      dispath(closeMenu())
    }
  }
  const openMenu = (keys) => {
    let data = {
      OpenKeys: keys,
      SelectedKeys: selectedMenu.SelectedKeys,
    }
    setSelectedMenuData(data)
  }
  let menuTop = collapsed ? 0 : 64
  return (
    <Sider
      style={{
        top: 0,
        background: colorBgContainer,
      }}
      collapsedWidth={0}
      collapsed={collapsed}
      onCollapse={() => setCollapsed()}
    >
      {/*<h3 style={{background: colorBgContainer, textAlign: "center", color: "#001529"}}>{'Wallet'}</h3>*/}
      <Menu
        style={{
          top: menuTop,
          overflow: 'auto',
          position: 'sticky',
          background: colorBgContainer
        }}
        mode="inline"
        selectedKeys={selectedMenu.SelectedKeys}
        openKeys={selectedMenu.OpenKeys}
        items={items}
        onClick={selectMenu}
        onOpenChange={openMenu}
        collapsed={collapsed.toString()}
      />
    </Sider>
  )
}

export default CommonSider
