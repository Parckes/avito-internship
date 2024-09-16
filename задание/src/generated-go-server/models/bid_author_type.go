package openapi

import (
	"fmt"
)

type BidAuthorType string

const (
	ORGANIZATION BidAuthorType = "Organization"
	USER         BidAuthorType = "User"
)

var AllowedBidAuthorTypeEnumValues = []BidAuthorType{
	"Organization",
	"User",
}

var validBidAuthorTypeEnumValues = map[BidAuthorType]struct{}{
	"Organization": {},
	"User":         {},
}

func (v BidAuthorType) IsValid() bool {
	_, ok := validBidAuthorTypeEnumValues[v]
	return ok
}

func NewBidAuthorTypeFromValue(v string) (BidAuthorType, error) {
	ev := BidAuthorType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for BidAuthorType: valid values are %v", v, AllowedBidAuthorTypeEnumValues)
}

func AssertBidAuthorTypeRequired(obj BidAuthorType) error {
	return nil
}

func AssertBidAuthorTypeConstraints(obj BidAuthorType) error {
	return nil
}
