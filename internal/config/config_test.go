package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() returned error: %v", err)
	}

	if cfg.EngagementRoot == "" {
		t.Error("Loaded config EngagementRoot should not be empty")
	}
	if cfg.TemplateDir == "" {
		t.Error("Loaded config TemplateDir should not be empty")
	}
}
