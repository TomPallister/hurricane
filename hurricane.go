package hurricane

import (
	"log"
)

func init() {
	//logger = log.New(os.Stderr, "Log: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Features is holds the dependencies required to identify if features are on or not
type Features struct {
	provider FeatureProvider
	logger   *log.Logger
}

// FeatureProvider is an interface to the thing that actually finds out if a feature is on or not.
// This public to allow people to implement their own if they want
type FeatureProvider interface {
	Enabled(key string) (bool, error)
}

// NewFeatures creates a pointer to features it takes a given FeatureProvider
func NewFeatures(provider FeatureProvider, logger *log.Logger) *Features {
	features := Features{provider: provider, logger: logger}
	return &features
}

// NewFileFeatures creates a features that reads the feature from the file at the path location
// This file must be in the json structure {"featureName":false,"my-feature":false} or it will
// not work
func NewFileFeatures(path string, logger *log.Logger) *Features {
	provider := &fileFeatureProvider{path: path}
	features := Features{provider: provider}
	return &features
}

// NewWatchingFileFeatures creates a features that watches the feature from the file at the path location
// This file must be in the json structure {"featureName":false,"my-feature":false} or it will
// not work.
func NewWatchingFileFeatures(path string, logger *log.Logger) *Features {
	provider := &watchingFileFeatureProvider{path: path, logger: logger}
	go provider.start()
	return &Features{provider: provider}
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
