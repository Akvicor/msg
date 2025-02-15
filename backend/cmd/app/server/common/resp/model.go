package resp

type Code int

const (
	Succeeded       Code = 0 // 成功
	Failed          Code = 1 // 失败
	NotFound        Code = 2 // 未找到
	BadRequest      Code = 3 // 错误输入
	UnAuthorized    Code = 4 // 未登录
	Forbidden       Code = 5 // 未授权
	Conflict        Code = 6 // 数据冲突
	TooManyRequests Code = 7 // 太多请求
)

type Model struct {
	Code Code       `json:"code"`
	Page *PageModel `json:"page,omitempty"`
	Msg  string     `json:"msg,omitempty"`
	Data any        `json:"data,omitempty"`
}

func NewModel(code Code, page *PageModel, msg string, data any) *Model {
	return &Model{
		Code: code,
		Page: page,
		Msg:  msg,
		Data: data,
	}
}
