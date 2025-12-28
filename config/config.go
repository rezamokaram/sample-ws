package config

import (
	"encoding/json"
	"fmt"
	"log"
)

type SampleAuthConfig struct {
	Server ServerConfig `json:"app" yaml:"app"`
}

func (SampleAuthConfig) configSignature() {}

func (cfg SampleAuthConfig) Print() {
	jsonData, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal struct to JSON: %v", err)
	}

	fmt.Printf("loaded config: %v", string(jsonData))
}
