package stats

import (
	"testing"
)

func BenchmarkStats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CurrentStats()
	}
}

func BenchmarkStatsUpdate(b *testing.B) {
	s := CurrentStats()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Update()
	}
}

func BenchmarkStatsJson(b *testing.B) {
	s := CurrentStats()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.String()
	}
}
