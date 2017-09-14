package hurricane_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"testing"
	"time"

	"github.com/TomPallister/hurricane"
)

type FakeProvider struct {
	enabled bool
	err     error
}

func (p FakeProvider) Enabled(key string) (bool, error) {
	return p.enabled, p.err
}

func TestReturnsFalse(t *testing.T) {
	p := FakeProvider{enabled: false}
	f := hurricane.NewFeatures(p)
	enabled := f.Enabled("my-feature")
	if enabled == true {
		t.Fatalf("Should be false")
	}
}

func TestReturnsFalseBecauseOfError(t *testing.T) {
	p := FakeProvider{enabled: true, err: errors.New("some kind of error")}
	f := hurricane.NewFeatures(p)
	enabled := f.Enabled("my-feature")
	if enabled == true {
		t.Fatalf("Should be false")
	}
}

func TestReturnsTrue(t *testing.T) {
	p := FakeProvider{enabled: true}
	f := hurricane.NewFeatures(p)
	enabled := f.Enabled("my-feature")
	if enabled == false {
		t.Fatalf("Should be true")
	}
}

func TestFileFeatureProvider(t *testing.T) {
	path := "features.json"
	featureName := "my-feature"
	features := map[string]bool{featureName: true}
	b, _ := json.Marshal(features)
	_ = ioutil.WriteFile(path, b, 0644)
	f := hurricane.NewFileFeatures(path)
	enabled := f.Enabled(featureName)
	if enabled == false {
		t.Fatalf("Should be true")
	}
}

func TestWatchingFileFeatureProvider(t *testing.T) {
	path := "features.json"
	featureName := "my-feature"
	features := map[string]bool{featureName: false}
	b, _ := json.Marshal(features)
	_ = ioutil.WriteFile(path, b, 0644)
	f := hurricane.NewWatchingFileFeatures(path)

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
