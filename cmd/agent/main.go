package main

import (
	"fmt"
	"log"
	"time"

	"github.com/your-org/gtssp/internal/config"
	"github.com/your-org/gtssp/internal/exporter"
	"github.com/your-org/gtssp/internal/processor"
	"github.com/your-org/gtssp/internal/relabel"
	"github.com/your-org/gtssp/internal/scraper"
	"github.com/your-org/gtssp/internal/server"
	"github.com/your-org/gtssp/internal/store"
)

func main() {

	// 1️⃣ Start Output Web Server FIRST
	server.StartOutputServer()
	fmt.Println("🔗 Output available at: http://localhost:8082/output")

	// 2️⃣ Initialize pipeline components
	s := scraper.NewDummyScraper()
	p := processor.NewBasicProcessor()
	store := store.NewInMemoryMetricStore()

	cfg, err := config.LoadRelabelConfig("configs/relabel.yaml")
	if err != nil {
		log.Fatal(err)
	}

	r := relabel.NewSimpleRelabeler(cfg.AddLabels, cfg.RenameLabels)
	exp := exporter.NewPrometheusExporter()

	fmt.Println("🚀 GTSPP Agent started")

	// 3️⃣ Periodic scrape loop
	for {

		metrics, err := s.Scrape()
		if err != nil {
			log.Println("scrape error:", err)
			continue
		}

		// Clear ONLY the MetricStore (NOT tracer)
		store.Clear()

		for _, m := range metrics {

			processed, err := p.Process(m)
			if err != nil {
				log.Println("processing error:", err)
				continue
			}

			relabeled := r.Apply(processed)
			store.Add(relabeled)
		}

		// Export to Prometheus
		exp.Export(store.GetAll())

		fmt.Println("FINAL OUTPUT: Metric exposed to Prometheus /metrics")

		time.Sleep(5 * time.Second)
	}
}
