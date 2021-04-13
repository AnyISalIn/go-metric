package output

import "github.com/AnyISalIn/go-metric/pkg/collector"

type Endpoint interface {
	Write([]collector.Metric) error
}
