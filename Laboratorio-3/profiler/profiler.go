package profiler

import (
	"runtime"
	"time"
)

func MeasureMemoryUsage(f func() error) (time.Duration, uint64, error) {
	var memStart, memEnd runtime.MemStats

	runtime.GC() // Forzar recolección de basura para valores más precisos
	runtime.ReadMemStats(&memStart)

	start := time.Now()
	err := f()
	duration := time.Since(start)

	runtime.ReadMemStats(&memEnd)
	memUsed := memEnd.Alloc - memStart.Alloc // Memoria usada en bytes

	return duration, memUsed, err
}
