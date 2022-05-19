package person

import (
	"encoding/json"
	"net/http"

	"github.com/snowmerak/workspace-demo/model/person"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	p := person.Person{
		Name: "Snowmerak",
		Age:  27,
	}
	write := json.NewEncoder(w)
	if err := write.Encode(p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}
