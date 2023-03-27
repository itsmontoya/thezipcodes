package thezipcodes

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

const (
	host = "https://thezipcodes.com"
	api  = "/api/v1"

	getZipcodeAddressEndpoint           = "/search"
	getDistanceBetweenLocationsEndpoint = "/distance"
	getStatesByCountryEndpoint          = "/states"
)

func New(apiKey string) (ref *Client, err error) {
	if len(apiKey) == 0 {
		err = errors.New("invalid apiKey, cannot be empty")
		return
	}

	var c Client
	if c.host, err = url.Parse(host); err != nil {
		return
	}

	c.apiKey = apiKey
	c.r = makeRateLimiter()
	ref = &c
	return
}

type Client struct {
	r    rateLimiter
	hc   http.Client
	host *url.URL

	apiKey string
}

func (c *Client) GetZipcodeAddress(zipCode, countryCode string) (locations []Location, err error) {
	var resp getZipcodeAddressResponse
	p := url.Values{}
	p.Add("zipCode", zipCode)
	p.Add("countryCode", countryCode)
	url := c.getURL(getZipcodeAddressEndpoint, p)
	if err = c.do(url, &resp); err != nil {
		return
	}

	locations = resp.Locations
	return
}

func (c *Client) GetDistanceBetweenLocations(fromZipCode, toZipCode, countryCode, unit string) (result *DistanceBetweenLocations, err error) {
	var resp getDistanceBetweenLocationsResponse
	p := url.Values{}
	p.Add("fromZipCode", fromZipCode)
	p.Add("toZipCode", toZipCode)
	p.Add("unit", unit)
	p.Add("countryCode", countryCode)
	url := c.getURL(getDistanceBetweenLocationsEndpoint, p)
	if err = c.do(url, &resp); err != nil {
		return
	}

	result = &resp.Result
	return
}

func (c *Client) GetStatesByCountry(countryCode string) (states []string, err error) {
	var resp getStatesByCountryResponse
	p := url.Values{}
	p.Add("countryCode", countryCode)
	url := c.getURL(getStatesByCountryEndpoint, p)
	if err = c.do(url, &resp); err != nil {
		return
	}

	states = resp.Results.States
	return
}

func (c *Client) do(url string, out apiResponse) (err error) {
	var req *http.Request
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return
	}

	req.Header.Set("apikey", c.apiKey)

	// Wait for rate limiter
	c.r.Request()

	var resp *http.Response
	if resp, err = c.hc.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(out); err != nil {
		return
	}

	if !out.IsSuccessful() {
		err = errors.New("request was not successful")
		return
	}

	return
}

func (c *Client) getURL(path string, params url.Values) string {
	u := *c.host
	u.Path = api + path
	if params != nil {
		u.RawQuery = params.Encode()
	}

	return u.String()
}
