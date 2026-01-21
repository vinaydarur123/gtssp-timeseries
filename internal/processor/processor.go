package processor

import (
	"errors"
	"time"

	"github.com/your-org/gtssp/internal/model"
	"github.com/your-org/gtssp/internal/util"
)

// Processor validates and normalizes incoming metrics
type Processor interface {
	Process(metric model.Metric) (model.Metric, error)
}

// BasicProcessor implements Processor interface
type BasicProcessor struct{}

// NewBasicProcessor creates a new processor
func NewBasicProcessor() *BasicProcessor {
	return &BasicProcessor{}
}

// Process validates and normalizes a metric
func (p *BasicProcessor) Process(metric model.Metric) (model.Metric, error) {

	// Trace metric before processing
	util.PrintStage("INGESTION & VALIDATION (Before Processing)", metric)

	// Validation: Metric name must exist
	if metric.Name == "" {
		return metric, errors.New("validation failed: metric name is empty")
	}

	// Validation: Timestamp must exist
	if metric.Timestamp.IsZero() {
		metric.Timestamp = time.Now()
	}

	// Validation: Labels must exist
	if metric.Labels == nil {
		metric.Labels = make(map[string]string)
	}

	// Trace metric after validation & normalization
	util.PrintStage("AFTER PROCESSING (Validated & Normalized)", metric)

	return metric, nil
}
