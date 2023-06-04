/*
*
サンプルアプリケーション
*/
package main

import (
	"context"
	"github.com/Media-Groove/openapi-go-echo/openapi"
	"github.com/Media-Groove/openapi-go-echo/sampleapi"
	"github.com/Media-Groove/openapi-go-echo/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

func main() {
	server.Start(&Initializer{}, 500*time.Millisecond)
}

type Initializer struct{}

func (i Initializer) Initialize(_ context.Context) {
}

func (i Initializer) EchoSetup(_ context.Context, e *echo.Echo) {
	e.Use(middleware.CORS()) // swaggerからのアクセスを許可
	openapi.SetRoutes(e, &openapi.NoMiddlewares{}, sampleapi.NewSampleApiController(NewSampleApiService()))
}

func (i Initializer) Shutdown(_ context.Context) {
}
