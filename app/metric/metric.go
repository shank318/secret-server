package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	Namespace = "router"
)

var register = prometheus.MustRegister
var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Name:      "requests",
			Help:      "How many HTTP requests processed, partitioned by status code and HTTP method.",
		},
		[]string{"status", "method", "handler"},
	)

	TotalRequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Name:      "requests_total",
			Help:      "How many HTTP requests received",
		},
		[]string{"handler"},
	)

	ErrorCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Name:      "rzp_errors",
			Help:      "Rzp error",
		},
		[]string{"error"},
	)

	ExternalClient = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Name:      "external_requests_count",
			Help:      "How many HTTP requests processed by external clients, partitioned by uri",
		},
		[]string{"uri", "status"},
	)
	TimeToProcess = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: Namespace,
			Name:      "response_time",
			Help:      "Time taken to process request and respond back",
			Buckets:   prometheus.LinearBuckets(5, 5, 20), // 20 buckets, each 5 wide starting 5
		})
)

func RegisterPrometheusMetrics() {
	register(RequestCount)
	register(TotalRequestCount)
	register(ErrorCounter)
	register(ExternalClient)
}
