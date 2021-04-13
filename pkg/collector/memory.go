package collector

import (
	"fmt"
	"github.com/mackerelio/go-osstat/memory"
	"os"
	"time"
)

type MemoryCollector struct {
}

func (m *MemoryCollector) Collect() ([]Metric, error) {
	var metrics []Metric

	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return metrics, err
	}
	currentTimestamp := time.Now()
	metrics = append(metrics, Metric{
		Timestamp: currentTimestamp,
		Key:       "memory_total",
		Value:     float64(memory.Total),
		Labels:    map[string]string{},
	})
	metrics = append(metrics, Metric{
		Timestamp: currentTimestamp,
		Key:       "memory_used",
		Value:     float64(memory.Used),
		Labels:    map[string]string{},
	})

	metrics = append(metrics, Metric{
		Timestamp: currentTimestamp,
		Key:       "memory_cached",
		Value:     float64(memory.Cached),
		Labels:    map[string]string{},
	})

	metrics = append(metrics, Metric{
		Timestamp: currentTimestamp,
		Key:       "memory_free",
		Value:     float64(memory.Free),
		Labels:    map[string]string{},
	})

	return metrics, nil

}
