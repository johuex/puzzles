package ratelimiter

import (
	"sync"
	"time"
)

// MUTEX impl

type MutexLimiter struct {
	mu     sync.Mutex
	base   int // init value
	tokens int // actual token counter
	rate   time.Duration
	status chan struct{}
}

func NewMutexLimiter(cnt int, rate time.Duration) *MutexLimiter {
	l := &MutexLimiter{
		base:   cnt,
		tokens: cnt,
		rate:   rate,
		status: make(chan struct{}),
	}
	go l.refill()

	return l
}

func (l *MutexLimiter) refill() {
	ticker := time.NewTicker(l.rate)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			l.mu.Lock()
			l.tokens = l.base // NOTE: made simple special
			l.mu.Unlock()
		case <-l.status:
			return
		}
	}
}

func (l *MutexLimiter) Close() {
	close(l.status)
}

func (l *MutexLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.tokens > 0 {
		l.tokens--
		return true
	}
	return false
}

// CHANNELS impl

type ChannelLimiter struct {
	tokens chan struct{}
	status chan struct{}
	rate   time.Duration
}

func NewChannelLimiter(capacity int, rate time.Duration) *ChannelLimiter {
	cl := &ChannelLimiter{
		tokens: make(chan struct{}, capacity),
		status: make(chan struct{}),
		rate:   rate,
	}

	// fill init tokens
	for i := 0; i < capacity; i++ {
		cl.tokens <- struct{}{}
	}

	go cl.refill()

	return cl
}

func (l *ChannelLimiter) refill() {
	ticker := time.NewTicker(l.rate)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			select {
			case l.tokens <- struct{}{}: // add token smooth
			default:
				// Ведро полное, пропускаем
			}
		case <-l.status:
			return
		}
	}
}

func (l *ChannelLimiter) Allow() bool {
	select {
	case <-l.tokens:
		return true
	default:
		return false
	}
}

func (l *ChannelLimiter) Close() {
	close(l.status)
}
