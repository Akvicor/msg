package mw

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
	"msg/cmd/app/server/common/ip_limiter"
	"msg/cmd/app/server/common/resp"
)

func NewIPLimiter(r rate.Limit, b int) func(next echo.HandlerFunc) echo.HandlerFunc {
	limiter := ip_limiter.NewIPRateLimiter(r, b)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ip := c.RealIP()
			if ip == "127.0.0.1" {
				return next(c)
			}
			l := limiter.GetLimiter(ip)

			if !l.Allow() {
				return resp.Fail(c, resp.TooManyRequests)
			}
			return next(c)
		}
	}
}
