package openapi

import (
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
)

func HeaderParam_string(ec echo.Context, name, defVal string) (string, error) {
	return param_string(ec.Request().Header.Get(name), defVal)
}

func HeaderParam_int32(ec echo.Context, name, defVal string) (int32, error) {
	return param_int32(ec.Request().Header.Get(name), defVal)
}

func HeaderParam_int64(ec echo.Context, name, defVal string) (int64, error) {
	return param_int64(ec.Request().Header.Get(name), defVal)
}

func HeaderParam_bool(ec echo.Context, name, defVal string) (bool, error) {
	return param_bool(ec.Request().Header.Get(name), defVal)
}

func HeaderParam_float32(ec echo.Context, name, defVal string) (float32, error) {
	return param_float32(ec.Request().Header.Get(name), defVal)
}

func HeaderParam_float64(ec echo.Context, name, defVal string) (float64, error) {
	return param_float64(ec.Request().Header.Get(name), defVal)
}

func headerToUrlValues(h http.Header, name string) url.Values {
	urlValues := url.Values{}
	vals := h.Values(name)
	if len(vals) == 0 {
		return urlValues
	}
	urlValues[name] = regexp.MustCompile(`,\s*`).Split(strings.Join(vals, ","), -1)
	return urlValues
}

func HeaderParam_stringAry(ec echo.Context, name string, required bool, defVal string) ([]string, error) {
	vals := headerToUrlValues(ec.Request().Header, name)
	return param_stringAry(vals, name, required, defVal)
}

func HeaderParam_int32Ary(ec echo.Context, name string, required bool, defVal string) ([]int32, error) {
	vals := headerToUrlValues(ec.Request().Header, name)
	return param_int32Ary(vals, name, required, defVal)
}

func HeaderParam_int64Ary(ec echo.Context, name string, required bool, defVal string) ([]int64, error) {
	vals := headerToUrlValues(ec.Request().Header, name)
	return param_int64Ary(vals, name, required, defVal)
}

func HeaderParam_boolAry(ec echo.Context, name string, required bool, defVal string) ([]bool, error) {
	vals := headerToUrlValues(ec.Request().Header, name)
	return param_boolAry(vals, name, required, defVal)
}

func HeaderParam_float32Ary(ec echo.Context, name string, required bool, defVal string) ([]float32, error) {
	vals := headerToUrlValues(ec.Request().Header, name)
	return param_float32Ary(vals, name, required, defVal)
}

func HeaderParam_float64Ary(ec echo.Context, name string, required bool, defVal string) ([]float64, error) {
	vals := headerToUrlValues(ec.Request().Header, name)
	return param_float64Ary(vals, name, required, defVal)
}

func HeaderParam_stringPtr(ec echo.Context, name string) (*string, error) {
	s, err := HeaderParam_string(ec, name, "")
	switch err {
	case ErrMsgRequiredMissing:
		return nil, nil
	case nil:
		return &s, nil
	default:
		return nil, err
	}
}

func HeaderParam_int32Ptr(ec echo.Context, name string) (*int32, error) {
	n, err := HeaderParam_int32(ec, name, "")
	switch err {
	case ErrMsgRequiredMissing:
		return nil, nil
	case nil:
		return &n, nil
	default:
		return nil, err
	}
}

func HeaderParam_int64Ptr(ec echo.Context, name string) (*int64, error) {
	n, err := HeaderParam_int64(ec, name, "")
	switch err {
	case ErrMsgRequiredMissing:
		return nil, nil
	case nil:
		return &n, nil
	default:
		return nil, err
	}
}

func HeaderParam_boolPtr(ec echo.Context, name string) (*bool, error) {
	b, err := HeaderParam_bool(ec, name, "")
	switch err {
	case ErrMsgRequiredMissing:
		return nil, nil
	case nil:
		return &b, nil
	default:
		return nil, err
	}
}

func HeaderParam_float32Ptr(ec echo.Context, name string) (*float32, error) {
	n, err := HeaderParam_float32(ec, name, "")
	switch err {
	case ErrMsgRequiredMissing:
		return nil, nil
	case nil:
		return &n, nil
	default:
		return nil, err
	}
}

func HeaderParam_float64Ptr(ec echo.Context, name string) (*float64, error) {
	n, err := HeaderParam_float64(ec, name, "")
	switch err {
	case ErrMsgRequiredMissing:
		return nil, nil
	case nil:
		return &n, nil
	default:
		return nil, err
	}
}
