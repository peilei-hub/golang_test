package concurrent

type ILoader interface {
	Load() error
	SetError(error)
	GetError() error
}

type CommonLoader struct {
	err error
}

func (c *CommonLoader) SetError(err error) {
	c.err = err
}

func (c *CommonLoader) GetError() error {
	return c.err
}
