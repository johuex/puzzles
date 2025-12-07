package metrics

import (
	"sync/atomic"
)

// unused now
type I interface {
	Add(k string, v int64)
	Get() int64
}

// shard
type shard struct {
	value int64
	_     [40]byte // padding
}

func newShard() *shard {
	return &shard{}
}

type Metrics struct {
	shards      []*shard
	keyShardMap map[string]int
}

func NewMetrics(keys []string) *Metrics {
	nKeys := len(keys)
	keyShardMap := make(map[string]int)
	for idx, val := range keys {
		keyShardMap[val] = idx
	}

	shards := make([]*shard, nKeys)
	for i := 0; i < nKeys; i++ {
		shards[i] = newShard()
	}
	return &Metrics{
		shards:      shards,
		keyShardMap: keyShardMap,
	}
}

func (m *Metrics) Add(key string, v int64) {
	shardID := m.keyShardMap[key]
	atomic.AddInt64(&m.shards[shardID].value, v)
}

func (m *Metrics) Get(key string) int64 {
	shardID := m.keyShardMap[key]
	return atomic.LoadInt64(&m.shards[shardID].value)
}
