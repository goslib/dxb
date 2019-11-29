package dxb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewTimoutContext(second int64) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(second*int64(time.Second)))
	return ctx
}

func NewBackgroundContext() context.Context {
	return context.Background()
}

var DefaultContext = NewBackgroundContext

func SortBy(option *options.FindOptions, key string, desc bool) *options.FindOptions {
	fmt.Println("TODO: Build the sort options:", key, desc)
	return option
}
