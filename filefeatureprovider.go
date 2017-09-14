package hurricane

import (
	"encoding/json"
	"io/ioutil"
)

type fileFeatureProvider struct {
	path string
}

func (p *fileFeatureProvider) Enabled(key string) (bool, error) {
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
