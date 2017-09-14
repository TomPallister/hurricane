package hurricane

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/fsnotify/fsnotify"
)

var features map[string]bool

type watchingFileFeatureProvider struct {
	path string
}

func (p *watchingFileFeatureProvider) makeFeatures() error {
	b, err := ioutil.ReadFile(p.path)
	if err != nil {
		return err
	}
	features = make(map[string]bool)
	err = json.Unmarshal(b, &features)
	if err != nil {
		return err
	}

	return nil
}

func (p *watchingFileFeatureProvider) start() error {
	p.makeFeatures()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case <-watcher.Events:
				// logger.Println("event:", event)
				// if event.Op&fsnotify.Write == fsnotify.Write {
				// 	logger.Println("modified file:", event.Name)
				// }
				p.makeFeatures()
			case err := <-watcher.Errors:
				logger.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(p.path)
	if err != nil {
		logger.Fatal(err)
	}
	<-done
	return nil
}

func (p *watchingFileFeatureProvider) Enabled(key string) (bool, error) {
	enabled, ok := features[key]
	if ok == false {
		message := fmt.Sprintf("key %v does not exist", key)
		logger.Printf(message)
		return false, nil
	}

	return enabled, nil
}
