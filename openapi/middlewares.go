package openapi

import (
	"github.com/labstack/echo/v4"
)

type Middlewares interface {
	Get(route Route) []echo.MiddlewareFunc
}

type NoMiddlewares struct {
	Middlewares
}

// noinspection GoUnusedParameter
func (m *NoMiddlewares) Get(route Route) []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{}
}
