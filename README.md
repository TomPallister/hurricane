# hurricane

## Credit

fsnotify - go get github.com/fsnotify/fsnotify

## How to

Example set up below..

```go
package main

import (
	"log"
	"os"
	"testing"

	"github.com/TomPallister/hurricane"
)

func main(){
    // This will build a features that reads from the file in the given path
    f := hurricane.NewFileFeatures(path)
    // This will build a features that watches the file in the given path
	f := hurricane.NewWatchingFileFeatures(path)
    //This will build a features that lets you pass any provider
	p := FakeProvider{enabled: true}
	f := hurricane.NewFeatures(p)
    //get a feature status
	enabled := f.Enabled("my-feature")
}

```

The features file must use this data structure...

```json
{"featureName":false,"my-feature":false}
```

## Further reading
To understand hurricane fully please take a look at the test class and the code itself. It's not complex.

## Future
It would be nice to do something distributed rather than a file...