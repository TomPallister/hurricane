package hurricane_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/TomPallister/hurricane"
)

func TestFileFeatureProvider(t *testing.T) {
	logger := log.New(os.Stderr, "Log: ", log.Ldate|log.Ltime|log.Lshortfile)
	path := "features.json"
	featureName := "my-feature"
	features := map[string]bool{featureName: true}
	b, _ := json.Marshal(features)
	_ = ioutil.WriteFile(path, b, 0644)
	f := hurricane.NewFileFeatures(path, logger)
	enabled := f.Enabled(featureName)
	if enabled == false {
		t.Fatalf("Should be true")
	}
}
