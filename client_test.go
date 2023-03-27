package thezipcodes

import (
	"log"
	"os"
	"reflect"
	"testing"
)

var testAPIKey string

func TestMain(m *testing.M) {
	testAPIKey = os.Getenv("THEZIPCODES_API_KEY")
	if len(testAPIKey) == 0 {
		log.Fatal("no api key set at the OS environment key of <THEZIPCODES_API_KEY>")
	}

	status := m.Run()
	os.Exit(status)
}

func TestClient_GetZipcodeAddress(t *testing.T) {
	type args struct {
		zipCode     string
		countryCode string
	}

	tests := []struct {
		name          string
		args          args
		wantLocations []Location
		wantErr       bool
	}{
		{
			name: "Portland",
			args: args{
				zipCode:     "97227",
				countryCode: "US",
			},
			wantLocations: []Location{
				{
					Country:      "US",
					CountryCode2: "US",
					CountryCode3: "USA",
					State:        "Oregon",
					StateCode2:   "OR",
					Latitude:     "45.5496",
					Longitude:    "-122.6743",
					ZipCode:      "97227",
					City:         "Portland",
				},
			},
			wantErr: false,
		},

		{
			name: "Australia",
			args: args{
				zipCode:     "2000",
				countryCode: "AU",
			},
			wantLocations: []Location{
				{
					Country:      "Australia",
					CountryCode2: "AU",
					CountryCode3: "AUS",
					State:        "New South Wales",
					StateCode2:   "NSW",
					Latitude:     "-33.8641",
					Longitude:    "151.2017",
					ZipCode:      "2000",
					City:         "Barangaroo",
					TimeZone:     "Australian Central Time",
				},
			},
			wantErr: false,
		},
	}

	c, err := New(testAPIKey)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLocations, err := c.GetZipcodeAddress(tt.args.zipCode, tt.args.countryCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotLocations, tt.wantLocations) {
				t.Errorf("Client.Get() = %+v, want %+v", gotLocations, tt.wantLocations)
			}
		})
	}
}
