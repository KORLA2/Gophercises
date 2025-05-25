package Handler

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v3"
)

func MapHandler(urls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if dest, ok := urls[r.URL.Path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		// fallback.ServeHTTP(w, r)

	}

}

type PathUrl struct {
	Path string `yaml: "path"`
	URL  string `yaml: "url"`
}

func YamlHandler(data []byte, fallback http.Handler) http.HandlerFunc {

	PathUrls := []PathUrl{}

	err := yaml.Unmarshal(data, &PathUrls)

	fmt.Println(err)
	pathToUrls := make(map[string]string)

	for _, val := range PathUrls {

		pathToUrls[val.Path] = val.URL
	}
	fmt.Println(string(data))
	fmt.Println(PathUrls)

	return MapHandler(pathToUrls, fallback)

}
