package model

import "time"

// Metric represents a single time-series data point
// This is the core data contract shared across the system
type Metric struct {
	// Name of the metric (example: cpu_usage_seconds_total)
	Name string

	// Value of the metric
	Value float64

	// Timestamp when the metric was recorded
	Timestamp time.Time

	// Labels provide metadata for the metric
	// example: instance, job, env, region
	Labels map[string]string
}
