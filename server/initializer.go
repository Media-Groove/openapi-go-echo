package server

import (
	"context"
	"github.com/labstack/echo/v4"
)

// Initializer server.Start() 関数に渡す初期化用インターフェース
// 最初に Initialize() が呼ばれ、 echo.Echo インスタンスの準備が出来てから EchoSetup() が呼ばれる。
type Initializer interface {
	// Initialize 初期化処理を記述する
	// echo.Echo インスタンスの生成前に呼ばれるため、最初に行っておきたい全体的な初期化処理を記述する
	Initialize(ctx context.Context)

	// EchoSetup echo.Echo インスタンスに対するセットアップ処理を記述する
	// 主にルートやミドルウェアの設定を行う。
	EchoSetup(ctx context.Context, e *echo.Echo)
}
