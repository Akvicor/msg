package app

import (
	"github.com/Akvicor/glog"
	"github.com/labstack/echo/v4"
	"io/fs"
	"msg/cmd/app/server/web"
	"msg/cmd/config"
	"net/http"
	"os"
)

func getFS() fs.FS {
	if config.Global.Debug {
		glog.Debug("using live mode")
		return os.DirFS(config.Global.Server.WebPath)
	}
	glog.Debug("using embed mode")
	fss, err := fs.Sub(web.Resource, config.Global.Server.WebPath)
	if err != nil {
		glog.Fatal("get embed web file failed: %v", err)
	}
	return fss
}

func WrapHandler(h http.Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Cache-Control", `public, max-age=31536000`)
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
