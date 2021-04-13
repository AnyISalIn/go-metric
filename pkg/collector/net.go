package collector

import (
	"github.com/mackerelio/go-osstat/network"
	"time"
)

type NetCollector struct {
}

func (c *NetCollector) Collect() ([]Metric, error) {
	networks, err := network.Get()
	var metrics []Metric
	currentTimestamp := time.Now()
	for _, network := range networks {
		metrics = append(metrics, Metric{
			Timestamp: currentTimestamp,
			Key:       "tx",
			Value:     float64(network.TxBytes),
			Labels: map[string]string{
				"nic": network.Name,
			},
		})
		metrics = append(metrics, Metric{
			Timestamp: currentTimestamp,
			Key:       "rx",
			Value:     float64(network.RxBytes),
			Labels: map[string]string{
				"nic": network.Name,
			},
		})
	}
	return metrics, err
}
