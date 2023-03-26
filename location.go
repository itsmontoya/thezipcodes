package zipcodes

type getZipcodeAddressResponse struct {
	baseResponse

	Locations []Location `json:"location"`
}

type Location struct {
	ZipCode      string `json:"zipCode"`
	Country      string `json:"country"`
	CountryCode2 string `json:"countryCode2"`
	CountryCode3 string `json:"countryCode3"`
	State        string `json:"state"`
	StateCode2   string `json:"stateCode2"`
	City         string `json:"city"`
	County       string `json:"county"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	TimeZone     string `json:"timeZone"`
}
