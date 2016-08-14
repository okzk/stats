package stats

import (
	"encoding/json"
	"github.com/okzk/ticker"
	"runtime"
	"time"
)

type Stats struct {
	Time         int64
	NumGoroutine int
	MemStats     *runtime.MemStats
}

func CurrentStats() *Stats {
	mem := new(runtime.MemStats)
	runtime.ReadMemStats(mem)

	return &Stats{
		Time:         time.Now().UnixNano(),
		NumGoroutine: runtime.NumGoroutine(),
		MemStats:     mem,
	}
}

func (s *Stats) String() string {
	json, _ := json.Marshal(s)
	return string(json)
}

func (s *Stats) Update() {
	s.Time = time.Now().UnixNano()
	s.NumGoroutine = runtime.NumGoroutine()
	runtime.ReadMemStats(s.MemStats)
}

func SchedulePeriodically(duration time.Duration, proc func(*Stats)) *ticker.Ticker {
	s := CurrentStats()
	return ticker.New(duration, func(_ time.Time) {
		s.Update()
		proc(s)
	})
}
