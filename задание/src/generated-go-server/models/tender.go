package openapi

import (
	"errors"
)

type Tender struct {
	Id             string             `json:"id"`
	Name           string             `json:"name"`
	Description    string             `json:"description"`
	Status         TenderStatus       `json:"status"`
	ServiceType    TenderServiceType   `json:"serviceType"`
	OrganizationId string             `json:"organizationId,omitempty"`
	Version        int32              `json:"version"`
	CreatedAt      string             `json:"createdAt"`
}

func AssertTenderRequired(obj Tender) error {
	elements := map[string]interface{}{
		"id":             obj.Id,
		"name":           obj.Name,
		"description":    obj.Description,
		"serviceType":    obj.ServiceType,
		"status":         obj.Status,
		"organizationId": obj.OrganizationId,
		"version":        obj.Version,
		"createdAt":      obj.CreatedAt,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}
	return nil
}

func AssertTenderConstraints(obj Tender) error {
	if obj.Version < 1 {
		return &ParsingError{Param: "Version", Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
