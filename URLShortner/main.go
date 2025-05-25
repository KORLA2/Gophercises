package main

import (
	"fmt"
	"myapp/Handler"
	"net/http"
)

func main() {

	mux := defaultMux()

	navigateUrls := map[string]string{
		"/about": "https://leetcode.com",
		"/ping":  "https://github.com",
	}
	mapHandler := Handler.MapHandler(navigateUrls, mux)

	yaml := `
- path: "/ytb"
  url: "https://youtube.com"
- path: "/glg"
  url: "https://google.com"
`

	// 	json :=
	// 		`
	// [
	//     {
	//         "path": "/ytb",
	//         "url": "https://youtube.com"
	//     },
	//     {
	//         "path": "/glg",
	//         "url": "https://google.com"
	//     }
	// ]
	// `
	ymlhandler := Handler.YamlHandler([]byte(yaml), mapHandler)

	http.ListenAndServe(":9000", ymlhandler)

}

func defaultMux() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		fmt.Fprint(w, "Hello World from Go web server")
	})
	return mux
}
