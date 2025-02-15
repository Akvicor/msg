package resp

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Healthy(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, data)
}

func Unhealthy(c echo.Context, data any) error {
	return c.JSON(http.StatusServiceUnavailable, data)
}

func Fail(c echo.Context, code Code) error {
	return c.JSON(http.StatusOK, NewModel(code, nil, "fail", nil))
}

func FailWithMsg(c echo.Context, code Code, msg string) error {
	return c.JSON(http.StatusOK, NewModel(code, nil, msg, nil))
}

func FailWithData(c echo.Context, code Code, data any) error {
	return c.JSON(http.StatusOK, NewModel(code, nil, "fail", data))
}

func FailWithPageData(c echo.Context, code Code, page *PageModel, data any) error {
	return c.JSON(http.StatusOK, NewModel(code, page, "fail", data))
}

func FailWithFull(c echo.Context, code Code, page *PageModel, msg string, data any) error {
	return c.JSON(http.StatusOK, NewModel(code, page, msg, data))
}

func Success(c echo.Context) error {
	return c.JSON(http.StatusOK, NewModel(Succeeded, nil, "ok", nil))
}

func SuccessWithMsg(c echo.Context, msg string) error {
	return c.JSON(http.StatusOK, NewModel(Succeeded, nil, msg, nil))
}

func SuccessWithData(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, NewModel(Succeeded, nil, "ok", data))
}

func SuccessWithPageData(c echo.Context, page *PageModel, data any) error {
	return c.JSON(http.StatusOK, NewModel(Succeeded, page, "ok", data))
}

func SuccessWithFull(c echo.Context, page *PageModel, msg string, data any) error {
	return c.JSON(http.StatusOK, NewModel(Succeeded, page, msg, data))
}
