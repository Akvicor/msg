import {ConfigProvider} from "antd";
import React from "react";
import {TinyColor} from "@ctrl/tinycolor";


export const ColorButtonProvider = ({danger, color, children}) => {
  if (danger) {
    color = '#ff4d4f'
  }
  const getHoverColors = new TinyColor(color).lighten(5).toString()
  const getActiveColors = new TinyColor(color).darken(5).toString()
  return (
    <ConfigProvider
      theme={{
        components: {
          Button: {
            colorPrimary: color,
            colorPrimaryHover: getHoverColors,
            colorPrimaryActive: getActiveColors
          },
        },
      }}
    >
      {children}
    </ConfigProvider>
  )
}
