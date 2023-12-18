package main

import (
	"fmt"
	"time"
)

type CounterLimiter struct {
	limit    int
	window   time.Duration // 窗口时间大小
	counter  int
	lastTime time.Time // 时间
}

func NewCounterLimiter(limit int, window time.Duration) *CounterLimiter {
	return &CounterLimiter{
		limit:    limit,
		window:   window,
		lastTime: time.Now(),
	}
}

func (l *CounterLimiter) Allow() bool {
	now := time.Now()
	if now.Sub(l.lastTime) <= l.window { // 在window内
		if l.counter > l.limit {
			return false
		}
		l.counter++
	} else {
		l.lastTime = now
		l.counter = 1
	}
	return l.counter <= l.limit
}

func main() {
	limiter := NewCounterLimiter(3, time.Second)
	var allow, notAllow int
	for {
		if limiter.Allow() {
			fmt.Println("allow")
			allow++
		} else {
			fmt.Println("now allow")
			notAllow++
		}
		time.Sleep(200 * time.Millisecond)
	}
}
