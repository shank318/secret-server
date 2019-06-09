package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterPrometheusMetrics(t *testing.T) {
	registerCount := 0
	oldREgister := register
	defer func() { register = oldREgister }()
	register = func(cs ...prometheus.Collector) {
		registerCount++
	}

	RegisterPrometheusMetrics()
	assert.Equal(t, 3, registerCount)
}
