package index

import "net/http"

func Index(rw http.ResponseWriter, _ *http.Request) {
	rw.Write([]byte("hello world"))
	rw.WriteHeader(http.StatusOK)
}
