package ratelimiter

import (
	"sync"
	"testing"
	"time"
)

func TestMutexLimiter(t *testing.T) {
	limiter := NewMutexLimiter(3, 50*time.Millisecond)
	defer limiter.Close()

	for i := 0; i < 3; i++ {
		if !limiter.Allow() {
			t.Errorf("Запрос %d должен быть разрешен", i)
		}
	}

	if limiter.Allow() {
		t.Error("4-й запрос должен быть отклонен")
	}

	// wait three new tokens
	time.Sleep(100 * time.Millisecond)

	for i := 0; i < 3; i++ {
		if !limiter.Allow() {
			t.Errorf("Запрос %d после ожидания должен быть разрешен", i)
		}
	}

	// rate exceed again
	if limiter.Allow() {
		t.Error("Запрос после исчерпания должен быть отклонен")
	}
}

func TestChannelLimiter(t *testing.T) {
	limiter := NewChannelLimiter(3, 50*time.Millisecond)
	defer limiter.Close()

	for i := 0; i < 3; i++ {
		if !limiter.Allow() {
			t.Errorf("Запрос %d должен быть разрешен", i)
		}
	}

	if limiter.Allow() {
		t.Error("4-й запрос должен быть отклонен")
	}

	// wait three new tokens
	time.Sleep(150 * time.Millisecond)

	for i := 0; i < 3; i++ {
		if !limiter.Allow() {
			t.Errorf("Запрос %d после ожидания должен быть разрешен", i)
		}
	}

	// rate exceed again
	if limiter.Allow() {
		t.Error("Запрос после исчерпания должен быть отклонен")
	}
}

func TestConcurrentMutex(t *testing.T) {
	limiter := NewMutexLimiter(1000, time.Microsecond)
	defer limiter.Close()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				limiter.Allow()
			}
		}()
	}
	wg.Wait()
}

func TestConcurrentChannel(t *testing.T) {
	limiter := NewChannelLimiter(1000, time.Microsecond)
	defer limiter.Close()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				limiter.Allow()
			}
		}()
	}
	wg.Wait()
}

// example with huge token inserting

func BenchmarkMutexLimiter(b *testing.B) {
	limiter := NewMutexLimiter(1000000, time.Nanosecond)
	defer limiter.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			limiter.Allow()
		}
	})
}

func BenchmarkChannelLimiter(b *testing.B) {
	limiter := NewChannelLimiter(1000000, time.Nanosecond)
	defer limiter.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			limiter.Allow()
		}
	})
}

// real usage examples

func BenchmarkMutexLimiterReal(b *testing.B) {
	limiter := NewMutexLimiter(5, 200*time.Millisecond)
	defer limiter.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			limiter.Allow()
		}
	})
}

// Бенчмарк для каналов с реальной частотой
func BenchmarkChannelLimiterReal(b *testing.B) {
	limiter := NewChannelLimiter(5, 200*time.Millisecond) // 10 RPS
	defer limiter.Close()

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			limiter.Allow()
		}
	})
}
