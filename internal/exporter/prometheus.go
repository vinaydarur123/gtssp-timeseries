package exporter

import (
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/your-org/gtssp/internal/model"
)

// Exporter defines how metrics are exported
type Exporter interface {
	Export(metrics []model.Metric) error
}

// PrometheusExporter exposes metrics via /metrics endpoint
type PrometheusExporter struct {
	mu     sync.Mutex
	gauges map[string]*prometheus.GaugeVec
}

// NewPrometheusExporter creates exporter
func NewPrometheusExporter() *PrometheusExporter {
	return &PrometheusExporter{
		gauges: make(map[string]*prometheus.GaugeVec),
	}
}

// Export registers and updates Prometheus metrics
func (e *PrometheusExporter) Export(metrics []model.Metric) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	for _, m := range metrics {
		gv, exists := e.gauges[m.Name]
		if !exists {
			labelKeys := make([]string, 0, len(m.Labels))
			for k := range m.Labels {
				labelKeys = append(labelKeys, k)
			}

			gv = prometheus.NewGaugeVec(
				prometheus.GaugeOpts{
					Name: m.Name,
					Help: "GTSPP generated metric",
				},
				labelKeys,
			)

			prometheus.MustRegister(gv)
			e.gauges[m.Name] = gv
		}

		gv.With(m.Labels).Set(m.Value)
	}
	return nil
}

// StartHTTPServer starts /metrics endpoint
func (e *PrometheusExporter) StartHTTPServer(addr string) {
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(addr, nil)
}
