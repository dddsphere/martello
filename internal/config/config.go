package config

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Config struct {
	namespace string // i.e.: MWZ, MYCVS, APP, etc...
	values    map[string]string
	defaults  map[string]string
	reload    bool
	location  *time.Location
}

// Load configuration.
func Load(namespace string) *Config {
	cfg := &Config{}
	cfg.SetNamespace(namespace)
	cfg.loadNamespaceEnvVars()
	cfg.reload = true
	cfg.location = time.Local
	cfg.EnableDefaults()
	return cfg
}

// SetNamespace for configuration.
func (cfg *Config) SetNamespace(namespace string) {
	cfg.namespace = strings.ToUpper(namespace)
}

func (cfg *Config) namespacePrefix() string {
	return fmt.Sprintf("%s_", cfg.namespace)
}

func (cfg *Config) SetValues(values map[string]string) {
	cfg.values = values
}

func (cfg *Config) SetDefaults(values map[string]string) {
	cfg.defaults = values
}

func (cfg *Config) SetTimeLocation(tl *time.Location) {
	cfg.location = tl
}

// EnableDefaults for well known configuration values.
// Most likely this method will be deprecated in the near future and the application that uses this package will become
// responsible for setting its default values.
func (cfg *Config) EnableDefaults() {
	cfg.defaults = map[string]string{
		"http.server.host":                  "0.0.0.0",
		"http.server.port":                  "8080",
		"http.server.shutdown.timeout.secs": "12",
	}
}

// Get reads all visible environment variables
// that belongs to the namespace.
// An optional reload parameter lets re-read
// the values from environment.
func (cfg *Config) Get(reload ...bool) map[string]string {
	if len(reload) > 1 && reload[0] {
		return cfg.get(true)
	}
	return cfg.get(false)
}

func (cfg *Config) get(reload bool) map[string]string {
	if reload || len(cfg.values) == 0 {
		return cfg.readNamespaceEnvVars()
	}
	return cfg.values
}

// Val read a specific namespaced  environment variable
// And return its value.
// An optional reload parameter lets re-read
// the values from environment.
func (cfg *Config) Val(key string, reload ...bool) (value string, ok bool) {
	vals := cfg.get(false)
	if len(reload) > 1 && reload[0] {
		vals = cfg.get(true)
	}
	val, ok := vals[key]
	return val, ok
}

// ValOrDef read a specific namespaced environment variables and return its value.
// A default value is returned if key value is not found.
// An optional reload parameter lets re-read the value from environment.
func (cfg *Config) ValOrDef(key string, defVal string, reload ...bool) (value string) {
	vals := cfg.get(false)
	if len(reload) > 1 && reload[0] {
		vals = cfg.get(true)
	}
	val, ok := vals[key]
	if !ok {
		val = defVal
	}
	return val
}

// loadNamespaceEnvVars load all visible environment variables that belongs to the namespace.
func (cfg *Config) loadNamespaceEnvVars() {
	cfg.values = cfg.readNamespaceEnvVars()
}

// readNamespaceEnvVars reads all visible environment variables that belongs to the namespace.
func (cfg *Config) readNamespaceEnvVars() map[string]string {
	nevs := make(map[string]string)
	np := cfg.namespacePrefix()

	for _, ev := range os.Environ() {
		if strings.HasPrefix(ev, np) {
			varval := strings.SplitN(ev, "=", 2)

			if len(varval) < 2 {
				continue
			}

			key := cfg.keyify(varval[0])
			nevs[key] = varval[1]
		}
	}

	return nevs
}

// keyify environment variable names
// i.e.: NAMESPACE_CONFIG_VALUE becomes config.value
func (cfg *Config) keyify(name string) string {
	split := strings.Split(name, "_")
	if len(split) < 1 {
		return ""
	}
	// Without namespace prefix
	wnsp := strings.Join(split[1:], ".")
	// Dot separated lowercased
	dots := strings.ToLower(strings.Replace(wnsp, "_", ".", -1))
	return dots
}

// getEnvOrDef returns the value of environment variable or a default value if this value is empty or an empty string.
// An empty string is returned if environment variable is nonexistent and a default was not provided.
func getEnvOrDef(envar string, def ...string) string {
	val := os.Getenv(envar)
	if val != "" {
		return val
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}
