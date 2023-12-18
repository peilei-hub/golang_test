package concurrent

import (
	"errors"
)

var (
	ConcurrentLoadTimeOutError = errors.New("ConcurrentLoadTimeOutError")
)
