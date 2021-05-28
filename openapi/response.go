package openapi

import (
	"github.com/labstack/echo/v4"
)

func Response(ec echo.Context, statusCode int, resp interface{}) error {
	if err, ok := resp.(error); ok {
		ec.Logger().Error(err.Error())
		if _, ok = err.(*Error); !ok {
			resp = Error{
				Message: err.Error(),
			}
		}
	}
	return ec.JSON(statusCode, resp)
}
