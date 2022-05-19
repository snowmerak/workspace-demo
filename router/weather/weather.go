package weather

import (
	"encoding/json"
	"net/http"

	"github.com/snowmerak/workspace-demo/model/weather"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	writer := json.NewEncoder(w)

	if err := writer.Encode(weather.Weather{City: "Seoul", Temp: "22"}); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}
