package server

import (
	"context"
	"github.com/DeNA/aelog/middleware"
	"github.com/Media-Groove/openapi-go-echo/openapi"
	"github.com/labstack/echo/v4"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start(initializer Initializer, cancelWaitingTime time.Duration) {
	ctx := context.Background()

	initializer.Initialize(ctx)

	e := echo.New()
	// バナー非表示
	e.HideBanner = true
	// 共通ミドルウェア
	e.Use(openapi.EchoContextMW, echo.WrapMiddleware(middleware.AELogger("Atto")))
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
