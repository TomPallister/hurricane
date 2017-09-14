package hurricane_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/TomPallister/hurricane"
)

func TestWatchingFileFeatureProvider(t *testing.T) {
	logger := log.New(os.Stderr, "Log: ", log.Ldate|log.Ltime|log.Lshortfile)
	path := "features.json"
	featureName := "my-feature"
	features := map[string]bool{featureName: false}
	b, _ := json.Marshal(features)
	_ = ioutil.WriteFile(path, b, 0644)
	f := hurricane.NewWatchingFileFeatures(path, logger)
	loops := 5
	count := 0
	passed := false
	for {
		if passed || count >= loops {
			break
		}
		features = map[string]bool{featureName: true}
		b, _ = json.Marshal(features)
		_ = ioutil.WriteFile(path, b, 0644)
		time.Sleep(time.Second)
		enabled := f.Enabled(featureName)
		if enabled == true {
			passed = true
		}
		count++
	}

	if passed == false {
		t.Fatalf("Should be true")
	}
}
