package openapi

import (
	"github.com/labstack/echo/v4"
)

func FormParam_string(ec echo.Context, name, defVal string) (string, error) {
	return param_string(ec.FormValue(name), defVal)
}

func FormParam_int32(ec echo.Context, name, defVal string) (int32, error) {
	return param_int32(ec.FormValue(name), defVal)
}

func FormParam_int64(ec echo.Context, name, defVal string) (int64, error) {
	return param_int64(ec.FormValue(name), defVal)
}

func FormParam_bool(ec echo.Context, name, defVal string) (bool, error) {
	return param_bool(ec.FormValue(name), defVal)
}

func FormParam_float32(ec echo.Context, name, defVal string) (float32, error) {
	return param_float32(ec.FormValue(name), defVal)
}

func FormParam_float64(ec echo.Context, name, defVal string) (float64, error) {
	return param_float64(ec.FormValue(name), defVal)
}

func FormParam_stringAry(ec echo.Context, name string, required bool, defVal string) ([]string, error) {
	vals, err := ec.FormParams()
	if err != nil {
		return nil, err
	}
	return param_stringAry(vals, name, required, defVal)
}

func FormParam_int32Ary(ec echo.Context, name string, required bool, defVal string) ([]int32, error) {
	vals, err := ec.FormParams()
	if err != nil {
		return nil, err
	}
	return param_int32Ary(vals, name, required, defVal)
}

func FormParam_int64Ary(ec echo.Context, name string, required bool, defVal string) ([]int64, error) {
	vals, err := ec.FormParams()
	if err != nil {
		return nil, err
	}
	return param_int64Ary(vals, name, required, defVal)
}

func FormParam_boolAry(ec echo.Context, name string, required bool, defVal string) ([]bool, error) {
	vals, err := ec.FormParams()
	if err != nil {
		return nil, err
	}
	return param_boolAry(vals, name, required, defVal)
}

func FormParam_float32Ary(ec echo.Context, name string, required bool, defVal string) ([]float32, error) {
	vals, err := ec.FormParams()
	if err != nil {
		return nil, err
	}
	return param_float32Ary(vals, name, required, defVal)
}

func FormParam_float64Ary(ec echo.Context, name string, required bool, defVal string) ([]float64, error) {
	vals, err := ec.FormParams()
	if err != nil {
		return nil, err
	}
	return param_float64Ary(vals, name, required, defVal)
}

func FormParam_stringPtr(ec echo.Context, name string) (*string, error) {
	s, err := FormParam_string(ec, name, "")
	switch err {
	case ErrMsgRequiredMissing:
		return nil, nil
	case nil:
		return &s, nil
	default:
		return nil, err
	}
}

func FormParam_int32Ptr(ec echo.Context, name string) (*int32, error) {
	n, err := FormParam_int32(ec, name, "")
	switch err {
	case ErrMsgRequiredMissing:
		return nil, nil
	case nil:
		return &n, nil
	default:
		return nil, err
	}
}

func FormParam_int64Ptr(ec echo.Context, name string) (*int64, error) {
	n, err := FormParam_int64(ec, name, "")
	switch err {
	case ErrMsgRequiredMissing:
		return nil, nil
	case nil:
		return &n, nil
	default:
		return nil, err
	}
}

func FormParam_boolPtr(ec echo.Context, name string) (*bool, error) {
	b, err := FormParam_bool(ec, name, "")
	switch err {
	case ErrMsgRequiredMissing:
		return nil, nil
	case nil:
		return &b, nil
	default:
		return nil, err
	}
}

func FormParam_float32Ptr(ec echo.Context, name string) (*float32, error) {
	n, err := FormParam_float32(ec, name, "")
	switch err {
	case ErrMsgRequiredMissing:
		return nil, nil
	case nil:
		return &n, nil
	default:
		return nil, err
	}
}

func FormParam_float64Ptr(ec echo.Context, name string) (*float64, error) {
	n, err := FormParam_float64(ec, name, "")
	switch err {
	case ErrMsgRequiredMissing:
		return nil, nil
	case nil:
		return &n, nil
	default:
		return nil, err
	}
}
