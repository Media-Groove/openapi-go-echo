package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Media-Groove/openapi-go-echo/openapi"
	"github.com/labstack/echo/v4"
)

var e *echo.Echo

// Start Echo HTTP サーバ起動関数
// 指定された initializer で初期化を行う。
// cancelWaitingTime はサーバ終了時のシャット断処理待ち時間。
func Start(initializer Initializer, cancelWaitingTime time.Duration) {
	ctx := context.Background()

	initializer.Initialize(ctx)

	e = echo.New()
	// バナー非表示
	e.HideBanner = true
	// 共通ミドルウェア
	e.Use(openapi.EchoContextMW)
	// /_ah/* ハンドラ
	SetAhHandlers(e)
	// Echoサーバのセットアップ
	initializer.EchoSetup(e)

	// サーバ起動
	go func() {
		if err := e.Start(":" + GetPort()); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// 終了シグナル検知
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)
	<-sigCh
	// サーバ終了
	c, cancel := context.WithTimeout(ctx, cancelWaitingTime)
	defer cancel()
	if err := e.Shutdown(c); err != nil {
		e.Logger.Fatal(err)
	}
}

// GetPort 待ち受けポート取得
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	return port
}

func Shutdown(ctx context.Context) error {
	return e.Shutdown(ctx)
}
