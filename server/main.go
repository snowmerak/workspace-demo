package main

import (
	"net/http"

	"github.com/snowmerak/workspace-demo/router/index"
	"github.com/snowmerak/workspace-demo/router/person"
	"github.com/snowmerak/workspace-demo/router/weather"
)

func main() {
	http.HandleFunc("/", index.Index)
	http.HandleFunc("/person", person.Index)
	http.HandleFunc("/weather", weather.Index)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
