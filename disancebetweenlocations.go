package thezipcodes

type getDistanceBetweenLocationsResponse struct {
	baseResponse

	Result DistanceBetweenLocations `json:"result"`
}

type DistanceBetweenLocations struct {
	FromZipCode string `json:"fromZipCode"`
	ToZipCode   string `json:"toZipCode"`
	Distance    string `json:"distance"`
	Unit        string `json:"unit"`
}
