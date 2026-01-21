package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// RelabelConfig defines relabeling rules
type RelabelConfig struct {
	AddLabels    map[string]string `yaml:"add_labels"`
	RenameLabels map[string]string `yaml:"rename_labels"`
}

// LoadRelabelConfig loads relabel config from YAML file
func LoadRelabelConfig(path string) (*RelabelConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg RelabelConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
