package thezipcodes

type apiResponse interface {
	IsSuccessful() bool
}

type baseResponse struct {
	Success bool `json:"success"`
}

func (b baseResponse) IsSuccessful() bool {
	return b.Success
}
