package istdiestrassedes17tenjunigesperrt

import (
	"encoding/json"
	"fmt"
	"net/http"

	"apiapi.rest/istdiestrassedes17tenjunigesperrt/availability"
)

func Availability(w http.ResponseWriter, r *http.Request) {
	data := availability.FetchDistance()

	// pretty.Println(data)

	json, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, string(json))
}
