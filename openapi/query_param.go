package openapi

import (
	"github.com/labstack/echo/v4"
)

func QueryParam_string(ec echo.Context, name, defVal string) (string, error) {
	return param_string(ec.QueryParam(name), defVal)
}

func QueryParam_int32(ec echo.Context, name, defVal string) (int32, error) {
	return param_int32(ec.QueryParam(name), defVal)
}

func QueryParam_int64(ec echo.Context, name, defVal string) (int64, error) {
	return param_int64(ec.QueryParam(name), defVal)
}

func QueryParam_bool(ec echo.Context, name, defVal string) (bool, error) {
	return param_bool(ec.QueryParam(name), defVal)
}

func QueryParam_float32(ec echo.Context, name, defVal string) (float32, error) {
	return param_float32(ec.QueryParam(name), defVal)
}

func QueryParam_float64(ec echo.Context, name, defVal string) (float64, error) {
	return param_float64(ec.QueryParam(name), defVal)
}

func QueryParam_stringAry(ec echo.Context, name string, required bool, defVal string) ([]string, error) {
	return param_stringAry(ec.QueryParams(), name, required, defVal)
}

func QueryParam_int32Ary(ec echo.Context, name string, required bool, defVal string) ([]int32, error) {
	return param_int32Ary(ec.QueryParams(), name, required, defVal)
}

func QueryParam_int64Ary(ec echo.Context, name string, required bool, defVal string) ([]int64, error) {
	return param_int64Ary(ec.QueryParams(), name, required, defVal)
}

func QueryParam_boolAry(ec echo.Context, name string, required bool, defVal string) ([]bool, error) {
	return param_boolAry(ec.QueryParams(), name, required, defVal)
}

func QueryParam_float32Ary(ec echo.Context, name string, required bool, defVal string) ([]float32, error) {
	return param_float32Ary(ec.QueryParams(), name, required, defVal)
}

func QueryParam_float64Ary(ec echo.Context, name string, required bool, defVal string) ([]float64, error) {
	return param_float64Ary(ec.QueryParams(), name, required, defVal)
}

func QueryParam_stringPtr(ec echo.Context, name string) (*string, error) {
	s, err := QueryParam_string(ec, name, "")
	switch err {
	case ErrMsgRequiredMissing:
		return nil, nil
	case nil:
		return &s, nil
	default:
		return nil, err
	}
}

func QueryParam_int32Ptr(ec echo.Context, name string) (*int32, error) {
	n, err := QueryParam_int32(ec, name, "")
	switch err {
	case ErrMsgRequiredMissing:
		return nil, nil
	case nil:
		return &n, nil
	default:
		return nil, err
	}
}

func QueryParam_int64Ptr(ec echo.Context, name string) (*int64, error) {
	n, err := QueryParam_int64(ec, name, "")
	switch err {
	case ErrMsgRequiredMissing:
		return nil, nil
	case nil:
		return &n, nil
	default:
		return nil, err
	}
}

func QueryParam_boolPtr(ec echo.Context, name string) (*bool, error) {
	b, err := QueryParam_bool(ec, name, "")
	switch err {
	case ErrMsgRequiredMissing:
		return nil, nil
	case nil:
		return &b, nil
	default:
		return nil, err
	}
}

func QueryParam_float32Ptr(ec echo.Context, name string) (*float32, error) {
	n, err := QueryParam_float32(ec, name, "")
	switch err {
	case ErrMsgRequiredMissing:
		return nil, nil
	case nil:
		return &n, nil
	default:
		return nil, err
	}
}

func QueryParam_float64Ptr(ec echo.Context, name string) (*float64, error) {
	n, err := QueryParam_float64(ec, name, "")
	switch err {
	case ErrMsgRequiredMissing:
		return nil, nil
	case nil:
		return &n, nil
	default:
		return nil, err
	}
}
