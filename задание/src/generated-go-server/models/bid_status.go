package openapi

import (
	"fmt"
)

type BidStatus string

const (
	CREATED_BID   BidStatus = "Created"
	PUBLISHED_BID BidStatus = "Published"
	CANCELED_BID  BidStatus = "Canceled"
)

var AllowedBidStatusEnumValues = []BidStatus{
	"Created",
	"Published",
	"Canceled",
}

var validBidStatusEnumValues = map[BidStatus]struct{}{
	"Created":   {},
	"Published": {},
	"Canceled":  {},
}

func (v BidStatus) IsValid() bool {
	_, ok := validBidStatusEnumValues[v]
	return ok
}

func NewBidStatusFromValue(v string) (BidStatus, error) {
	ev := BidStatus(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for BidStatus: valid values are %v", v, AllowedBidStatusEnumValues)
}

func AssertBidStatusRequired(obj BidStatus) error {
	return nil
}

func AssertBidStatusConstraints(obj BidStatus) error {
	return nil
}
