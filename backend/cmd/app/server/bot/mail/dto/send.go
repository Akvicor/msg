package dto

import "github.com/wneessen/go-mail"

type SendModel struct {
	To      string           `json:"to" form:"to" query:"to"`
	Subject string           `json:"subject" form:"subject" query:"subject"`
	Type    mail.ContentType `json:"type" form:"type" query:"type"`
	Msg     string           `json:"msg" form:"msg" query:"msg"`
}
