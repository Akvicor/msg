import React, {useEffect, useState} from 'react'
import {Layout} from "antd";
import './index.css';
import {getVersion, getVersionFull} from "../../api/public";

const {Footer} = Layout;

const CommonFooter = () => {
  const {version} = require("../../../package.json");
  const [apiVersion, setApiVersionData] = useState(localStorage.getItem('api-version'))
  const [apiVersionFull, setApiVersionFullData] = useState(localStorage.getItem('api-version-full'))
  useEffect(() => {
    getVersion().then(({data}) => {
      if (data.code === 0) {
        localStorage.setItem('api-version', data.data)
        setApiVersionData(data.data)
      }
    })
    getVersionFull().then(({data}) => {
      if (data.code === 0) {
        localStorage.setItem('api-version-full', data.data)
        setApiVersionFullData(data.data)
      }
    })
  }, [])
  const clickVersion = (e) => {
    if (e.target.textContent === apiVersion) {
      e.target.textContent = apiVersionFull
    } else {
      e.target.textContent = apiVersion
    }
  }
  return (
    <Footer
      style={{
        textAlign: 'center',
      }}
    >
      MSG ©{new Date().getFullYear()} Created by Akvicor<br/>WEB: <span>{version}</span> / API: <span
      onClick={(e) => {
        clickVersion(e)
      }}>{apiVersion}</span>
    </Footer>
  )
}

export default CommonFooter
