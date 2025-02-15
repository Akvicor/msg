package dto

type UserBindHomeTipsSave struct {
	Content string `json:"content" form:"content" query:"content"`
}
