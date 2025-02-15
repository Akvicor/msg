package dro

import "encoding/xml"

type CallbackResponseMessageModel struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
}

func (r *CallbackResponseMessageModel) String() string {
	msgStr, err := xml.Marshal(r)
	if err != nil {
		return `{}`
	}
	return string(msgStr)
}
