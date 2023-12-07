package loaders

import (
	"context"
	"time"
	"wangpeilei/leetcode/concurrent"
)

type InfoCtx struct {
	// req info
	Id int

	// resp info
	Name    string
	Age     int
	NameAge string
}

type LoaderResultHandler func() error

type LoaderHandler func(context.Context, *InfoCtx, *concurrent.ConcurrentLoader) (LoaderResultHandler, error)

func EmptyResultHandler() error {
	return nil
}

type LevelHandler struct {
	Name        string
	TimeOut     time.Duration
	HandlerList []LoaderHandler
}

var Handlers = []LevelHandler{
	{
		Name:    "level1",
		TimeOut: time.Second,
		HandlerList: []LoaderHandler{
			NameLoaderHandler,
			AgeLoaderHandler,
		},
	},
	{
		Name:    "level2",
		TimeOut: time.Second,
		HandlerList: []LoaderHandler{
			NameAndAgeLoaderHandler,
		},
	},
}
