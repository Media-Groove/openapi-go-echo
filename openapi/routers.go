package openapi

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// A Route defines the parameters for an api endpoint
type Route struct {
	Name            string
	Method          string
	Pattern         string
	HandlerFunc     func(c echo.Context) error
	Tags            []string
	SecuritySchemes []string
}

// Routes are a collection of defined api endpoints
type Routes []Route

// Router defines the required methods for retrieving api routes
type Router interface {
	Routes() Routes
}

// noinspection GoUnusedExportedFunction
func SetRoutes(e *echo.Echo, middlewares Middlewares, routers ...Router) {
	for _, api := range routers {
		for _, route := range api.Routes() {
			pattern := strings.ReplaceAll(strings.ReplaceAll(route.Pattern, "{", ":"), "}", "")
			switch route.Method {
			case http.MethodHead:
				e.HEAD(pattern, route.HandlerFunc, middlewares.Get(route)...)
			case http.MethodGet:
				e.GET(pattern, route.HandlerFunc, middlewares.Get(route)...)
			case http.MethodPost:
				e.POST(pattern, route.HandlerFunc, middlewares.Get(route)...)
			case http.MethodPut:
				e.PUT(pattern, route.HandlerFunc, middlewares.Get(route)...)
			case http.MethodDelete:
				e.DELETE(pattern, route.HandlerFunc, middlewares.Get(route)...)
			case http.MethodPatch:
				e.PATCH(pattern, route.HandlerFunc, middlewares.Get(route)...)
			}
		}
	}
}
