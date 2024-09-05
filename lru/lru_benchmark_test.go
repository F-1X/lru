package lru

import (
	"crypto/rand"
	"math"
	"math/big"
	"testing"
	"time"
)

func BenchmarkAdd(b *testing.B) {
	c := NewCache(1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.AddWithTTL(i, i, time.Minute)
	}
}

func BenchmarkAddAndGet(b *testing.B) {
	c := NewCache(1000)

	b.Run("Adding", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.AddWithTTL(i, i, time.Minute)
		}
	})

	b.Run("Getting", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Get(i)
		}
	})
}

func BenchmarkRandomAddAndGet(b *testing.B) {
	c := NewCache(1000)

	trace := make([]int64, b.N*2)
	for i := 0; i < b.N*2; i++ {
		trace[i] = getRand(b) % 32768
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.AddWithTTL(trace[i], trace[i], time.Minute)
		_, _ = c.Get(trace[i])
	}
}

func BenchmarkLRU_Rand_NoExpire(b *testing.B) {
	l := NewCache(8192)

	trace := make([]int64, b.N*2)
	for i := 0; i < b.N*2; i++ {
		trace[i] = getRand(b) % 32768
	}

	b.ResetTimer()

	var hit, miss int
	for i := 0; i < 2*b.N; i++ {
		if i%2 == 0 {
			l.Add(trace[i], trace[i])
		} else {
			if _, ok := l.Get(trace[i]); ok {
				hit++
			} else {
				miss++
			}
		}
	}
	b.Logf("hit: %d miss: %d ratio: %f", hit, miss, float64(hit)/float64(hit+miss))
}

func getRand(tb testing.TB) int64 {
	out, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		tb.Fatal(err)
	}
	return out.Int64()
}
