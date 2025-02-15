package dto

import "encoding/xml"

type Verify struct {
	MsgSignature string `json:"msg_signature" form:"msg_signature" query:"msg_signature"`
	TimeStamp    string `json:"timestamp" form:"timestamp" query:"timestamp"`
	Nonce        string `json:"nonce" form:"nonce" query:"nonce"`
	EchoStr      string `json:"echostr" form:"echostr" query:"echostr"`
}

type Callback struct {
	MsgSignature string `json:"msg_signature" form:"msg_signature" query:"msg_signature"`
	TimeStamp    string `json:"timestamp" form:"timestamp" query:"timestamp"`
	Nonce        string `json:"nonce" form:"nonce" query:"nonce"`
	EchoStr      string `json:"echostr" form:"echostr" query:"echostr"`
}

type CallbackReceivedMessagePackageModel struct {
	XMLName    xml.Name `xml:"xml"`
	ToUserName string   `xml:"ToUserName"`
	Encrypt    string   `xml:"Encrypt"`
	AgentID    string   `xml:"AgentID"`
}

type CallbackReceivedMessageModel struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   string `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MsgId        string `xml:"MsgId"`
	AgentID      int64  `xml:"AgentID"`
}
