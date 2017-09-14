package hurricane

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr, "Log: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Features is holds the dependencies required to identify if features are on or not
type Features struct {
	provider FeatureProvider
}

// FeatureProvider is an interface to the thing that actually finds out if a feature is on or not
type FeatureProvider interface {
	Enabled(key string) (bool, error)
}

// NewFeatures creates a pointer to features it takes a provider and a logger
func NewFeatures(provider FeatureProvider) *Features {
	features := Features{provider: provider}
	return &features
}

func NewFileFeatures(path string) *Features {
	provider := &fileFeatureProvider{path: path}
	features := Features{provider: provider}
	return &features
}

func NewWatchingFileFeatures(path string) *Features {
	provider := &watchingFileFeatureProvider{path: path}
	go provider.start()
	return &Features{provider: provider}
}

// Enabled is used to check if feature is enabled
func (features *Features) Enabled(key string) bool {
	enabled, err := features.provider.Enabled(key)
	if err == nil {
		return enabled
	}
	logger.Printf("Error getting value for key %v. Error is %v", key, err)
	return false
}
