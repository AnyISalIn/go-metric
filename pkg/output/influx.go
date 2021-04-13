package output

import (
	"github.com/AnyISalIn/go-metric/pkg/collector"
	"github.com/influxdata/influxdb-client-go"
	"os"
)

type InfluxEndpoint struct {
	Client influxdb2.Client
}

func (i *InfluxEndpoint) Write(metrics []collector.Metric) error {
	writeAPI := i.Client.WriteAPI("my-org", "mydb")
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	for _, metric := range metrics {
		tags := metric.Labels
		tags["hostname"] = hostname

		p := influxdb2.NewPoint(
			"system",
			tags,
			map[string]interface{}{
				metric.Key: metric.Value,
			},
			metric.Timestamp)

		writeAPI.WritePoint(p)
	}
	writeAPI.Flush()
	return nil
}

func NewInfluxEndpoint(URL string, Token string) InfluxEndpoint {
	client := influxdb2.NewClient(URL, Token)
	return InfluxEndpoint{Client: client}
}
