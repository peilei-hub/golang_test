package main

import (
	"fmt"
	"time"
)

type TokenBucketLimiter struct {
	capacity      int       // 容量
	currentTokens int       // 令牌数量
	rate          int       // 发令牌速率 s
	lastTime      time.Time // 上次发令牌时间
}

func NewTokenBucketLimiter(capacity, rate int) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		capacity: capacity,
		rate:     rate,
		lastTime: time.Now(),
	}
}

func (l *TokenBucketLimiter) Allow() bool {
	// 尝试发放令牌
	now := time.Now()
	// 距离上次发放令牌时间
	interval := now.Sub(l.lastTime)
	if interval >= time.Second {
		l.currentTokens = l.currentTokens + int(interval/time.Second)*l.rate
		if l.currentTokens > l.capacity {
			l.currentTokens = l.capacity
		}
		l.lastTime = now
	}

	// 没有令牌
	if l.currentTokens == 0 {
		return false
	}
	l.currentTokens--
	return true
}

func main() {
	limiter := NewTokenBucketLimiter(20, 3)
	time.Sleep(2 * time.Second) // 放2s令牌
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
