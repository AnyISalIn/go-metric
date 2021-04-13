package main

import (
	"flag"
	"fmt"
	"github.com/AnyISalIn/go-metric/pkg/collector"
	"github.com/AnyISalIn/go-metric/pkg/output"
	"time"
)

func startCollector(c collector.Collector, ch chan []collector.Metric) {
	metrics, err := c.Collect()
	if err != nil {
	}
	ch <- metrics
}

func writeEndpoint(url string, token string, ch chan []collector.Metric) {
	i := output.NewInfluxEndpoint(url, token)
	defer i.Client.Close()
	for {
		select {
		case metrics := <-ch:
			i.Write(metrics)
		}
	}
}

func main() {
	endpoint := flag.String("endpoint", "http://localhost:8086", "influx api endpoint url")
	token := flag.String("token", "token", "influx api token")
	interval := flag.Int("interval", 10, "collect metric interval")

	flag.Parse()

	c := collector.New()
	ch := make(chan []collector.Metric)
	quit := make(chan struct{})

	ticker := time.NewTicker(time.Duration(*interval) * time.Second)
	go writeEndpoint(*endpoint, *token, ch)

	go func() {
		for {
			select {
			case <-ticker.C:
				go startCollector(&c.Memory, ch)
				go startCollector(&c.Cpu, ch)
				go startCollector(&c.Net, ch)
			}
		}
	}()
	fmt.Println("Starting metric collector...")
	<-quit

}
