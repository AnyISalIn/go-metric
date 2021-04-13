package collector

import "time"

type Metric struct {
	Timestamp time.Time
	Key       string
	Value     float64
	Labels    map[string]string
}

type Collector interface {
	Collect() ([]Metric, error)
}

type CollectorServices struct {
	Cpu    CpuCollector
	Memory MemoryCollector
	Net    NetCollector
}

func New() CollectorServices {
	return CollectorServices{
		CpuCollector{},
		MemoryCollector{},
		NetCollector{},
	}
}
