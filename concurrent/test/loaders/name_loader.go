package loaders

import (
	"context"
	"errors"
	"wangpeilei/leetcode/concurrent"
)

func NameLoaderHandler(ctx context.Context, info *InfoCtx, cl *concurrent.ConcurrentLoader) (LoaderResultHandler, error) {
	if info.Id == 0 {
		return EmptyResultHandler, errors.New("id is nil")
	}

	l := &NameLoader{
		Id: info.Id,
	}
	cl.AppendLoader(l)

	return func() error {
		if l.GetError() != nil {
			return l.GetError()
		}

		info.Name = l.Resp
		return nil
	}, nil
}

type NameLoader struct {
	concurrent.CommonLoader
	Id int

	Resp string
}

func (l *NameLoader) Load() error {
	l.Resp = "wpl"
	return nil
}
