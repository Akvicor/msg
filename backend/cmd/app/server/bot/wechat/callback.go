package wechat

import (
	"encoding/xml"
	"github.com/Akvicor/glog"
	"github.com/labstack/echo/v4"
	"io"
	"msg/cmd/app/server/bot/status"
	"msg/cmd/app/server/bot/wechat/crypt"
	"msg/cmd/app/server/bot/wechat/dro"
	"msg/cmd/app/server/bot/wechat/dto"
	"net/http"
)

type callbackFunc func(toUserName, fromUsername, createTime, msgType, content, msgId string, agentId int64) (rToUserName, rFromUsername string, rCreateTime int64, rMsgType, rContent string)

type CallbackModel struct {
	CorpID   string
	Token    string
	AesKey   string
	wx       *crypt.WXBizMsgCrypt
	callback callbackFunc
}

func NewCallbackModel(corpID, token, aesKey string, callback callbackFunc) *CallbackModel {
	return &CallbackModel{
		CorpID:   corpID,
		Token:    token,
		AesKey:   aesKey,
		wx:       crypt.NewWXBizMsgCrypt(token, aesKey, corpID, crypt.XmlType),
		callback: callback,
	}
}

func (s *CallbackModel) Status() status.SenderStatus {
	return status.SenderStatusOK
}

// Verify 验证
func (a *CallbackModel) Verify(c echo.Context) (err error) {
	input := new(dto.Verify)
	if err = c.Bind(input); err != nil {
		return c.String(http.StatusBadRequest, "错误输入: "+err.Error())
	}
	err = input.Validate()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	var echoStr []byte
	{
		data, e := a.wx.VerifyURL(input.MsgSignature, input.TimeStamp, input.Nonce, input.EchoStr)
		if e != nil {
			return c.String(http.StatusOK, "verify failed"+e.ErrMsg)
		}
		echoStr = data
	}

	return c.String(http.StatusOK, string(echoStr))
}

// Callback 信息接收
func (a *CallbackModel) Callback(c echo.Context) (err error) {
	input := new(dto.Callback)
	input.MsgSignature = c.QueryParam("msg_signature")
	input.TimeStamp = c.QueryParam("timestamp")
	input.Nonce = c.QueryParam("nonce")

	err = input.Validate()
	if err != nil {
		glog.Info("验证错误: " + err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	// parse body
	resB, err := io.ReadAll(c.Request().Body)
	if err != nil {
		glog.Info("读取body错误" + err.Error())
		return
	}
	res := dto.CallbackReceivedMessagePackageModel{}
	err = xml.Unmarshal(resB, &res)
	if err != nil {
		glog.Info("解析Body错误: " + err.Error())
		return c.String(http.StatusBadRequest, "解析Body错误: "+err.Error())
	}

	// verify
	var echoStr []byte
	{
		data, e := a.wx.VerifyURL(input.MsgSignature, input.TimeStamp, input.Nonce, res.Encrypt)
		if e != nil {
			glog.Info("verify failed" + e.ErrMsg)
			return c.String(http.StatusOK, "verify failed"+e.ErrMsg)
		}
		echoStr = data
	}

	req := dto.CallbackReceivedMessageModel{}
	err = xml.Unmarshal(echoStr, &req)
	if err != nil {
		glog.Info("unmarshal message failed" + err.Error())
		return c.String(http.StatusBadRequest, "unmarshal message failed"+err.Error())
	}

	rToUsername, rFromUsername, rCreateTime, rMsgType, rContent := a.callback(req.ToUserName, req.FromUserName, req.CreateTime, req.MsgType, req.Content, req.MsgId, req.AgentID)

	resp := dro.CallbackResponseMessageModel{
		ToUserName:   rToUsername,
		FromUserName: rFromUsername,
		CreateTime:   rCreateTime,
		MsgType:      rMsgType,
		Content:      rContent,
	}

	var respStr []byte
	{
		data, e := a.wx.EncryptMsg(resp.String(), input.TimeStamp, input.Nonce)
		if e != nil {
			glog.Info("encrypt failed" + e.ErrMsg)
			return c.String(http.StatusOK, "encrypt failed"+e.ErrMsg)
		}
		respStr = data
	}

	return c.String(http.StatusOK, string(respStr))
}
