/*
    * api
    *
* Generated by: OpenAPI Generator (https://openapi-generator.tech)
*/

package sampleapi

import (
	"net/http"
	"strings"

	"github.com/Media-Groove/openapi-go-echo/openapi"
	"github.com/labstack/echo/v4"
)

// A SampleApiController binds http requests to an api service and writes the service results to the http response
type SampleApiController struct {
	service SampleApiServicer
}

// NewSampleApiController creates a default api controller
// noinspection GoUnusedExportedFunction
func NewSampleApiController(s SampleApiServicer) openapi.Router {
	return &SampleApiController{service: s}
}

// Routes returns all of the api route for the SampleApiController
func (c *SampleApiController) Routes() openapi.Routes {
	return openapi.Routes{
		{
			"DeleteTestId",
			strings.ToUpper("Delete"),
			"/test/{id}",
			c.DeleteTestId,
			[]string{"sample"},
			[]string{"BasicAuth"},
		},
		{
			"GetTestId",
			strings.ToUpper("Get"),
			"/test/{id}",
			c.GetTestId,
			[]string{"sample"},
			[]string{},
		},
		{
			"PutTestId",
			strings.ToUpper("Put"),
			"/test/{id}",
			c.PutTestId,
			[]string{"sample"},
			[]string{"BasicAuth"},
		},
	}
}

// DeleteTestId - サンプルDELETE
func (c *SampleApiController) DeleteTestId(ec echo.Context) error {
	id, err := openapi.PathParam_int32(ec, "id", "")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	statusCode, json := c.service.DeleteTestId(ec.(*openapi.EchoContext), id)
	return openapi.Response(ec, statusCode, json)
}

// GetTestId - サンプルGET
func (c *SampleApiController) GetTestId(ec echo.Context) error {
	id, err := openapi.PathParam_int32(ec, "id", "")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	statusCode, json := c.service.GetTestId(ec.(*openapi.EchoContext), id)
	return openapi.Response(ec, statusCode, json)
}

// PutTestId - サンプルPUT
func (c *SampleApiController) PutTestId(ec echo.Context) error {
	id, err := openapi.PathParam_int32(ec, "id", "")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	var data Data
	if err := ec.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	statusCode, json := c.service.PutTestId(ec.(*openapi.EchoContext), id, &data)
	return openapi.Response(ec, statusCode, json)
}
