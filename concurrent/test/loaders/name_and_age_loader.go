package loaders

import (
	"context"
	"strconv"
	"wangpeilei/leetcode/concurrent"
)

func NameAndAgeLoaderHandler(ctx context.Context, info *InfoCtx, cl *concurrent.ConcurrentLoader) (LoaderResultHandler, error) {
	if info.Name == "" || info.Age == 0 {
		return EmptyResultHandler, nil // 弱依赖不返回err
	}

	l := &NameAndAgeLoader{
		Name: info.Name,
		Age:  info.Age,
	}

	cl.AppendLoader(l)

	return func() error {
		if l.GetError() != nil {
			// warn log
			return nil
		}
		info.NameAge = l.NameAge
		return nil
	}, nil
}

type NameAndAgeLoader struct {
	concurrent.CommonLoader

	// req
	Name string
	Age  int

	// resp
	NameAge string
}

func (l *NameAndAgeLoader) Load() error {
	l.NameAge = l.Name + strconv.FormatInt(int64(l.Age), 10)
	return nil
}
