package metrics

import "sync"

type Metrics struct {
	lock    sync.RWMutex
	metrics map[string]float64
}
