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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Max-Age", "3600")

	fmt.Fprint(w, string(json))
}
