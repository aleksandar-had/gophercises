package urlshort

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// if path is matched
		// redirect to the path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		// else
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	decodedYaml, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}
	pathsToUrls := buildMap(decodedYaml)
	// Return a map handler using the map
	return MapHandler(pathsToUrls, fallback), nil
}

// Convert yaml array to map
func buildMap(data []pathUrl) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range data {
		pathsToUrls[pu.Path] = pu.URL
	}
	return pathsToUrls
}

// Parse the yaml
func parseYaml(data []byte) ([]pathUrl, error) {
	var decodedYaml []pathUrl
	err := yaml.Unmarshal(data, &decodedYaml)
	if err != nil {
		return nil, err
	}

	return decodedYaml, nil
}
