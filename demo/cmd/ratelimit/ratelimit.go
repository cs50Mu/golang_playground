package main

import (
	"fmt"
	"sync"
	"time"
)

type CounterLimiter struct {
	Interval  time.Duration
	Threshold int64
	currCnt   int64
	startTime time.Time
	mu        sync.Mutex
}

func NewCounterLimiter(interval time.Duration, threshold int64) *CounterLimiter {
	return &CounterLimiter{
		Interval:  interval,
		Threshold: threshold,
		startTime: time.Now(),
	}
}

func (cl *CounterLimiter) Allow() bool {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	now := time.Now()
	// 区间内
	if cl.startTime.Add(cl.Interval).After(now) {
		// fmt.Printf("in range, now: %+v\n", now)
		cl.currCnt++
		return cl.currCnt <= cl.Threshold
	} else {
		cl.startTime = now
		cl.currCnt = 1
		return true
	}
}

// 漏桶限流算法
// 输入速率任意，输出速率固定，当桶满时，拒绝访问
type LeakBucketLimiter struct {
	cap        int64 // 桶容量
	lastLeakTs int64 // 上次计算漏水的时间戳
	leakRate   int64 // 漏水的速率，单位：秒
	water      int64 // 当前桶内的水量
	mu         sync.Mutex
}

func NewLeakBucketLimiter(cap, leakRate int64) *LeakBucketLimiter {
	return &LeakBucketLimiter{
		cap:        cap,
		leakRate:   leakRate,
		lastLeakTs: time.Now().Unix(),
	}
}

func (lb *LeakBucketLimiter) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	// 计算当前到上次请求期间漏掉的水
	currTs := time.Now().Unix()
	elapsedSecs := currTs - lb.lastLeakTs
	if elapsedSecs < 0 {
		elapsedSecs = 0
	}
	waterLeaked := elapsedSecs * lb.leakRate
	fmt.Printf("leaked: %v, currWater: %v\n", waterLeaked, lb.water)
	// 更新水量
	lb.water = lb.water - waterLeaked
	// 注意这里把水量重置为零
	if lb.water < 0 {
		lb.water = 0
	}
	// 更新时间戳为当前时间
	// 因为水量已经更新过了，下次计算新的水量肯定要根据最新的时间戳才行
	lb.lastLeakTs = currTs
	// 若水量已经超过桶容量则拒绝
	// 若水量尚未超过桶容量则放行
	if lb.water < lb.cap {
		// 本次请求会增加水量
		lb.water++
		return true
	} else {
		return false
	}
}

// TokenBucketLimiter 令牌桶算法
type TokenBucketLimiter struct {
	cap        int64 // 桶容量
	rate       int64 // 令牌产生的速率
	ts         int64 // 时间戳
	currTokens int64 // 当前桶内的 token 数
	mu         sync.Mutex
}

func NewTokenBucketLimiter(cap, rate int64) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		cap:  cap,
		rate: rate,
		// 初始的时候，桶是满的
		currTokens: cap,
		ts:         time.Now().Unix(),
	}
}

func (tb *TokenBucketLimiter) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now().Unix()
	elapsedSecs := now - tb.ts
	filledTokens := elapsedSecs * tb.rate
	tb.ts = now
	tb.currTokens += filledTokens
	if tb.currTokens > tb.cap {
		tb.currTokens = tb.cap
	}
	if tb.currTokens == 0 {
		return false
	} else {
		tb.currTokens--
		return true
	}
}
