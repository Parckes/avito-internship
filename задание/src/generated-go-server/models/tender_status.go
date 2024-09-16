package openapi

import (
	"fmt"
)

type TenderStatus string

const (
	CREATED   TenderStatus = "Created"
	PUBLISHED TenderStatus = "Published"
	CLOSED    TenderStatus = "Closed"
)

var AllowedTenderStatusEnumValues = []TenderStatus{
	"Created",
	"Published",
	"Closed",
}

var validTenderStatusEnumValues = map[TenderStatus]struct{}{
	"Created":   {},
	"Published": {},
	"Closed":    {},
}

func (v TenderStatus) IsValid() bool {
	_, ok := validTenderStatusEnumValues[v]
	return ok
}

func NewTenderStatusFromValue(v string) (TenderStatus, error) {
	ev := TenderStatus(v)
	if ev.IsValid() {
		return ev, nil
	}
	return "", fmt.Errorf("invalid value '%v' for TenderStatus: valid values are %v", v, AllowedTenderStatusEnumValues)
}

func AssertTenderStatusRequired(obj TenderStatus) error {
	return nil
}

func AssertTenderStatusConstraints(obj TenderStatus) error {
	return nil
}
