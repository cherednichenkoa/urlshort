package urlshort

import (
	"gopkg.in/yaml.v2"
	"net/http"
)


// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
			val, exists := pathsToUrls[req.URL.Path]
			if exists {
				http.Redirect(w, req, val, http.StatusMovedPermanently)
			} else {
				fallback.ServeHTTP(w,req)
			}
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

	values, error := unpackYml(yml)
	if error != nil {
		return nil, error
	}
	handler := func(w http.ResponseWriter, req *http.Request) {
		for _, url := range values {
			if url.Url == req.URL.Path {
				http.Redirect(w, req, url.Path, http.StatusMovedPermanently)
				return
			}
		}
		fallback.ServeHTTP(w,req)
	}
	return handler, nil
}

func unpackYml(yml []byte)([]UrlMappingYaml, error) {
	var out []UrlMappingYaml
	if err := yaml.Unmarshal(yml,&out); err != nil {
		return nil, err
	}
	return out, nil
}

type UrlMappingYaml struct  {
	Path string `yaml:"url"`
	Url string  `yaml:"path"`
}