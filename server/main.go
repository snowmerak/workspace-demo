package main

import (
	"net/http"

	"github.com/snowmerak/workspace-demo/router/index"
)

func main() {
	http.HandleFunc("/", index.Index)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
