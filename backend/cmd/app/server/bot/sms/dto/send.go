package dto

type SendModel struct {
	Phone string `json:"phone" form:"phone" query:"phone"`
	Msg   string `json:"msg" form:"msg" query:"msg"`
}
