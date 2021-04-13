package collector

import (
	"fmt"
	"github.com/mackerelio/go-osstat/cpu"
	"os"
	"time"
)

type CpuCollector struct {
}

func (c *CpuCollector) Collect() ([]Metric, error) {
	before, err := cpu.Get()
	var metrics []Metric
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return metrics, err
	}
	time.Sleep(time.Duration(1) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return metrics, err
	}
	currentTimestamp := time.Now()
	total := float64(after.Total - before.Total)
	metrics = append(metrics, Metric{
		Timestamp: currentTimestamp,
		Key:       "cpu_user",
		Value:     float64(after.User-before.User) / total * 100,
		Labels:    map[string]string{},
	})
	metrics = append(metrics, Metric{
		Timestamp: currentTimestamp,
		Key:       "cpu_system",
		Value:     float64(after.System-before.System) / total * 100,
		Labels:    map[string]string{},
	})

	metrics = append(metrics, Metric{
		Timestamp: currentTimestamp,
		Key:       "cpu_idle",
		Value:     float64(after.Idle-before.Idle) / total * 100,
		Labels:    map[string]string{},
	})

	return metrics, err
}
