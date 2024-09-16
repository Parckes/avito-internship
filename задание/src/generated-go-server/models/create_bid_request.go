package openapi

import (
	"fmt"
	"github.com/google/uuid"
)

type CreateBidRequest struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	TenderId    string        `json:"tenderId"`
	AuthorType  BidAuthorType `json:"authorType"`
	AuthorId    string        `json:"authorId"`
}

func AssertCreateBidRequestRequired(obj CreateBidRequest) error {
	elements := map[string]interface{}{
		"name":        obj.Name,
		"description": obj.Description,
		"tenderId":    obj.TenderId,
		"authorType":  obj.AuthorType,
		"authorId":    obj.AuthorId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}
	return nil
}

func AssertCreateBidRequestConstraints(obj CreateBidRequest) error {
	if len(obj.Name) > 100 {
		return fmt.Errorf("name exceeds the maximum length of 100 characters")
	}
	if len(obj.Description) > 500 {
		return fmt.Errorf("description exceeds the maximum length of 500 characters")
	}
	if _, err := uuid.Parse(obj.TenderId); err != nil {
		return fmt.Errorf("invalid tenderId: must be a valid UUID")
	}
	validAuthorTypes := map[string]bool{
		"Organization": true,
		"User":         true,
	}
	if !validAuthorTypes[obj.AuthorType.String()] {
		return fmt.Errorf("invalid authorType: must be one of 'Organization' or 'User'")
	}
	if _, err := uuid.Parse(obj.AuthorId); err != nil {
		return fmt.Errorf("invalid authorId: must be a valid UUID")
	}
	return nil
}

func (c *BidAuthorType) String() string {
	return string(*c)
}
