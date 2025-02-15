package dto

type SendModel struct {
	Touser string `json:"touser" form:"touser" query:"touser"`
	Type   string `json:"type" form:"type" query:"type"`
	Title  string `json:"title" form:"title" query:"title"`
	Msg    string `json:"msg" form:"msg" query:"msg"`
	Url    string `json:"url" form:"url" query:"url"`
	Btn    string `json:"btn" form:"btn" query:"btn"`
}
