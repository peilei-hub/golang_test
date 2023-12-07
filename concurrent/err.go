package concurrent

import (
	"context"
	"errors"
)

var (
	ConcurrentLoadTimeOutError = errors.New("ConcurrentLoadTimeOutError")
)

func parseError(ctx context.Context, i interface{}) error {
	switch r := i.(type) {
	case string:
		return errors.New(r)
	case error:
		return r
	default:
		return errors.New("unknown")
	}
}
