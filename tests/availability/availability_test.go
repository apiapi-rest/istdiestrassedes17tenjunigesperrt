package availability

import (
	"net/http"
	"os"
	"testing"

	"apiapi.rest/istdiestrassedes17tenjunigesperrt/availability"
	"googlemaps.github.io/maps"
)

func TestSuccssResponse(t *testing.T) {

	response, statusCode := availability.AvailabilityResponse()

	if statusCode != http.StatusOK {
		t.Errorf("Wrong statuscode: %d - expected: %d", statusCode, http.StatusOK)
	}

	if response.Success != true {
		t.Errorf("Wrong response attribute for success: %t - expected: %t", response.Success, true)
	}
	if len(response.Error) > 0 {
		t.Errorf("Error message thrown. Expected a string with length == 0. Erorr Message says: %s", response.Error)
	}
}
func TestErrorResponse(t *testing.T) {
	original := os.Getenv("GOOGLE_API_KEY")
	os.Setenv("GOOGLE_API_KEY", "")
	defer func() {
		os.Setenv("GOOGLE_API_KEY", original)
	}()

	response, statusCode := availability.AvailabilityResponse()

	if statusCode != http.StatusServiceUnavailable {
		t.Errorf("Wrong statuscode: %d - expected: %d", statusCode, http.StatusServiceUnavailable)
	}

	if response.Success != false {
		t.Errorf("Wrong response attribute for success: %t - expected: %t", response.Success, false)
	}
	if len(response.Error) == 0 {
		t.Errorf("No error message thrown. Expected a string with length > 0.")
	}
}

func TestBuildData(t *testing.T) {
	testResponse := maps.DistanceMatrixResponse{
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
		{Response: testResponse, Threshold: 3700, Blocked: false},
		{Response: testResponse, Threshold: 1000, Blocked: true},
	}

	for _, test := range tests {
		data := availability.BuildData(&test.Response, test.Threshold)

		if data.Blocked != test.Blocked {
			t.Errorf("Blocked: %t - with Threshold %d, expected: %t", data.Blocked, test.Threshold, test.Blocked)
		}
	}
}
