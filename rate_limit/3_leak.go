package main

import (
	"fmt"
	"time"
)

type LeakBucketLimiter struct {
	peakLevel    int       // 最大水位，桶的容量
	currentLevel int       // 当前水位
	rate         int       // 水流速度 s
	lastTime     time.Time // 上次放水时间
}

func NewLeadBucketLimiter(peakLevel, rate int) *LeakBucketLimiter {
	return &LeakBucketLimiter{
		peakLevel: peakLevel,
		rate:      rate,
		lastTime:  time.Now(),
	}
}

func (l *LeakBucketLimiter) Allow() bool {
	// 尝试放水
	now := time.Now()
	// 距离上次放水时间
	intervals := now.Sub(l.lastTime)
	if intervals >= time.Second {
		l.currentLevel = l.currentLevel - int(intervals/time.Second)*l.rate // 上次到这次放了多少水
		if l.currentLevel < 0 {
			l.currentLevel = 0
		}
		l.lastTime = now
	}

	// 到达最高水位
	if l.currentLevel >= l.peakLevel {
		return false
	}
	l.currentLevel++
	return true
}

func main() {
	limiter := NewLeadBucketLimiter(20, 3)
	time.Sleep(10 * time.Second)
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
