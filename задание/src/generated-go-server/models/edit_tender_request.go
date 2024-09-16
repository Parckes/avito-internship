package openapi

type EditTenderRequest struct {
	Name        string             `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	ServiceType TenderServiceType  `json:"serviceType,omitempty"`
}

func AssertEditTenderRequestRequired(obj EditTenderRequest) error {
	return nil
}

func AssertEditTenderRequestConstraints(obj EditTenderRequest) error {
	return nil
}
