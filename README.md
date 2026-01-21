# GTSPP – Time-Series Processing Library

GTSPP is a Go library for collecting, validating, transforming, and exporting time-series metrics in distributed and cloud-native environments.

The library is designed to be embedded within agents or services that require a structured and configurable metrics processing pipeline before forwarding data to Prometheus-compatible time-series databases.

---

## What This Library Provides

- Metric collection from application and system sources
- Validation and normalization of time-series data
- Configuration-driven label relabeling
- Temporary in-memory storage for processed metrics
- Prometheus exposition and remote write compatibility
- Metric lifecycle visibility for debugging and analysis

---

## Processing Flow

Metrics pass through the following stages:

1. **Scraper** – Collects raw metrics  
2. **Validation** – Ensures correctness and consistency  
3. **Processor** – Normalizes metric fields  
4. **Relabeling** – Applies label transformations via configuration  
5. **Metric Store** – Temporarily caches metrics  
6. **Exporter** – Exposes metrics to downstream systems  

Each stage is modular and can be extended or replaced independently.

---

## Configuration

Relabeling rules are defined using YAML configuration files, allowing metric transformation behavior to be modified without changing application code.

---

## Intended Usage

- Observability agents
- Metrics ingestion pipelines
- Cloud-native monitoring platforms
- Internal telemetry standardization systems

---

## Technology Stack

- Go
- Prometheus client libraries
- YAML-based configuration
- Cloud-native runtime environments

  ## Getting Started (Run From Scratch)

> The following steps demonstrate how to run the GTSPP agent locally and view live metric output.

```bash
# 1. Install Go (1.20 or higher)
https://go.dev/dl/

# 2. Clone the repository
git clone https://github.com/<your-org>/gtssp.git
cd gtssp

# 3. Initialize dependencies
go mod tidy

# 4. Configure relabel rules
# Edit configs/relabel.yaml if required

# 5. Run the agent
go run cmd/agent/main.go

# 6. Open output dashboard
http://localhost:8082/output

# 7. (Optional) View Prometheus metrics
http://localhost:8082/metrics

