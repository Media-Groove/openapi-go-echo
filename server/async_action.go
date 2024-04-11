package server

import (
	"context"
)

func AsyncAction(f func(ctx context.Context)) {
	go f(ctx)
}
