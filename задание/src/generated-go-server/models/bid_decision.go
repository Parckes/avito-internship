package openapi

import (
	"fmt"
)

type BidDecision string

const (
	APPROVED BidDecision = "Approved"
	REJECTED BidDecision = "Rejected"
)

var AllowedBidDecisionEnumValues = []BidDecision{
	"Approved",
	"Rejected",
}

var validBidDecisionEnumValues = map[BidDecision]struct{}{
	"Approved": {},
	"Rejected": {},
}

func (v BidDecision) IsValid() bool {
	_, ok := validBidDecisionEnumValues[v]
	return ok
}

func NewBidDecisionFromValue(v string) (BidDecision, error) {
	ev := BidDecision(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for BidDecision: valid values are %v", v, AllowedBidDecisionEnumValues)
}

func AssertBidDecisionRequired(obj BidDecision) error {
	return nil
}

func AssertBidDecisionConstraints(obj BidDecision) error {
	return nil
}
