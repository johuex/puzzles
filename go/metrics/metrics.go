package metrics

import (
	"sync/atomic"
)

// unused now
type I interface {
	Add(k string, v int64)
	Get() int64
}

type Metrics struct {
	shards      []atomic.Int64
	keyShardMap map[string]int
}

func NewMetrics(keys []string) *Metrics {
	keyShardMap := make(map[string]int)
	for idx, val := range keys {
		keyShardMap[val] = idx
	}

	return &Metrics{
		shards:      make([]atomic.Int64, len(keys)),
		keyShardMap: keyShardMap,
	}
}

func (m *Metrics) Add(key string, v int64) {
	shardID := m.keyShardMap[key]
	m.shards[shardID].Add(v)
}

func (m *Metrics) Get(key string) int64 {
	shardID := m.keyShardMap[key]
	return m.shards[shardID].Load()
}
