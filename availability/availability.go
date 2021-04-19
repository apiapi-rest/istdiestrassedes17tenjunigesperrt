package availability

import (
	"context"
	"net/http"
	"os"
	"time"

	"googlemaps.github.io/maps"
)

const (
	travelMode        = maps.TravelModeDriving
	travelOrigin      = "52.5162467446992,13.376336268075974"
	travelDestination = "52.51280515463461,13.323774503739166"
	threshold         = 3700
)

type Response struct {
	Data    Data   `json:"data"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type Data struct {
	Blocked  bool          `json:"blocked"`
	Distance int           `json:"distance"`
	Duration time.Duration `json:"duration"`
}

func AvailabilityResponse() (Response, int) {
	matrixResponse, err := FetchDistance()
	if err != nil {
		return Response{
			Data:    Data{},
			Error:   err.Error(),
			Success: false,
		}, http.StatusServiceUnavailable
	}
	data := BuildData(matrixResponse, threshold)

	return Response{
		Data:    data,
		Error:   "",
		Success: true,
	}, http.StatusOK
}

func FetchDistance() (*maps.DistanceMatrixResponse, error) {
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_API_KEY")))

	if err != nil {
		return nil, err
	}

	r := &maps.DistanceMatrixRequest{
		Mode:         travelMode,
		Origins:      []string{travelOrigin},
		Destinations: []string{travelDestination},
	}

	matrixResponse, err := c.DistanceMatrix(context.Background(), r)
	if err != nil {
		return nil, err
	}
	return matrixResponse, nil

}

func BuildData(matrixResponse *maps.DistanceMatrixResponse, currentThreshold int) Data {
	distance := matrixResponse.Rows[0].Elements[0].Distance.Meters

	data := Data{
		Blocked:  distance > currentThreshold,
		Distance: distance,
		Duration: matrixResponse.Rows[0].Elements[0].Duration / 1000 / 1000 / 1000,
	}
	return data
}
