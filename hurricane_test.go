package hurricane_test

import (
	"errors"
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
