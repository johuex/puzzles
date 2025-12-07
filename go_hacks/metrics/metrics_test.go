package metrics

import (
	"math/rand"
	"runtime"
	"sync"
	"testing"
)

func randomKeys(cnt, length int) []string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	res := make([]string, cnt)

	for i := range cnt {
		key := make([]byte, length)
		for j := range key {
			key[j] = letters[rand.Intn(len(letters))]
		}
		res[i] = string(key)
	}

	return res
}

// test speed of Metrics structure, each metric has own shard
func BenchmarkMetricsRandomKeys(b *testing.B) {
	numWorkers := runtime.NumCPU()
	keyLen := 8
	keysCnt := 1 << 20

	keys := randomKeys(keysCnt, keyLen)

	metrics := NewMetrics(keys)

	b.ResetTimer()
	var wg sync.WaitGroup

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx, val := range keys {
				i := min(idx, len(keys))
				metrics.Get(keys[i])

				metrics.Add(val, 1)

				d := max(idx, 0)
				metrics.Get(keys[d])

			}
		}()
	}

	wg.Wait()
}
