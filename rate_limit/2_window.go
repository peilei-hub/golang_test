package main

import (
	"errors"
	"fmt"
	"time"
)

type SlideWindowLimiter struct {
	limit           int
	window          int64         // 窗口时间大小
	smallWindow     int64         // 小窗口时间大小
	smallWindowNums int           // 小窗口数量
	counters        map[int64]int // 小窗口计数器
}

func NewSlideWindowLimiter(limit int, window, smallWindow time.Duration) (*SlideWindowLimiter, error) {
	if window%smallWindow != 0 {
		return nil, errors.New("err")
	}

	return &SlideWindowLimiter{
		limit:           limit,
		window:          int64(window),
		smallWindow:     int64(smallWindow),
		smallWindowNums: int(window / smallWindow),
		counters:        make(map[int64]int),
	}, nil
}

func (l *SlideWindowLimiter) Allow() bool {
	// 获取当前小窗口值
	currentSmallWindow := time.Now().UnixNano() / l.smallWindow * l.smallWindow
	// 获取起始小窗口值
	startSmallWindow := currentSmallWindow - l.smallWindow*int64(l.smallWindowNums-1)

	var count int
	// 计算当前窗口的请求总数
	for smallWindow, counter := range l.counters {
		if smallWindow < startSmallWindow {
			delete(l.counters, smallWindow)
		} else {
			count += counter
		}
	}

	if count >= l.limit {
		return false
	}

	l.counters[currentSmallWindow]++
	return true
}

func main() {
	limiter, _ := NewSlideWindowLimiter(3, time.Second, time.Millisecond*10)
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
