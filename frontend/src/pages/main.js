import React from 'react';

import {Layout, theme} from 'antd';
import CommonSider from "../components/commonSider";
import CommonHeader from "../components/commonHeader";
import CommonFooter from "../components/commonFooter";
import {useSelector} from "react-redux";
import {Outlet} from "react-router-dom";
import {RouterAuth} from '../router/routerAuth'

const {Content} = Layout;

const Main = () => {
  const {
    token: {colorBgContainer, borderRadiusLG},
  } = theme.useToken();
  const collapsed = useSelector(state => state.tab.isCollapse)
  const mode = useSelector(state => state.mode)
  return (
    <RouterAuth>
      <Layout>
        <CommonHeader mode={mode}/>
        <Layout
          style={{
            paddingTop: 64,
            minHeight: '100vh',
          }}
        >
            <CommonSider collapsed={collapsed} mode={mode}/>
            <Layout style={{
              // width: 'fit-content'
            }}>
              <Content
                style={{
                  margin: '16px 16px',
                  padding: 24,
                  minHeight: '89vh',
                  background: colorBgContainer,
                  borderRadius: borderRadiusLG,
                }}
              >
                <Outlet/>
              </Content>
              <CommonFooter/>
            </Layout>

        </Layout>
      </Layout>
    </RouterAuth>
  );
};
export default Main;