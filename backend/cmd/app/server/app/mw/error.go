package mw

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"msg/cmd/app/server/common/resp"
)

func Error(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			var he *echo.HTTPError
			if errors.As(err, &he) {
				message := fmt.Sprintf("%v", he.Message)
				return resp.FailWithMsg(c, resp.Failed, message)
			}
			return resp.FailWithMsg(c, resp.Failed, "服务器错误")
		}
		return nil
	}
}
