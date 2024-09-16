package openapi

import (
	"fmt"
)

type TenderServiceType string

const (
	CONSTRUCTION TenderServiceType = "Construction"
	DELIVERY     TenderServiceType = "Delivery"
	MANUFACTURE  TenderServiceType = "Manufacture"
	FREE         TenderServiceType = ""
)

var AllowedTenderServiceTypeEnumValues = []TenderServiceType{
	"Construction",
	"Delivery",
	"Manufacture",
}

var validTenderServiceTypeEnumValues = map[TenderServiceType]struct{}{
	"Construction": {},
	"Delivery":     {},
	"Manufacture":  {},
}

func (v TenderServiceType) IsValid() bool {
	_, ok := validTenderServiceTypeEnumValues[v]
	return ok
}

func NewTenderServiceTypeFromValue(v string) (TenderServiceType, error) {
	ev := TenderServiceType(v)
	if ev.IsValid() {
		return ev, nil
	}
	return "", fmt.Errorf("invalid value '%v' for TenderServiceType: valid values are %v", v, AllowedTenderServiceTypeEnumValues)
}

func AssertTenderServiceTypeRequired(obj TenderServiceType) error {
	return nil
}

func AssertTenderServiceTypeConstraints(obj TenderServiceType) error {
	return nil
}
