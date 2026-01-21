package util

import (
	"fmt"
	"sync"
	"time"

	"github.com/your-org/gtssp/internal/model"
)

type TraceEntry struct {
	Stage       string
	Metric      model.Metric
	Summary     string
	AddedLabels map[string]string
	RemovedKeys []string
	NoOp        bool
	Timestamp   time.Time
}

var (
	traceMu sync.Mutex
	traces  []TraceEntry
	last    *model.Metric
)

// PrintStage stores metric lifecycle trace with diff detection
func PrintStage(stage string, metric model.Metric) {
	traceMu.Lock()
	defer traceMu.Unlock()

	entry := TraceEntry{
		Stage:       stage,
		Metric:      metric,
		AddedLabels: make(map[string]string),
		Timestamp:   time.Now(),
	}

	// Diff detection
	if last != nil {
		// Detect added / changed labels
		for k, v := range metric.Labels {
			if oldV, ok := last.Labels[k]; !ok || oldV != v {
				entry.AddedLabels[k] = v
			}
		}
		// Detect removed labels
		for k := range last.Labels {
			if _, ok := metric.Labels[k]; !ok {
				entry.RemovedKeys = append(entry.RemovedKeys, k)
			}
		}
	}

	if len(entry.AddedLabels) == 0 && len(entry.RemovedKeys) == 0 {
		entry.NoOp = true
		entry.Summary = "No changes applied at this stage."
	} else {
		entry.Summary = "Metric transformed at this stage."
	}

	// Save trace
	traces = append(traces, entry)

	// Update last snapshot
	copyMetric := metric
	last = &copyMetric

	// Console trace (for debugging)
	fmt.Println("--------------------------------------------------")
	fmt.Println("STAGE:", stage)
	fmt.Printf("Name      : %s\n", metric.Name)
	fmt.Printf("Value     : %v\n", metric.Value)
	fmt.Printf("Timestamp : %v\n", metric.Timestamp)
	fmt.Printf("Labels    : %+v\n", metric.Labels)
	fmt.Println("--------------------------------------------------")
}

// GetTraces returns stored traces
func GetTraces() []TraceEntry {
	traceMu.Lock()
	defer traceMu.Unlock()

	result := make([]TraceEntry, len(traces))
	copy(result, traces)
	return result
}

// ClearTraces removes all stored trace entries
func ClearTraces() {
	traceMu.Lock()
	defer traceMu.Unlock()

	traces = []TraceEntry{}
	last = nil
}
