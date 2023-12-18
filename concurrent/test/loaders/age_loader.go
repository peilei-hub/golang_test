package loaders

import (
	"context"
	"errors"
	"time"
	"wangpeilei/leetcode/concurrent"
)

func AgeLoaderHandler(ctx context.Context, info *InfoCtx, cl *concurrent.ConcurrentLoader) (LoaderResultHandler, error) {
	if info.Id == 0 {
		return EmptyResultHandler, errors.New("id is nil")
	}

	l := &AgeLoader{
		Id: info.Id,
	}

	cl.AppendLoader(l)

	return func() error {
		if l.GetError() != nil {
			return errors.New("age_loader error: " + l.GetError().Error())
		}
		info.Age = l.Resp
		return nil
	}, nil
}

type AgeLoader struct {
	concurrent.CommonLoader
	Id int

	Resp int
}

func (l *AgeLoader) Load() error {
	if l.Id == 1 {
		time.Sleep(2 * time.Second) // fixme
	} else if l.Id == 2 {
		return errors.New("age err")
	}
	l.Resp = 11
	return nil
}
