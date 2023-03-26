package zipcodes

type getStatesByCountryResponse struct {
	baseResponse

	Results StatesByCountry `json:"results"`
}

type StatesByCountry struct {
	Country string   `json:"country"`
	States  []string `json:"states"`
}
