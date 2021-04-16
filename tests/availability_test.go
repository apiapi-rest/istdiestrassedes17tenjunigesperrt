package availability

import (
	"testing"

	"apiapi.rest/istdiestrassedes17tenjunigesperrt/availability"
	"googlemaps.github.io/maps"
)

func TestFetchDistance(t *testing.T) {
	response := maps.DistanceMatrixResponse{
		OriginAddresses:      []string{"B2 4, 10557 Berlin, Germany"},
		DestinationAddresses: []string{"Str. des 17. Juni 150, 10623 Berlin, Germany"},
		Rows: []maps.DistanceMatrixElementsRow{
			{
				Elements: []*maps.DistanceMatrixElement{
					{
						Status:            "OK",
						Duration:          390000000000,
						DurationInTraffic: 0,
						Distance:          maps.Distance{HumanReadable: "3.6 km", Meters: 3622},
					},
				},
			},
		},
	}

	tests := []struct {
		Response  maps.DistanceMatrixResponse
		Threshold int
		Blocked   bool
	}{
		{Response: response, Threshold: 3700, Blocked: false},
		{Response: response, Threshold: 1000, Blocked: true},
	}

	for _, test := range tests {
		data := availability.BuildData(&test.Response, test.Threshold)

		if data.Blocked != test.Blocked {
			t.Errorf("Blocked: %t - with Threshold %d, expected: %t", data.Blocked, test.Threshold, test.Blocked)
		}
	}
}
