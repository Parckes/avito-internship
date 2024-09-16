package openapi

type EditBidRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func AssertEditBidRequestRequired(obj EditBidRequest) error {
	return nil
}

func AssertEditBidRequestConstraints(obj EditBidRequest) error {
	return nil
}
