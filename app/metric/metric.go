package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	Namespace = "secretserver"
)

var register = prometheus.MustRegister
var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Name:      "requests",
			Help:      "How many HTTP requests processed, partitioned by api",
		},
		[]string{"api"},
	)

	ErrorCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Name:      "rzp_errors",
			Help:      "Rzp error",
		},
		[]string{"error"},
	)

	TimeToProcess = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: Namespace,
			Name:      "response_time_add_secret",
			Help:      "Time taken to process request and respond back",
			Buckets:   prometheus.LinearBuckets(5, 5, 20), // 20 buckets, each 5 wide starting 5
		}, []string{"api"})

)

func RegisterPrometheusMetrics() {
	register(RequestCount)
	register(ErrorCounter)
	register(TimeToProcess)
}
