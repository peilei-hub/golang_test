package concurrent

import (
	"context"
	"time"
)

type ConcurrentLoader struct {
	ctx     context.Context
	loaders []ILoader
}

type loaderResult struct {
	Id  int
	Err error
}

type loaderInfo struct {
	Id     int
	Loader ILoader
}

func NewConcurrentLoader(ctx context.Context) *ConcurrentLoader {
	l := &ConcurrentLoader{
		ctx: ctx,
	}

	return l
}

func (c *ConcurrentLoader) AppendLoader(loaders ...ILoader) {
	c.loaders = append(c.loaders, loaders...)
}

func (c *ConcurrentLoader) ConcurrentLoad(timeout time.Duration) {
	loaderCount := len(c.loaders)
	resultsChan := make(chan *loaderResult, loaderCount) // 存放结果的chan

	loaderInfoMap := map[int]*loaderInfo{}

	for i, loader := range c.loaders {
		info := &loaderInfo{
			Id:     i,
			Loader: loader,
		}
		loaderInfoMap[i] = info

		go func(info *loaderInfo) {
			resultsChan <- &loaderResult{
				Id:  info.Id,
				Err: info.Loader.Load(),
			}
		}(info)
	}

	timer := time.After(timeout)
	resultsMap := map[int]*loaderResult{}

LoaderLoop:
	for len(resultsMap) < loaderCount {
		select {
		case result := <-resultsChan:
			resultsMap[result.Id] = result
			if result.Err != nil {
				loaderInfoMap[result.Id].Loader.SetError(result.Err)
			}
		case <-timer:
			for loaderId, info := range loaderInfoMap {
				if _, ok := resultsMap[loaderId]; !ok {
					info.Loader.SetError(ConcurrentLoadTimeOutError)
				}
			}
			break LoaderLoop
		}
	}
}
