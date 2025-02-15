import {Navigate} from "react-router-dom";

export const RouterAuth = ({children}) => {
  const token = localStorage.getItem('login-token')
  if (!token) {
    return <Navigate to="/login" replace/>
  }
  return (
    children
  )
}

export const AdminAuth = ({children}) => {
  const role = localStorage.getItem('role')
  if (role !== 'admin') {
    return <Navigate to="/home" replace/>
  }
  return (
    children
  )
}

export const UserAuth = ({children}) => {
  const role = localStorage.getItem('role')
  if (role !== 'admin' && role !== 'user') {
    return <Navigate to="/home" replace/>
  }
  return (
    children
  )
}

export const ViewerAuth = ({children}) => {
  const role = localStorage.getItem('role')
  if (role !== 'admin' && role !== 'user' && role !== 'viewer') {
    return <Navigate to="/home" replace/>
  }
  return (
    children
  )
}
