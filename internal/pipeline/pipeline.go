package pipeline

import (
	"github.com/your-org/gtssp/internal/exporter"
	"github.com/your-org/gtssp/internal/processor"
	"github.com/your-org/gtssp/internal/relabel"
	"github.com/your-org/gtssp/internal/scraper"
	"github.com/your-org/gtssp/internal/store"
)

// Pipeline wires all components together
type Pipeline struct {
	Scraper   scraper.Scraper
	Processor processor.Processor
	Store     store.MetricStore
	Relabeler relabel.Relabeler
	Exporter  exporter.Exporter
}
