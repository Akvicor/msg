package dto

type SendModel struct {
	Chat int64  `json:"chat" form:"chat" query:"chat"`
	Mode string `json:"mode" form:"mode" query:"mode"`
	Msg  string `json:"msg" form:"msg" query:"msg"`
}
