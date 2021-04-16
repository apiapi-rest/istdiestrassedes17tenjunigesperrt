package availability

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

const (
	travelMode        = maps.TravelModeDriving
	travelOrigin      = "52.5162467446992,13.376336268075974"
	travelDestination = "52.51280515463461,13.323774503739166"
	threshold         = 3700
)

type Response struct {
	Data Data
}

type Data struct {
	Blocked  bool          `json:"blocked"`
	Distance int           `json:"distance"`
	Duration time.Duration `json:"duration"`
}

// func main() {
// 	data := fetch()

// 	pretty.Println(data)

// 	json, err := json.MarshalIndent(data, "", "	")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(json))

// }

func FetchDistance() Data {
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_API_KEY")))

	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.DistanceMatrixRequest{
		Mode:         travelMode,
		Origins:      []string{travelOrigin},
		Destinations: []string{travelDestination},
	}

	matrixResponse, err := c.DistanceMatrix(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	distance := matrixResponse.Rows[0].Elements[0].Distance.Meters

	response := Data{
		Blocked:  distance > threshold,
		Distance: distance,
		Duration: matrixResponse.Rows[0].Elements[0].Duration / 1000 / 1000 / 1000,
	}
	pretty.Println(matrixResponse)
	return response
}
