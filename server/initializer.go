package server

import (
	"context"
	"github.com/labstack/echo/v4"
)

type Initializer interface {
	Initialize(ctx context.Context)

	EchoSetup(e *echo.Echo)
}
