package hurricane_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"testing"

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
	l := log.New(os.Stdout, "Log: ", log.Ldate|log.Ltime|log.Lshortfile)
	p := FakeProvider{enabled: false}
	f := hurricane.NewFeatures(p, l)
	enabled := f.Enabled("my-feature")
	if enabled == true {
		t.Fatalf("Should be false")
	}
}

func TestReturnsFalseBecauseOfError(t *testing.T) {
	l := log.New(os.Stdout, "Log: ", log.Ldate|log.Ltime|log.Lshortfile)
	p := FakeProvider{enabled: true, err: errors.New("some kind of error")}
	f := hurricane.NewFeatures(p, l)
	enabled := f.Enabled("my-feature")
	if enabled == true {
		t.Fatalf("Should be false")
	}
}

func TestReturnsTrue(t *testing.T) {
	l := log.New(os.Stdout, "Log: ", log.Ldate|log.Ltime|log.Lshortfile)
	p := FakeProvider{enabled: true}
	f := hurricane.NewFeatures(p, l)
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
	l := log.New(os.Stdout, "Log: ", log.Ldate|log.Ltime|log.Lshortfile)
	p := hurricane.NewFileFeatureProvider(path)
	f := hurricane.NewFeatures(p, l)
	enabled := f.Enabled(featureName)
	if enabled == false {
		t.Fatalf("Should be true")
	}
}
