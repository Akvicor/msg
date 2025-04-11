import {createBrowserRouter, Navigate} from 'react-router-dom'
import Main from '../pages/main'
import Home from '../pages/home'
import User from "../pages/user";
import Channel from "../pages/channel";
import Send from "../pages/send";
import Login from '../pages/login'
import {AdminAuth, RouterAuth, UserAuth, ViewerAuth} from './routerAuth'
import UserCenter from "../pages/userCenter";
import AccessToken from "../pages/accessToken";
import Schedule from "../pages/schedule";


const routes = [
  {
    path: '/',
    Component: Main,
    children: [
      {
        path: '/',
        element: (<RouterAuth><Navigate to="/home" replace/></RouterAuth>)
      },
      {
        path: 'home',
        Component: Home
      },
      {
        path: 'channel',
        element: (<ViewerAuth><Channel/></ViewerAuth>)
      },
      {
        path: 'schedule',
        element: (<ViewerAuth><Schedule/></ViewerAuth>)
      },
      {
        path: 'send',
        element: (<ViewerAuth><Send/></ViewerAuth>)
      },
      {
        path: 'admin',
        children: [
          {
            path: 'user-center',
            element: (<ViewerAuth><UserCenter/></ViewerAuth>)
          },
          {
            path: 'user',
            element: (<AdminAuth><User/></AdminAuth>)
          },
          {
            path: 'access_token',
            element: (<UserAuth><AccessToken/></UserAuth>)
          }
        ]
      }
    ]
  }, {
    path: '/login',
    Component: Login
  }
]

export default createBrowserRouter(routes)
