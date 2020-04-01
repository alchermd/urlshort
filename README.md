# urlshort
Package that handles HTTP redirects from various file formats

## Usage

Included with this package is a `MapHandler`, `YAMLHandler`, and a `JSONHandler`. These handlers accept their respective data formats and a fallback handler, and then returns an `http.HandlerFunc` (and an error, if any). These handlers can be composed, allowing multiple data stores to be used for the URL shortener:

```golang
import "github.com/alchermd/urlshort"

pathsToUrls := map[string]string{
	"/aws":   "https://aws.amazon.com/",
	"/gcp":   "https://cloud.google.com/",
	"/azure": "https://azure.microsoft.com/",
}

// Build handler using a Map (with a defaultHandler implemented somewhere.)
mapHandler := urlshort.MapHandler(pathsToUrls, defaultHandler)

// The mapHandler can then be used as the fallback for the yamlHandler
yaml := `
- path: /foo
  url: https://example.com/
`
yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
if err != nil {
	panic(err)
}

// Same thing with the jsonHandler
json := `[{"path": "/ms", "url": "https://microsoft.com/"}, {"path": "/ub", "url": "https://ubuntu.com/"}]`
jsonHandler, err := urlshort.JSONHandler([]byte(json), yamlHandler)
if err != nil {
	panic(err)
}

// data not found on jsonHandler will then be passed onto the previous handlers.
```

See [example/main.go](/example/main.go) for example usage.

## LICENSE

See [LICENSE](LICENSE)