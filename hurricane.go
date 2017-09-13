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

type feature struct {
	key     string
	enabled bool
}

// FeatureProvider is an interface to the thing that actually finds out if a feature is on or not
type FeatureProvider interface {
	Enabled(key string) (bool, error)
}

type FileFeatureProvider struct {
	path string
}

func NewFileFeatureProvider(path string) *FileFeatureProvider {
	return &FileFeatureProvider{path: path}
}

func (p FileFeatureProvider) Enabled(key string) (bool, error) {
	b, err := ioutil.ReadFile(p.path)
	if err != nil {
		return false, err
	}
	//get the features..
	j, err := json.Unmarshal(b, []feature)
	//get our features...
	//return its value..
	return true, nil
}

// New creates a pointer to features it takes a provider and a logger
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
