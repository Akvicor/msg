package sms

import (
	"fmt"
	"github.com/Akvicor/util"
	"github.com/labstack/echo/v4"
	"msg/cmd/app/server/bot/sms/dto"
	"msg/cmd/app/server/bot/status"
	"msg/cmd/app/server/common/resp"
	"time"
)

type Model struct {
	Api   string `toml:"api"`
	Token string `toml:"token"`
}

func NewModel(api string, token string) *Model {
	return &Model{
		Api:   api,
		Token: token,
	}
}

func (m *Model) Status() status.SenderStatus {
	return status.SenderStatusOK
}

func (m *Model) Send(phone, msg string) error {
	try := 0
	var err error = nil
	for try < 3 {
		var res []byte
		res, err = util.HttpPost(m.Api, map[string]string{
			"key":     m.Token,
			"sender":  "MSG",
			"phone":   phone,
			"message": msg,
		}, util.HTTPContentTypeUrlencoded, nil)
		if err != nil {
			err = fmt.Errorf("短信发送失败, HTTP请求错误: %v", err)
			try++
			time.Sleep(time.Second)
			continue
		}
		jsonRes := util.NewJSONResult(res)
		if jsonRes == nil {
			err = fmt.Errorf("短信发送失败, 解析数据错误: %s", string(res))
			try++
			time.Sleep(time.Second)
			continue
		}

		result := jsonRes.Map()
		if fmt.Sprint(result["code"]) != "0" {
			err = fmt.Errorf("短信发送失败, 请求错误: %v", result["msg"])
			try++
			time.Sleep(time.Second)
			continue
		}

		return nil
	}

	return err
}

// APISend 发送消息
func (m *Model) APISend(c echo.Context) (err error) {
	input := new(dto.SendModel)
	if err = c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, fmt.Sprintf("错误输入: %v", err))
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, fmt.Sprintf("错误输入: %v", err))
	}
	err = m.Send(input.Phone, input.Msg)
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, fmt.Sprintf("%v", err))
	}
	return resp.Success(c)
}
