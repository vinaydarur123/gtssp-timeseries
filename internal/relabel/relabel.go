package relabel

import (
	"github.com/your-org/gtssp/internal/model"
	"github.com/your-org/gtssp/internal/util"
)

// Relabeler defines label transformation behavior
type Relabeler interface {
	Apply(metric model.Metric) model.Metric
}

// SimpleRelabeler applies config-based relabeling
type SimpleRelabeler struct {
	addLabels    map[string]string
	renameLabels map[string]string
}

// NewSimpleRelabeler creates a new relabeler
func NewSimpleRelabeler(add, rename map[string]string) *SimpleRelabeler {
	return &SimpleRelabeler{
		addLabels:    add,
		renameLabels: rename,
	}
}

// Apply applies relabeling rules to a metric
func (r *SimpleRelabeler) Apply(metric model.Metric) model.Metric {

	// Trace metric before relabeling
	util.PrintStage("BEFORE RELABELING (Validated Metric)", metric)

	if metric.Labels == nil {
		metric.Labels = make(map[string]string)
	}

	// Rename labels
	for oldKey, newKey := range r.renameLabels {
		if val, ok := metric.Labels[oldKey]; ok {
			delete(metric.Labels, oldKey)
			metric.Labels[newKey] = val
		}
	}

	// Add labels
	for k, v := range r.addLabels {
		metric.Labels[k] = v
	}

	// Trace metric after relabeling
	util.PrintStage("AFTER RELABELING (Config Applied)", metric)

	return metric
}
