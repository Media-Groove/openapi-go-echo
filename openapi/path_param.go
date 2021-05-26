package openapi

import (
	"strings"

	"github.com/Media-Groove/go-utils/sliconv"
	"github.com/labstack/echo/v4"
)

func PathParam_int32(ec echo.Context, name, defVal string) (int32, error) {
	return param_int32(ec.Param(name), defVal)
}

func PathParam_int64(ec echo.Context, name, defVal string) (int64, error) {
	return param_int64(ec.Param(name), defVal)
}

func PathParam_bool(ec echo.Context, name, defVal string) (bool, error) {
	return param_bool(ec.Param(name), defVal)
}

func PathParam_string(ec echo.Context, name, defVal string) (string, error) {
	return param_string(ec.Param(name), defVal)
}

func PathParam_float32(ec echo.Context, name, defVal string) (float32, error) {
	return param_float32(ec.Param(name), defVal)
}

func PathParam_float64(ec echo.Context, name, defVal string) (float64, error) {
	return param_float64(ec.Param(name), defVal)
}

func PathParam_int32Ary(ec echo.Context, name, defVal string) ([]int32, error) {
	ss, err := PathParam_stringAry(ec, name, defVal)
	if err != nil {
		return nil, err
	}
	return sliconv.String{ss}.ToInt32()
}

func PathParam_int64Ary(ec echo.Context, name, defVal string) ([]int64, error) {
	ss, err := PathParam_stringAry(ec, name, defVal)
	if err != nil {
		return nil, err
	}
	return sliconv.String{ss}.ToInt64()
}

func PathParam_boolAry(ec echo.Context, name, defVal string) ([]bool, error) {
	ss, err := PathParam_stringAry(ec, name, defVal)
	if err != nil {
		return nil, err
	}
	return sliconv.String{ss}.ToBool()
}

func PathParam_stringAry(ec echo.Context, name, defVal string) ([]string, error) {
	val, err := param_string(ec.Param(name), defVal)
	if err != nil {
		return nil, ErrMsgRequiredMissing
	}
	return strings.Split(val, ","), nil
}

func PathParam_float32Ary(ec echo.Context, name, defVal string) ([]float32, error) {
	ss, err := PathParam_stringAry(ec, name, defVal)
	if err != nil {
		return nil, err
	}
	return sliconv.String{ss}.ToFloat32()
}

func PathParam_float64Ary(ec echo.Context, name, defVal string) ([]float64, error) {
	ss, err := PathParam_stringAry(ec, name, defVal)
	if err != nil {
		return nil, err
	}
	return sliconv.String{ss}.ToFloat64()
}
