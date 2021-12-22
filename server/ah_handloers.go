package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// SetAhHandlers GAE の /_ah/* へのアクセスをエラーにしないためのハンドラ設定
func SetAhHandlers(e *echo.Echo) {
	// warmup
	e.GET("/_ah/warmup", handlerNoContentOK)
	// basic_scaling start/stop
	e.GET("/_ah/start", handlerNoContentOK)
	e.GET("/_ah/stop", handlerNoContentOK)
}

func handlerNoContentOK(ec echo.Context) error {
	return ec.NoContent(http.StatusOK)
}
