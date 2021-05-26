package openapi

import (
	"encoding/json"
	"errors"
	"github.com/Media-Groove/go-utils/sliconv"
	"net/url"
	"strconv"
)

var ErrMsgRequiredMissing = errors.New("required parameter is missing")

func param_int32(val, defVal string) (int32, error) {
	val, err := param_string(val, defVal)
	if err != nil {
		return 0, err
	}
	n, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(n), nil
}

func param_int64(val, defVal string) (int64, error) {
	val, err := param_string(val, defVal)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(val, 10, 64)
}

func param_float32(val, defVal string) (float32, error) {
	val, err := param_string(val, defVal)
	if err != nil {
		return 0, err
	}
	n, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return 0, err
	}
	return float32(n), nil
}

func param_float64(val, defVal string) (float64, error) {
	val, err := param_string(val, defVal)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(val, 64)
}

func param_bool(val, defVal string) (bool, error) {
	val, err := param_string(val, defVal)
	if err != nil {
		return false, err
	}
	return strconv.ParseBool(val)
}

func param_string(val, defVal string) (string, error) {
	if val == "" {
		if defVal == "" {
			return "", ErrMsgRequiredMissing
		}
		val = defVal
	}
	return val, nil
}

func param_stringAry(params url.Values, name string, required bool, defVal string) ([]string, error) {
	vals, ok := params[name]
	if required && !ok {
		return nil, ErrMsgRequiredMissing // 必須なのにパラメータが存在しない場合
	}
	if len(vals) == 0 && defVal != "" {
		// デフォルト値取得
		err := json.Unmarshal([]byte(defVal), &vals)
		if err != nil {
			return nil, err // デフォルト値のパースエラー
		}
		if len(vals) == 0 {
			return vals, nil // デフォルト値をそのまま返す
		}
	}
	return vals, nil
}

func param_int32Ary(params url.Values, name string, required bool, defVal string) ([]int32, error) {
	vals, ok := params[name]
	if required && !ok {
		return nil, ErrMsgRequiredMissing // 必須なのにパラメータが存在しない場合
	}
	if len(vals) == 0 && defVal != "" {
		// デフォルト値取得
		var ns []int32
		err := json.Unmarshal([]byte(defVal), &ns)
		if err != nil {
			return nil, err // デフォルト値のパースエラー
		}
		return ns, nil // デフォルト値をそのまま返す
	}
	return sliconv.String{vals}.ToInt32()
}

func param_int64Ary(params url.Values, name string, required bool, defVal string) ([]int64, error) {
	vals, ok := params[name]
	if required && !ok {
		return nil, ErrMsgRequiredMissing // 必須なのにパラメータが存在しない場合
	}
	if len(vals) == 0 && defVal != "" {
		// デフォルト値取得
		var ns []int64
		err := json.Unmarshal([]byte(defVal), &ns)
		if err != nil {
			return nil, err // デフォルト値のパースエラー
		}
		return ns, nil // デフォルト値をそのまま返す
	}
	return sliconv.String{vals}.ToInt64()
}

func param_boolAry(params url.Values, name string, required bool, defVal string) ([]bool, error) {
	vals, ok := params[name]
	if required && !ok {
		return nil, ErrMsgRequiredMissing // 必須なのにパラメータが存在しない場合
	}
	if len(vals) == 0 && defVal != "" {
		// デフォルト値取得
		var ns []bool
		err := json.Unmarshal([]byte(defVal), &ns)
		if err != nil {
			return nil, err // デフォルト値のパースエラー
		}
		return ns, nil // デフォルト値をそのまま返す
	}
	return sliconv.String{vals}.ToBool()
}

func param_float32Ary(params url.Values, name string, required bool, defVal string) ([]float32, error) {
	vals, ok := params[name]
	if required && !ok {
		return nil, ErrMsgRequiredMissing // 必須なのにパラメータが存在しない場合
	}
	if len(vals) == 0 && defVal != "" {
		// デフォルト値取得
		var ns []float32
		err := json.Unmarshal([]byte(defVal), &ns)
		if err != nil {
			return nil, err // デフォルト値のパースエラー
		}
		return ns, nil // デフォルト値をそのまま返す
	}
	return sliconv.String{vals}.ToFloat32()
}

func param_float64Ary(params url.Values, name string, required bool, defVal string) ([]float64, error) {
	vals, ok := params[name]
	if required && !ok {
		return nil, ErrMsgRequiredMissing // 必須なのにパラメータが存在しない場合
	}
	if len(vals) == 0 && defVal != "" {
		// デフォルト値取得
		var ns []float64
		err := json.Unmarshal([]byte(defVal), &ns)
		if err != nil {
			return nil, err // デフォルト値のパースエラー
		}
		return ns, nil // デフォルト値をそのまま返す
	}
	return sliconv.String{vals}.ToFloat64()
}
