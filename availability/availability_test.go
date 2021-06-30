package availability

import (
	"errors"
	"net/http"
	"testing"

	"googlemaps.github.io/maps"
)

func TestAvailabilityResponse(t *testing.T) {

	response, statusCode := AvailabilityResponse()

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
func TestSuccessResponse(t *testing.T) {
	data := Data{
		Blocked:  true,
		Distance: 3800,
		Duration: 3000,
	}
	response, statusCode := SuccessResponse(data)

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
	name := "test: error."
	err := errors.New(name)
	response, statusCode := ErrorResponse(err)

	if statusCode != http.StatusServiceUnavailable {
		t.Errorf("Wrong statuscode: %d - expected: %d", statusCode, http.StatusServiceUnavailable)
	}

	if response.Success != false {
		t.Errorf("Wrong response attribute for success: %t - expected: %t", response.Success, false)
	}
	if response.Error != name {
		t.Errorf("Wrong error message thrown. Got: \"%s\", expected \"%s\".", response.Error, name)
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
		data := BuildData(&test.Response, test.Threshold)

		if data.Blocked != test.Blocked {
			t.Errorf("Blocked: %t - with Threshold %d, expected: %t", data.Blocked, test.Threshold, test.Blocked)
		}
	}
}
