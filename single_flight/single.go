package single_flight

import (
	"fmt"
	"sync"
)

type Result struct {
	Val    interface{}
	Err    error
	Shared bool
}

type call struct {
	wg sync.WaitGroup

	val interface{}
	err error

	dups  int
	chans []chan<- Result
}

type SingleFlight struct {
	l sync.Mutex
	m map[string]*call
}

func (s *SingleFlight) Do(key string, fn func() (any, error)) (v any, err error, shared bool) {
	s.l.Lock()
	if s.m == nil {
		s.m = make(map[string]*call)
	}

	if c, ok := s.m[key]; ok {
		c.dups++
		s.l.Unlock()
		c.wg.Wait()

		return c.val, c.err, true
	}

	c := new(call)
	c.wg.Add(1)
	s.m[key] = c
	s.l.Unlock()
	s.doCall(c, key, fn)
	return c.val, c.err, c.dups > 0
}

func (s *SingleFlight) DoChan(key string, fn func() (any, error)) <-chan Result {
	ch := make(chan Result, 1)
	s.l.Lock()
	if s.m == nil {
		s.m = make(map[string]*call)
	}
	if c, ok := s.m[key]; ok {
		c.dups++
		c.chans = append(c.chans, ch)
		s.l.Unlock()
		return ch
	}

	c := new(call)
	c.wg.Add(1)
	c.chans = []chan<- Result{ch}
	s.m[key] = c
	s.l.Unlock()
	go s.doCall(c, key, fn)

	return ch
}

func (s *SingleFlight) doCall(c *call, key string, fn func() (any, error)) {

	defer func() {
		s.l.Lock()
		defer s.l.Unlock()

		c.wg.Done()
		if s.m[key] == c {
			delete(s.m, key)
		}
		for _, ch := range c.chans {
			ch <- Result{c.val, c.err, c.dups > 0}
		}
	}()
	fmt.Println("in doCall")
	c.val, c.err = fn()
}

func (s *SingleFlight) Forget(key string) {
	s.l.Lock()
	delete(s.m, key)
	s.l.Unlock()
}
