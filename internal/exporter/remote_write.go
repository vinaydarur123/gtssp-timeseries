package exporter

import (
	"bytes"
	"net/http"
	"time"

	"github.com/your-org/gtssp/internal/model"
)

// RemoteWriteExporter is a partial implementation
// of Prometheus Remote Write exporter
type RemoteWriteExporter struct {
	endpoint string
	client   *http.Client
}

// NewRemoteWriteExporter initializes the exporter
func NewRemoteWriteExporter(endpoint string) *RemoteWriteExporter {
	return &RemoteWriteExporter{
		endpoint: endpoint,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// Export prepares metrics for remote write (partial)
func (r *RemoteWriteExporter) Export(metrics []model.Metric) error {
	// TODO:
	// 1. Convert metrics to Prometheus TimeSeries
	// 2. Marshal using protobuf
	// 3. Compress using snappy
	// 4. POST to remote write endpoint

	req, err := http.NewRequest(
		http.MethodPost,
		r.endpoint,
		bytes.NewBuffer([]byte("remote-write-payload")),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-protobuf")
	req.Header.Set("Content-Encoding", "snappy")

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
