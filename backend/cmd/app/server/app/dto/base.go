package dto

type Id struct {
	ID int64 `json:"id" form:"id" query:"id"`
}

type Validator interface {
	Validate() error
}
