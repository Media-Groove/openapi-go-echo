package main

import (
	"context"
	"github.com/Media-Groove/openapi-go-echo/openapi"
	"github.com/Media-Groove/openapi-go-echo/sampleapi"
	"github.com/Media-Groove/openapi-go-echo/server"
	"github.com/labstack/echo/v4"
	"time"
)

func main() {
	server.Start(&Initializer{}, 500*time.Millisecond)
}

type Initializer struct{}

func (i Initializer) Initialize(_ context.Context) {
}

func (i Initializer) EchoSetup(e *echo.Echo) {
	openapi.SetRoutes(e, &openapi.NoMiddlewares{}, sampleapi.NewSampleApiController(NewSampleApiService()))
}
