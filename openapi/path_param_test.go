package openapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPathParam_int32(t *testing.T) {
	e := echo.New()
	e.GET("/:p1/:p2/:p3/:p4/:p5", func(ec echo.Context) error {
		n, err := PathParam_int32(ec, "p1", "")
		require.Nil(t, err)
		assert.Equal(t, int32(123), n)

		n, err = PathParam_int32(ec, "p2", "")
		require.Nil(t, err)
		assert.Equal(t, int32(4), n)

		n, err = PathParam_int32(ec, "p3", "11")
		require.Nil(t, err)
		assert.Equal(t, int32(11), n)

		n, err = PathParam_int32(ec, "p3", "")
		require.Equal(t, ErrMsgRequiredMissing, err)

		n, err = PathParam_int32(ec, "p4", "")
		require.NotNil(t, err)

		n, err = PathParam_int32(ec, "p5", "")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodGet, "/123/4//ERROR/5.6", nil)
	e.ServeHTTP(nil, req)
}

func TestPathParam_int64(t *testing.T) {
	e := echo.New()
	e.GET("/:p1/:p2/:p3/:p4/:p5", func(ec echo.Context) error {
		n, err := PathParam_int64(ec, "p1", "")
		require.Nil(t, err)
		assert.Equal(t, int64(123), n)

		n, err = PathParam_int64(ec, "p2", "")
		require.Nil(t, err)
		assert.Equal(t, int64(4), n)

		n, err = PathParam_int64(ec, "p3", "11")
		require.Nil(t, err)
		assert.Equal(t, int64(11), n)

		n, err = PathParam_int64(ec, "p3", "")
		require.Equal(t, ErrMsgRequiredMissing, err)

		n, err = PathParam_int64(ec, "p4", "")
		require.NotNil(t, err)

		n, err = PathParam_int64(ec, "p5", "")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodGet, "/123/4//ERROR/5.6", nil)
	e.ServeHTTP(nil, req)
}

func TestPathParam_float32(t *testing.T) {
	e := echo.New()
	e.GET("/:p1/:p2/:p3/:p4", func(ec echo.Context) error {
		n, err := PathParam_float32(ec, "p1", "")
		require.Nil(t, err)
		assert.Equal(t, float32(12.3), n)

		n, err = PathParam_float32(ec, "p2", "")
		require.Nil(t, err)
		assert.Equal(t, float32(4), n)

		n, err = PathParam_float32(ec, "p3", "11.1")
		require.Nil(t, err)
		assert.Equal(t, float32(11.1), n)

		n, err = PathParam_float32(ec, "p3", "")
		require.Equal(t, ErrMsgRequiredMissing, err)

		n, err = PathParam_float32(ec, "p4", "")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodGet, "/12.3/4//ERROR", nil)
	e.ServeHTTP(nil, req)
}

func TestPathParam_float64(t *testing.T) {
	e := echo.New()
	e.GET("/:p1/:p2/:p3/:p4", func(ec echo.Context) error {
		n, err := PathParam_float64(ec, "p1", "")
		require.Nil(t, err)
		assert.Equal(t, 12.3, n)

		n, err = PathParam_float64(ec, "p2", "")
		require.Nil(t, err)
		assert.Equal(t, float64(4), n)

		n, err = PathParam_float64(ec, "p3", "11.1")
		require.Nil(t, err)
		assert.Equal(t, 11.1, n)

		n, err = PathParam_float64(ec, "p3", "")
		require.Equal(t, ErrMsgRequiredMissing, err)

		n, err = PathParam_float64(ec, "p4", "")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodGet, "/12.3/4//ERROR", nil)
	e.ServeHTTP(nil, req)
}

func TestPathParam_bool(t *testing.T) {
	e := echo.New()
	e.GET("/:p1/:p2/:p3/:p4/:p5/:p6/:p7/:p8", func(ec echo.Context) error {
		b, err := PathParam_bool(ec, "p1", "")
		require.Nil(t, err)
		assert.True(t, b)

		b, err = PathParam_bool(ec, "p2", "")
		require.Nil(t, err)
		assert.True(t, b)

		b, err = PathParam_bool(ec, "p3", "")
		require.Nil(t, err)
		assert.True(t, b)

		b, err = PathParam_bool(ec, "p4", "")
		require.Nil(t, err)
		assert.True(t, b)

		b, err = PathParam_bool(ec, "p5", "true")
		require.Nil(t, err)
		assert.True(t, b)

		b, err = PathParam_bool(ec, "p5", "")
		require.Equal(t, ErrMsgRequiredMissing, err)

		b, err = PathParam_bool(ec, "p6", "")
		require.Nil(t, err)
		assert.False(t, b)

		b, err = PathParam_bool(ec, "p7", "")
		require.Nil(t, err)
		assert.False(t, b)

		b, err = PathParam_bool(ec, "p8", "")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodGet, "/true/True/TRUE/1//0/false/test", nil)
	e.ServeHTTP(nil, req)
}

func TestPathParam_string(t *testing.T) {
	e := echo.New()
	e.GET("/:p1/:p2/:p3", func(ec echo.Context) error {
		s, err := PathParam_string(ec, "p1", "")
		require.Nil(t, err)
		assert.Equal(t, "test1", s)

		s, err = PathParam_string(ec, "p2", "test0")
		require.Nil(t, err)
		assert.Equal(t, "test0", s)

		s, err = PathParam_string(ec, "p2", "")
		require.Equal(t, ErrMsgRequiredMissing, err)

		s, err = PathParam_string(ec, "p3", "")
		require.Nil(t, err)
		assert.Equal(t, "test2", s)

		return nil
	})

	req := httptest.NewRequest(http.MethodGet, "/test1//test2", nil)
	e.ServeHTTP(nil, req)
}

func TestPathParam_int32Ary(t *testing.T) {
	e := echo.New()
	e.GET("/:p1/:p2/:p3/:p4", func(ec echo.Context) error {
		ns, err := PathParam_int32Ary(ec, "p1", "")
		require.Nil(t, err)
		assert.Equal(t, []int32{123, 4}, ns)

		ns, err = PathParam_int32Ary(ec, "p2", "")
		require.NotNil(t, err)

		ns, err = PathParam_int32Ary(ec, "p3", "1,2,3")
		require.Nil(t, err)
		assert.Equal(t, []int32{1, 2, 3}, ns)

		ns, err = PathParam_int32Ary(ec, "p3", "")
		require.Equal(t, ErrMsgRequiredMissing, err)

		ns, err = PathParam_int32Ary(ec, "p4", "")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodGet, "/123,4/56,ERROR//7.8", nil)
	e.ServeHTTP(nil, req)
}

func TestPathParam_int64Ary(t *testing.T) {
	e := echo.New()
	e.GET("/:p1/:p2/:p3/:p4", func(ec echo.Context) error {
		ns, err := PathParam_int64Ary(ec, "p1", "")
		require.Nil(t, err)
		assert.Equal(t, []int64{123, 4}, ns)

		ns, err = PathParam_int64Ary(ec, "p2", "")
		require.NotNil(t, err)

		ns, err = PathParam_int64Ary(ec, "p3", "1,2,3")
		require.Nil(t, err)
		assert.Equal(t, []int64{1, 2, 3}, ns)

		ns, err = PathParam_int64Ary(ec, "p3", "")
		require.Equal(t, ErrMsgRequiredMissing, err)

		ns, err = PathParam_int64Ary(ec, "p4", "")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodGet, "/123,4/56,ERROR//7.8", nil)
	e.ServeHTTP(nil, req)
}

func TestPathParam_boolAry(t *testing.T) {
	e := echo.New()
	e.GET("/:p1/:p2/:p3", func(ec echo.Context) error {
		bs, err := PathParam_boolAry(ec, "p1", "")
		require.Nil(t, err)
		assert.Equal(t, []bool{true, true, true, true, false, false, false}, bs)

		bs, err = PathParam_boolAry(ec, "p2", "true,false")
		require.Nil(t, err)
		assert.Equal(t, []bool{true, false}, bs)

		bs, err = PathParam_boolAry(ec, "p2", "")
		require.Equal(t, ErrMsgRequiredMissing, err)

		bs, err = PathParam_boolAry(ec, "p3", "")
		require.NotNil(t, err)

		return nil
	})

	req := httptest.NewRequest(http.MethodGet, "/true,True,TRUE,1,0,false,FALSE//test", nil)
	e.ServeHTTP(nil, req)
}

func TestPathParam_stringAry(t *testing.T) {
	e := echo.New()
	e.GET("/:p1/:p2/:p3", func(ec echo.Context) error {
		ss, err := PathParam_stringAry(ec, "p1", "")
		require.Nil(t, err)
		assert.Equal(t, []string{"test1", "test2"}, ss)

		ss, err = PathParam_stringAry(ec, "p2", "t1,t2,t3")
		require.Nil(t, err)
		assert.Equal(t, []string{"t1", "t2", "t3"}, ss)

		ss, err = PathParam_stringAry(ec, "p2", "")
		require.Equal(t, ErrMsgRequiredMissing, err)

		ss, err = PathParam_stringAry(ec, "p3", "")
		require.Nil(t, err)
		assert.Equal(t, []string{"", "test2"}, ss)
		return nil
	})

	req := httptest.NewRequest(http.MethodGet, "/test1,test2//,test2", nil)
	e.ServeHTTP(nil, req)
}

func TestPathParam_float32Ary(t *testing.T) {
	e := echo.New()
	e.GET("/:p1/:p2/:p3/:p4", func(ec echo.Context) error {
		ns, err := PathParam_float32Ary(ec, "p1", "")
		require.Nil(t, err)
		assert.Equal(t, []float32{12.3, 4}, ns)

		ns, err = PathParam_float32Ary(ec, "p2", "")
		require.NotNil(t, err)

		ns, err = PathParam_float32Ary(ec, "p3", "1.2,3")
		require.Nil(t, err)
		assert.Equal(t, []float32{1.2, 3}, ns)

		ns, err = PathParam_float32Ary(ec, "p3", "")
		require.Equal(t, ErrMsgRequiredMissing, err)

		ns, err = PathParam_float32Ary(ec, "p4", "")
		require.Nil(t, err)
		assert.Equal(t, []float32{7.8}, ns)

		return nil
	})

	req := httptest.NewRequest(http.MethodGet, "/12.3,4/56,ERROR//7.8", nil)
	e.ServeHTTP(nil, req)
}

func TestPathParam_float64Ary(t *testing.T) {
	e := echo.New()
	e.GET("/:p1/:p2/:p3/:p4", func(ec echo.Context) error {
		ns, err := PathParam_float64Ary(ec, "p1", "")
		require.Nil(t, err)
		assert.Equal(t, []float64{12.3, 4}, ns)

		ns, err = PathParam_float64Ary(ec, "p2", "")
		require.NotNil(t, err)

		ns, err = PathParam_float64Ary(ec, "p3", "1.2,3")
		require.Nil(t, err)
		assert.Equal(t, []float64{1.2, 3}, ns)

		ns, err = PathParam_float64Ary(ec, "p3", "")
		require.Equal(t, ErrMsgRequiredMissing, err)

		ns, err = PathParam_float64Ary(ec, "p4", "")
		require.Nil(t, err)
		assert.Equal(t, []float64{7.8}, ns)

		return nil
	})

	req := httptest.NewRequest(http.MethodGet, "/12.3,4/56,ERROR//7.8", nil)
	e.ServeHTTP(nil, req)
}
