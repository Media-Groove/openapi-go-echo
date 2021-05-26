package openapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormParam_string(t *testing.T) {
	e := echo.New()
	e.POST("/", func(ec echo.Context) error {
		s, err := FormParam_string(ec, "a", "")
		require.Nil(t, err)
		assert.Equal(t, "AAA", s)

		s, err = FormParam_string(ec, "b", "")
		require.Equal(t, ErrMsgRequiredMissing, err)

		s, err = FormParam_string(ec, "b", "BBB")
		require.Nil(t, err)
		assert.Equal(t, "BBB", s)

		return nil
	})

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("a=AAA"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	e.ServeHTTP(nil, req)
}

func TestFormParam_int32(t *testing.T) {
	e := echo.New()
	e.POST("/", func(ec echo.Context) error {
		n, err := FormParam_int32(ec, "a", "")
		require.Nil(t, err)
		assert.Equal(t, int32(123), n)

		n, err = FormParam_int32(ec, "b", "")
		require.Equal(t, ErrMsgRequiredMissing, err)

		n, err = FormParam_int32(ec, "b", "4")
		require.Nil(t, err)
		assert.Equal(t, int32(4), n)

		n, err = FormParam_int32(ec, "c", "")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("a=123&c=3.4"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	e.ServeHTTP(nil, req)
}

func TestFormParam_int64(t *testing.T) {
	e := echo.New()
	e.POST("/", func(ec echo.Context) error {
		n, err := FormParam_int64(ec, "a", "")
		require.Nil(t, err)
		assert.Equal(t, int64(123), n)

		n, err = FormParam_int64(ec, "b", "")
		require.Equal(t, ErrMsgRequiredMissing, err)

		n, err = FormParam_int64(ec, "b", "4")
		require.Nil(t, err)
		assert.Equal(t, int64(4), n)

		n, err = FormParam_int64(ec, "c", "")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("a=123&c=3.4"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	e.ServeHTTP(nil, req)
}

func TestFormParam_bool(t *testing.T) {
	e := echo.New()
	e.POST("/", func(ec echo.Context) error {
		b, err := FormParam_bool(ec, "a", "")
		require.Nil(t, err)
		assert.True(t, b)

		b, err = FormParam_bool(ec, "b", "")
		require.Equal(t, ErrMsgRequiredMissing, err)

		b, err = FormParam_bool(ec, "b", "TRUE")
		require.Nil(t, err)
		assert.True(t, b)

		b, err = FormParam_bool(ec, "c", "")
		require.Nil(t, err)
		assert.False(t, b)

		b, err = FormParam_bool(ec, "d", "")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("a=true&c=false&d=ERROR"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	e.ServeHTTP(nil, req)
}

func TestFormParam_float32(t *testing.T) {
	e := echo.New()
	e.POST("/", func(ec echo.Context) error {
		n, err := FormParam_float32(ec, "a", "")
		require.Nil(t, err)
		assert.Equal(t, float32(12.3), n)

		n, err = FormParam_float32(ec, "b", "")
		require.Equal(t, ErrMsgRequiredMissing, err)

		n, err = FormParam_float32(ec, "b", "4")
		require.Nil(t, err)
		assert.Equal(t, float32(4), n)

		n, err = FormParam_float32(ec, "c", "")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("a=12.3&c=ERROR"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	e.ServeHTTP(nil, req)
}

func TestFormParam_float64(t *testing.T) {
	e := echo.New()
	e.POST("/", func(ec echo.Context) error {
		n, err := FormParam_float64(ec, "a", "")
		require.Nil(t, err)
		assert.Equal(t, 12.3, n)

		n, err = FormParam_float64(ec, "b", "")
		require.Equal(t, ErrMsgRequiredMissing, err)

		n, err = FormParam_float64(ec, "b", "4")
		require.Nil(t, err)
		assert.Equal(t, float64(4), n)

		n, err = FormParam_float64(ec, "c", "")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("a=12.3&c=ERROR"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	e.ServeHTTP(nil, req)
}

func TestFormParam_stringAry(t *testing.T) {
	e := echo.New()
	e.POST("/", func(ec echo.Context) error {
		ss, err := FormParam_stringAry(ec, "a", false, "")
		require.Nil(t, err)
		assert.Equal(t, []string{"A", "AA", "AAA"}, ss)
		// 必須では無いので nil を返す
		ss, err = FormParam_stringAry(ec, "b", false, "")
		require.Nil(t, err)
		assert.Nil(t, ss)
		// 必須なのでErrMsgRequiredMissing
		ss, err = FormParam_stringAry(ec, "b", true, "")
		require.Equal(t, ErrMsgRequiredMissing, err)
		// 必須では無いのでデフォルト値
		ss, err = FormParam_stringAry(ec, "b", false, `["aaa", "bbb"]`)
		require.Nil(t, err)
		assert.Equal(t, []string{"aaa", "bbb"}, ss)
		// 必須なのでErrMsgRequiredMissing
		ss, err = FormParam_stringAry(ec, "b", true, `["aaa", "bbb"]`)
		require.Equal(t, ErrMsgRequiredMissing, err)
		// 必須では無いが、デフォルト値に問題があるのでエラー
		ss, err = FormParam_stringAry(ec, "b", false, `ERROR`)
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("a=A&a=AA&a=AAA"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	e.ServeHTTP(nil, req)
}

func TestFormParam_int32Ary(t *testing.T) {
	e := echo.New()
	e.POST("/", func(ec echo.Context) error {
		ns, err := FormParam_int32Ary(ec, "a", false, "")
		require.Nil(t, err)
		assert.Equal(t, []int32{1, 23, 456}, ns)
		// 必須では無いので空スライスを返す
		ns, err = FormParam_int32Ary(ec, "b", false, "")
		require.Nil(t, err)
		assert.Len(t, ns, 0)
		// 必須なのでErrMsgRequiredMissing
		ns, err = FormParam_int32Ary(ec, "b", true, "")
		require.Equal(t, ErrMsgRequiredMissing, err)
		// 必須では無いのでデフォルト値
		ns, err = FormParam_int32Ary(ec, "b", false, `[1, 2]`)
		require.Nil(t, err)
		assert.Equal(t, []int32{1, 2}, ns)
		// 必須なのでErrMsgRequiredMissing
		ns, err = FormParam_int32Ary(ec, "b", true, `[1, 2]`)
		require.Equal(t, ErrMsgRequiredMissing, err)
		// 必須では無いが、デフォルト値に問題があるのでエラー
		ns, err = FormParam_int32Ary(ec, "b", false, `["ERROR"]`)
		require.NotNil(t, err)
		// 数値では無い
		ns, err = FormParam_int32Ary(ec, "c", false, "")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("a=1&a=23&a=456&c=error"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	e.ServeHTTP(nil, req)
}

func TestFormParam_int64Ary(t *testing.T) {
	e := echo.New()
	e.POST("/", func(ec echo.Context) error {
		ns, err := FormParam_int64Ary(ec, "a", false, "")
		require.Nil(t, err)
		assert.Equal(t, []int64{1, 23, 456}, ns)
		// 必須では無いので空スライスを返す
		ns, err = FormParam_int64Ary(ec, "b", false, "")
		require.Nil(t, err)
		assert.Len(t, ns, 0)
		// 必須なのでErrMsgRequiredMissing
		ns, err = FormParam_int64Ary(ec, "b", true, "")
		require.Equal(t, ErrMsgRequiredMissing, err)
		// 必須では無いのでデフォルト値
		ns, err = FormParam_int64Ary(ec, "b", false, `[1, 2]`)
		require.Nil(t, err)
		assert.Equal(t, []int64{1, 2}, ns)
		// 必須なのでErrMsgRequiredMissing
		ns, err = FormParam_int64Ary(ec, "b", true, `[1, 2]`)
		require.Equal(t, ErrMsgRequiredMissing, err)
		// 必須では無いが、デフォルト値に問題があるのでエラー
		ns, err = FormParam_int64Ary(ec, "b", false, `["ERROR"]`)
		require.NotNil(t, err)
		// 数値では無い
		ns, err = FormParam_int64Ary(ec, "c", false, "")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("a=1&a=23&a=456&c=error"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	e.ServeHTTP(nil, req)
}

func TestFormParam_boolAry(t *testing.T) {
	e := echo.New()
	e.POST("/", func(ec echo.Context) error {
		ns, err := FormParam_boolAry(ec, "a", false, "")
		require.Nil(t, err)
		assert.Equal(t, []bool{true, false, true}, ns)
		// 必須では無いので空スライスを返す
		ns, err = FormParam_boolAry(ec, "b", false, "")
		require.Nil(t, err)
		assert.Len(t, ns, 0)
		// 必須なのでErrMsgRequiredMissing
		ns, err = FormParam_boolAry(ec, "b", true, "")
		require.Equal(t, ErrMsgRequiredMissing, err)
		// 必須では無いのでデフォルト値
		ns, err = FormParam_boolAry(ec, "b", false, `[true, false]`)
		require.Nil(t, err)
		assert.Equal(t, []bool{true, false}, ns)
		// 必須なのでErrMsgRequiredMissing
		ns, err = FormParam_boolAry(ec, "b", true, `[true, false]`)
		require.Equal(t, ErrMsgRequiredMissing, err)
		// 必須では無いが、デフォルト値に問題があるのでエラー
		ns, err = FormParam_boolAry(ec, "b", false, `["ERROR"]`)
		require.NotNil(t, err)
		// 数値では無い
		ns, err = FormParam_boolAry(ec, "c", false, "")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("a=1&a=false&a=TRUE&c=error"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	e.ServeHTTP(nil, req)
}

func TestFormParam_float32Ary(t *testing.T) {
	e := echo.New()
	e.POST("/", func(ec echo.Context) error {
		ns, err := FormParam_float32Ary(ec, "a", false, "")
		require.Nil(t, err)
		assert.Equal(t, []float32{1, 2.3, 4.56}, ns)
		// 必須では無いので空スライスを返す
		ns, err = FormParam_float32Ary(ec, "b", false, "")
		require.Nil(t, err)
		assert.Len(t, ns, 0)
		// 必須なのでErrMsgRequiredMissing
		ns, err = FormParam_float32Ary(ec, "b", true, "")
		require.Equal(t, ErrMsgRequiredMissing, err)
		// 必須では無いのでデフォルト値
		ns, err = FormParam_float32Ary(ec, "b", false, `[1.2, 3]`)
		require.Nil(t, err)
		assert.Equal(t, []float32{1.2, 3}, ns)
		// 必須なのでErrMsgRequiredMissing
		ns, err = FormParam_float32Ary(ec, "b", true, `[1.2, 3]`)
		require.Equal(t, ErrMsgRequiredMissing, err)
		// 必須では無いが、デフォルト値に問題があるのでエラー
		ns, err = FormParam_float32Ary(ec, "b", false, `["ERROR"]`)
		require.NotNil(t, err)
		// 数値では無い
		ns, err = FormParam_float32Ary(ec, "c", false, "")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("a=1&a=2.3&a=4.56&c=error"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	e.ServeHTTP(nil, req)
}

func TestFormParam_float64Ary(t *testing.T) {
	e := echo.New()
	e.POST("/", func(ec echo.Context) error {
		ns, err := FormParam_float64Ary(ec, "a", false, "")
		require.Nil(t, err)
		assert.Equal(t, []float64{1, 2.3, 4.56}, ns)
		// 必須では無いので空スライスを返す
		ns, err = FormParam_float64Ary(ec, "b", false, "")
		require.Nil(t, err)
		assert.Len(t, ns, 0)
		// 必須なのでErrMsgRequiredMissing
		ns, err = FormParam_float64Ary(ec, "b", true, "")
		require.Equal(t, ErrMsgRequiredMissing, err)
		// 必須では無いのでデフォルト値
		ns, err = FormParam_float64Ary(ec, "b", false, `[1.2, 3]`)
		require.Nil(t, err)
		assert.Equal(t, []float64{1.2, 3}, ns)
		// 必須なのでErrMsgRequiredMissing
		ns, err = FormParam_float64Ary(ec, "b", true, `[1.2, 3]`)
		require.Equal(t, ErrMsgRequiredMissing, err)
		// 必須では無いが、デフォルト値に問題があるのでエラー
		ns, err = FormParam_float64Ary(ec, "b", false, `["ERROR"]`)
		require.NotNil(t, err)
		// 数値では無い
		ns, err = FormParam_float64Ary(ec, "c", false, "")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("a=1&a=2.3&a=4.56&c=error"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	e.ServeHTTP(nil, req)
}

func TestFormParam_stringPtr(t *testing.T) {
	e := echo.New()
	e.POST("/", func(ec echo.Context) error {
		s, err := FormParam_stringPtr(ec, "a")
		require.Nil(t, err)
		assert.Equal(t, "AAA", *s)

		s, err = FormParam_stringPtr(ec, "b")
		require.Nil(t, err)
		assert.Nil(t, s)

		return nil
	})

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("a=AAA"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	e.ServeHTTP(nil, req)
}

func TestFormParam_int32Ptr(t *testing.T) {
	e := echo.New()
	e.POST("/", func(ec echo.Context) error {
		n, err := FormParam_int32Ptr(ec, "a")
		require.Nil(t, err)
		assert.Equal(t, int32(123), *n)

		n, err = FormParam_int32Ptr(ec, "b")
		require.Nil(t, err)
		assert.Nil(t, n)

		n, err = FormParam_int32Ptr(ec, "c")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("a=123&c=3.4"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	e.ServeHTTP(nil, req)
}

func TestFormParam_int64Ptr(t *testing.T) {
	e := echo.New()
	e.POST("/", func(ec echo.Context) error {
		n, err := FormParam_int64Ptr(ec, "a")
		require.Nil(t, err)
		assert.Equal(t, int64(123), *n)

		n, err = FormParam_int64Ptr(ec, "b")
		require.Nil(t, err)
		assert.Nil(t, n)

		n, err = FormParam_int64Ptr(ec, "c")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("a=123&c=3.4"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	e.ServeHTTP(nil, req)
}

func TestFormParam_boolPtr(t *testing.T) {
	e := echo.New()
	e.POST("/", func(ec echo.Context) error {
		b, err := FormParam_boolPtr(ec, "a")
		require.Nil(t, err)
		assert.True(t, *b)

		b, err = FormParam_boolPtr(ec, "b")
		require.Nil(t, err)
		assert.Nil(t, b)

		b, err = FormParam_boolPtr(ec, "c")
		require.Nil(t, err)
		assert.False(t, *b)

		b, err = FormParam_boolPtr(ec, "d")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("a=true&c=false&d=ERROR"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	e.ServeHTTP(nil, req)
}

func TestFormParam_float32Ptr(t *testing.T) {
	e := echo.New()
	e.POST("/", func(ec echo.Context) error {
		n, err := FormParam_float32Ptr(ec, "a")
		require.Nil(t, err)
		assert.Equal(t, float32(12.3), *n)

		n, err = FormParam_float32Ptr(ec, "b")
		require.Nil(t, err)
		assert.Nil(t, n)

		n, err = FormParam_float32Ptr(ec, "c")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("a=12.3&c=ERROR"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	e.ServeHTTP(nil, req)
}

func TestFormParam_float64Ptr(t *testing.T) {
	e := echo.New()
	e.POST("/", func(ec echo.Context) error {
		n, err := FormParam_float64Ptr(ec, "a")
		require.Nil(t, err)
		assert.Equal(t, 12.3, *n)

		n, err = FormParam_float64Ptr(ec, "b")
		require.Nil(t, err)
		assert.Nil(t, n)

		n, err = FormParam_float64Ptr(ec, "c")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("a=12.3&c=ERROR"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	e.ServeHTTP(nil, req)
}
