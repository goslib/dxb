package dxb

import (
	"context"
	"time"
)

func NewTimoutContext(second int64) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(second*int64(time.Second)))
	return ctx
}

func NewBackgroundContext() context.Context {
	return context.Background()
}

var DefaultContext = NewBackgroundContext
