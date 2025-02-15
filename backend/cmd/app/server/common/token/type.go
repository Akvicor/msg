package token

type Type string

const (
	TypeLoginToken  Type = "login-token"  // 登录令牌
	TypeAccessToken Type = "access-token" // 授权令牌
)
