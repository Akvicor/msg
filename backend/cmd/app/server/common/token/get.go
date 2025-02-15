package token

import (
	"github.com/labstack/echo/v4"
)

// GetXToken 获取授权令牌
func GetXToken(c echo.Context) string {
	key := c.Request().Header.Get(SignToken)
	if len(key) > 0 {
		return key
	}
	key = c.QueryParam(SignToken)
	if len(key) > 0 {
		return key
	}
	key = c.Request().Header.Get(SignXToken)
	if len(key) > 0 {
		return key
	}
	return c.QueryParam(SignXToken)
}

// GetAccessToken 获取授权令牌
func GetAccessToken(c echo.Context) string {
	key := c.Request().Header.Get(SignAccessToken)
	if len(key) > 0 {
		return key
	}
	return c.QueryParam(SignAccessToken)
}

// GetLoginToken 获取登录令牌
func GetLoginToken(c echo.Context) string {
	token := c.Request().Header.Get(SignLoginToken)
	if len(token) > 0 {
		return token
	}
	return c.QueryParam(SignLoginToken)
}

// GetToken 获取令牌
func GetToken(c echo.Context) string {
	var token string
	token = GetLoginToken(c)
	if len(token) > 0 {
		return token
	}
	token = GetAccessToken(c)
	if len(token) > 0 {
		return token
	}
	token = GetXToken(c)
	if len(token) > 0 {
		return token
	}
	return ""
}
