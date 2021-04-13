package collector

import (
	"testing"
)

func TestCollectCpu(t *testing.T) {
	c := CpuCollector{}
	_, err := c.Collect()
	if err != nil {
		t.Fatalf("error should be nil but got: %v", err)
	}
}
