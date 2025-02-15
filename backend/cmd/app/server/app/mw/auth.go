package mw

import (
	"github.com/labstack/echo/v4"
	"msg/cmd/app/server/common/resp"
	"msg/cmd/app/server/common/token"
	"msg/cmd/app/server/common/types/role"
	"msg/cmd/app/server/global/auth"
)

// AuthAdmin 允许管理员访问，通过LoginToken或AccessToken鉴权
func AuthAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenStr := token.GetToken(c)
		authorization := auth.GetAuthorizationByToken(tokenStr)

		if authorization == nil {
			return resp.FailWithMsg(c, resp.UnAuthorized, "您的登录信息已失效，请重新登录后再试。")
		}

		if authorization.User.Disabled != 0 {
			return resp.FailWithMsg(c, resp.Forbidden, "您的账户已停用")
		}

		if authorization.User.Role != role.TypeAdmin {
			return resp.FailWithMsg(c, resp.Forbidden, "您没有权限")
		}

		// 更新登录有效期
		authorization.UpdateExpiration()

		return next(c)
	}
}

// AuthUser 允许管理员、普通用户访问，通过LoginToken或AccessToken鉴权
func AuthUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenStr := token.GetToken(c)
		authorization := auth.GetAuthorizationByToken(tokenStr)

		if authorization == nil {
			return resp.FailWithMsg(c, resp.UnAuthorized, "您的登录信息已失效，请重新登录后再试。")
		}

		if authorization.User.Disabled != 0 {
			return resp.FailWithMsg(c, resp.Forbidden, "您的账户已停用")
		}

		if authorization.User.Role != role.TypeAdmin && authorization.User.Role != role.TypeUser {
			return resp.FailWithMsg(c, resp.Forbidden, "您没有权限")
		}

		// 更新登录有效期
		authorization.UpdateExpiration()

		return next(c)
	}
}

// AuthViewer 允许管理员、普通用户、浏览者访问，通过LoginToken或AccessToken鉴权
func AuthViewer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenStr := token.GetToken(c)
		authorization := auth.GetAuthorizationByToken(tokenStr)

		if authorization == nil {
			return resp.FailWithMsg(c, resp.UnAuthorized, "您的登录信息已失效，请重新登录后再试。")
		}

		if authorization.User.Disabled != 0 {
			return resp.FailWithMsg(c, resp.Forbidden, "您的账户已停用")
		}

		if authorization.User.Role != role.TypeAdmin && authorization.User.Role != role.TypeUser && authorization.User.Role != role.TypeViewer {
			return resp.FailWithMsg(c, resp.Forbidden, "您没有权限")
		}

		// 更新登录有效期
		authorization.UpdateExpiration()

		return next(c)
	}
}

// Auth 通过LoginToken或AccessToken鉴权
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenStr := token.GetToken(c)
		authorization := auth.GetAuthorizationByToken(tokenStr)

		if authorization == nil {
			return resp.FailWithMsg(c, resp.UnAuthorized, "您的登录信息已失效，请重新登录后再试。")
		}

		if authorization.User.Disabled != 0 {
			return resp.FailWithMsg(c, resp.Forbidden, "您的账户已停用")
		}

		// 更新登录有效期
		authorization.UpdateExpiration()

		return next(c)
	}
}
