package openapi

type ErrorResponse struct {
	Reason string `json:"reason"`
}

func AssertErrorResponseRequired(obj ErrorResponse) error {
	elements := map[string]interface{}{
		"reason": obj.Reason,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}
	return nil
}

func AssertErrorResponseConstraints(obj ErrorResponse) error {
	return nil
}
