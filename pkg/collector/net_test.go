package collector

import (
	"testing"
)

func TestCollectNet(t *testing.T) {
	n := NetCollector{}
	_, err := n.Collect()
	if err != nil {
		t.Fatalf("error should be nil but got: %v", err)
	}
}
