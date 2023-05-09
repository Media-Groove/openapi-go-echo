package openapi

import (
	"context"

	"github.com/labstack/echo/v4"
)

type EchoContext struct {
	echo.Context

	ctx context.Context
}

// noinspection GoUnusedExportedFunction
func NewEchoContext(ec echo.Context) *EchoContext {
	c := EchoContext{Context: ec}
	c.ctx = ec.Request().Context()
	return &c
}

func (ec *EchoContext) Ctx() context.Context {
	if ec.ctx == nil {
		ec.ctx = ec.Request().Context()
	}
	return ec.ctx
}

func (ec *EchoContext) SetCtx(ctx context.Context) {
	ec.ctx = ctx
}

// EchoContextMW EchoContextをセットするミドルウェアを返す関数
// EchoContextをセットするミドルウェア
func EchoContextMW(next echo.HandlerFunc) echo.HandlerFunc {
	// ミドルウェアが実際に行う処理
	return func(c echo.Context) error {
		ec := NewEchoContext(c)
		return next(ec)
	}
}
