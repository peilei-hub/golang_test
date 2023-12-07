package test

import (
	"context"
	"fmt"
	"testing"
	"wangpeilei/leetcode/concurrent"
	"wangpeilei/leetcode/concurrent/test/loaders"
)

func Test_Loader(t *testing.T) {
	err := Load()
	fmt.Println(err)
}

func Load() error {
	ctx := context.Background()

	infoCtx := &loaders.InfoCtx{
		Id: 2,
	}

	defer func() {
		fmt.Println(infoCtx)
	}()

	for _, level := range loaders.Handlers {
		concurrentLoader := concurrent.NewConcurrentLoader(ctx)

		resultHandlerList := make([]loaders.LoaderResultHandler, 0, len(level.HandlerList))
		for _, h := range level.HandlerList {
			resultHandler, err := h(ctx, infoCtx, concurrentLoader)
			if err != nil {
				return err
			}
			resultHandlerList = append(resultHandlerList, resultHandler)
		}

		// load
		concurrentLoader.ConcurrentLoad(level.TimeOut)

		// resultHandler
		for _, resultHandler := range resultHandlerList {
			if err := resultHandler(); err != nil {
				return err
			}
		}
	}

	return nil
}
