package durtymemory

import (
	"testing"
)

var mem []int

const size = 1 << 20

func BenchmarkGetMem(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		mem = GetMem(size)
	}

}

func BenchmarkGetDurtyMem(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		mem = GetDurtyMem(size)
	}

}
