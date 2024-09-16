package openapi

import (
	"errors"
)

type Bid struct {
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description,omitempty"`
	Status      BidStatus     `json:"status"`
	TenderId    string        `json:"tenderId,omitempty"`
	AuthorType  BidAuthorType `json:"authorType"`
	AuthorId    string        `json:"authorId"`
	Version     int32         `json:"version"`
	CreatedAt   string        `json:"createdAt"`
}

func AssertBidRequired(obj Bid) error {
	elements := map[string]interface{}{
		"id":          obj.Id,
		"name":        obj.Name,
		"description": obj.Description,
		"status":      obj.Status,
		"tenderId":    obj.TenderId,
		"authorType":  obj.AuthorType,
		"authorId":    obj.AuthorId,
		"version":     obj.Version,
		"createdAt":   obj.CreatedAt,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

func AssertBidConstraints(obj Bid) error {
	if obj.Version < 1 {
		return &ParsingError{Param: "Version", Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
