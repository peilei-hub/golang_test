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

func (c *ConcurrentLoader) run(ctx context.Context, concurrentSize int, resultsChain chan<- *loaderResult, messagesChan <-chan *loaderInfo, timeout time.Duration) context.CancelFunc {
	cancelCtx, cancelFunc := context.WithTimeout(ctx, timeout)
	//concurrentCh := make(chan struct{}, concurrentSize)

	go func(ctx context.Context) {
	Exit:
		for {
			select {
			case <-ctx.Done():
				break Exit

			// create a goroutine to run loaders task
			case m := <-messagesChan:
				//concurrentCh <- struct{}{}

				go func(l ILoader, loaderId int) {
					//defer func() {
					//	<-concurrentCh
					//}()
					resultsChain <- &loaderResult{
						Id:  loaderId,
						Err: l.Load(),
					}
				}(m.Loader, m.Id)
			}
		}
	}(cancelCtx)

	return cancelFunc
}

func (c *ConcurrentLoader) ConcurrentLoad(timeout time.Duration) {
	loaderCount := len(c.loaders)
	resultsChan := make(chan *loaderResult, loaderCount)
	messagesChan := make(chan *loaderInfo, loaderCount)

	loaderMap := map[int]*loaderInfo{}

	for i, loader := range c.loaders {
		info := &loaderInfo{
			Id:     i,
			Loader: loader,
		}
		loaderMap[i] = info
		messagesChan <- info
	}

	cancelFunc := c.run(context.Background(), loaderCount, resultsChan, messagesChan, timeout)

	timer := time.After(timeout)
	resultsMap := map[int]*loaderResult{}

LoaderLoop:
	for len(resultsMap) < loaderCount {
		select {
		case result := <-resultsChan:
			resultsMap[result.Id] = result
			if result.Err != nil {
				loaderMap[result.Id].Loader.SetError(result.Err)
			}
		case <-timer:
			for loaderId, info := range loaderMap {
				if _, ok := resultsMap[loaderId]; !ok {
					info.Loader.SetError(ConcurrentLoadTimeOutError)
				}
			}
			break LoaderLoop
		}
	}

	cancelFunc()
}
