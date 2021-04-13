package collector

import (
	"testing"
)

func TestCollectMemory(t *testing.T) {
	m := MemoryCollector{}
	_, err := m.Collect()
	if err != nil {
		t.Fatalf("error should be nil but got: %v", err)
	}
}
