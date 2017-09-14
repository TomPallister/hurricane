package hurricane

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Features is holds the dependencies required to identify if features are on or not
type Features struct {
	provider FeatureProvider
	logger   *log.Logger
}

// FeatureProvider is an interface to the thing that actually finds out if a feature is on or not
type FeatureProvider interface {
	Enabled(key string) (bool, error)
}

type FileFeatureProvider struct {
	path string
}

// NewFileFeatureProvider is an interface to the thing that reads from a file containing json data for map[string]bool
func NewFileFeatureProvider(path string) *FileFeatureProvider {
	return &FileFeatureProvider{path: path}
}

// Enabled tries to get the feature from the provider path
func (p FileFeatureProvider) Enabled(key string) (bool, error) {
	b, err := ioutil.ReadFile(p.path)
	if err != nil {
		return false, err
	}
	features := make(map[string]bool)
	err = json.Unmarshal(b, &features)
	if err != nil {
		return false, err
	}
	feature, ok := features[key]
	if ok == false {
		return false, nil
	}
	return feature, nil
}

// NewFeatures creates a pointer to features it takes a provider and a logger
func NewFeatures(provider FeatureProvider, logger *log.Logger) *Features {
	features := Features{provider: provider, logger: logger}
	return &features
}

// Enabled is used to check if feature is enabled
func (features *Features) Enabled(key string) bool {
	enabled, err := features.provider.Enabled(key)
	if err == nil {
		return enabled
	}
	features.logger.Printf("Error getting value for key %v. Error is %v", key, err)
	return false
}
