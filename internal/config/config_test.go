package config_test

import (
	"os"
	"testing"

	"github.com/dddsphere/martello/internal/config"
)

// TestGet default (no reload)
func TestGet(t *testing.T) {
	// Set some namespaced environment variables.
	err := os.Setenv("CFG_SERVER_HOST", "localhost")
	if err != nil {
		t.Errorf("Test setup error: %s", err.Error())
	}
	err = os.Setenv("CFG_SERVER_PORT", "8080")
	if err != nil {
		t.Errorf("Test setup error: %s", err.Error())
	}

	// Init a config handler.
	cfg := config.Config{}
	cfg.SetNamespace("CFG")

	// Read and process environment variables.
	res := cfg.Get()

	// Assert result.
	host, ok1 := res["server.host"]
	port, ok2 := res["server.port"]

	if !ok1 || !ok2 || host != "localhost" || port != "8080" {
		t.Error("error processing config environment variables")
	}
}

// TestGet no reload
func TestGetNoReload(t *testing.T) {
	// Set some namespaced environment variables.
	err := os.Setenv("CFG_SERVER_HOST", "localhost")
	if err != nil {
		t.Errorf("Test setup error: %s", err.Error())
	}
	err = os.Setenv("CFG_SERVER_PORT", "8080")
	if err != nil {
		t.Errorf("Test setup error: %s", err.Error())
	}

	// Init a config handler.
	cfg := config.Config{}
	cfg.SetNamespace("CFG")

	// Read and process environment variables.
	res := cfg.Get(false)

	// Assert result.
	host, ok1 := res["server.host"]
	port, ok2 := res["server.port"]

	if !ok1 || !ok2 || host != "localhost" || port != "8080" {
		t.Error("error processing config environment variables")
	}
}

// TestGet reload
func TestGetReload(t *testing.T) {
	// Set some namespaced environment variables.
	err := os.Setenv("CFG_SERVER_HOST", "localhost")
	if err != nil {
		t.Errorf("Test setup error: %s", err.Error())
	}
	err = os.Setenv("CFG_SERVER_PORT", "8080")
	if err != nil {
		t.Errorf("Test setup error: %s", err.Error())
	}

	// Init a config handler.
	cfg := config.Config{}
	cfg.SetNamespace("CFG")

	// Read and process environment variables.
	res := cfg.Get(true)

	// Assert result.
	host, ok1 := res["server.host"]
	port, ok2 := res["server.port"]

	if !ok1 || !ok2 || host != "localhost" || port != "8080" {
		t.Error("error processing config environment variables")
	}
}

func TestLoadEnvVars(t *testing.T) {
	// Set some namespaced environment variables.
	err := os.Setenv("CFG_SERVER_HOST", "localhost")
	if err != nil {
		t.Errorf("Test setup error: %s", err.Error())
	}
	err = os.Setenv("CFG_SERVER_PORT", "8080")
	if err != nil {
		t.Errorf("Test setup error: %s", err.Error())
	}

	// Init a config handler.
	cfg := config.Config{}
	cfg.SetNamespace("CFG")

	// Read and process environment variables.
	res := cfg.ReadNamespaceEnvVars()

	// Assert result.
	host, ok1 := res["server.host"]
	port, ok2 := res["server.port"]

	if !ok1 || !ok2 || host != "localhost" || port != "8080" {
		t.Error("error processing config environment variables")
	}
}
