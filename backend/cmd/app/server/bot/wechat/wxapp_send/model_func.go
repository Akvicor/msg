package wxapp_send

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Akvicor/glog"
	"github.com/patrickmn/go-cache"
	"io"
	"msg/cmd/config"
	"net/http"
	"time"
)

func NewMessageMarkdown(agentId int64, toUser, msgType, content string) *MessageTypeMarkdownModel {
	return &MessageTypeMarkdownModel{
		MessageTypeBaseModel: MessageTypeBaseModel{
			Touser:  toUser,
			Msgtype: msgType,
			Agentid: fmt.Sprint(agentId),
		},
		Markdown: MessageTypeBaseMarkdownModel{
			Content: content,
		},
	}
}

func NewMessageTextCard(agentId int64, toUser, msgType, title, content, url, btn string) *MessageTypeTextCardModel {
	return &MessageTypeTextCardModel{
		MessageTypeBaseModel: MessageTypeBaseModel{
			Touser:  toUser,
			Msgtype: msgType,
			Agentid: fmt.Sprint(agentId),
		},
		TextCard: MessageTypeBaseTextCardModel{
			Title:       title,
			Description: content,
			URL:         url,
			BtnTxt:      btn,
		},
	}
}

func NewMessageText(agentId int64, toUser, msgType, content string) *MessageTypeTextModel {
	return &MessageTypeTextModel{
		MessageTypeBaseModel: MessageTypeBaseModel{
			Touser:  toUser,
			Msgtype: msgType,
			Agentid: fmt.Sprint(agentId),
		},
		Text: MessageTypeBaseTextModel{
			Content: content,
		},
	}
}

func NewSendModel(corpID, corpSecret string, agentId int64) *SendModel {
	return &SendModel{
		cache:      cache.New(2*time.Hour-5*time.Second, 10*time.Minute),
		CorpId:     corpID,
		CorpSecret: corpSecret,
		AgentId:    agentId,
	}
}

func (s *SendModel) getToken(newToken bool) string {
	if !newToken {
		token, ok := s.cache.Get("token")
		if ok {
			tokenS, ok := token.(string)
			if ok {
				return tokenS
			}
		}
	}
	try := 0
	for try < 3 {
		url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", s.CorpId, s.CorpSecret)
		resp, err := http.Get(url)
		if err != nil {
			try++
			glog.Warning("get token failed [%d]", try)
			time.Sleep(time.Second)
			continue
		}
		tokenB, _ := io.ReadAll(resp.Body)
		var tokenObj GetTokenResultModel
		err = json.Unmarshal(tokenB, &tokenObj)
		if err != nil {
			try++
			glog.Warning("unmarshal token failed [%d] [%s]", try, string(tokenB))
			time.Sleep(time.Second)
			continue
		}
		if tokenObj.Errcode != 0 {
			try++
			glog.Warning("token error code [%d] [%d][%s]", try, tokenObj.Errcode, tokenObj.Errmsg)
			time.Sleep(time.Second)
			continue
		}
		s.cache.Set("token", tokenObj.AccessToken, time.Duration(tokenObj.ExpiresIn-3)*time.Second)
		return tokenObj.AccessToken
	}
	glog.Error("get token error [%d]", try)
	return ""
}

func (s *SendModel) Send(to string, messageType string, title, message, url, btn string) error {
	if to == "" {
		to = "@all"
	}
	if messageType == "" {
		messageType = "text"
	}
	if url == "" {
		url = config.Global.Server.BaseUrl
	}
	if btn == "" {
		btn = "Button"
	}
	var contentB []byte
	var err error
	if messageType == "markdown" {
		msg := NewMessageMarkdown(s.AgentId, to, messageType, message)
		contentB, err = json.Marshal(msg)
	} else if messageType == "textcard" {
		msg := NewMessageTextCard(s.AgentId, to, messageType, title, message, url, btn)
		contentB, err = json.Marshal(msg)
	} else if messageType == "text" {
		msg := NewMessageText(s.AgentId, to, messageType, message)
		contentB, err = json.Marshal(msg)
	}
	if err != nil {
		return err
	}
	try := 0
	for try < 3 {
		token := s.getToken(try != 0)
		url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token)
		_, err = s.postJson(url, contentB)
		if err != nil {
			try++
			time.Sleep(time.Second)
			continue
		}
		return nil
	}
	return fmt.Errorf("send failed [%d][%v]", try, err)
}

func (s *SendModel) postJson(url string, content []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(content))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var res PostJsonResultModel
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	if res.Errcode != 0 {
		return nil, fmt.Errorf("post json failed [%d][%s]", res.Errcode, res.Errmsg)
	}
	return body, nil
}
