package openapi

type BidReview struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}

func AssertBidReviewRequired(obj BidReview) error {
	elements := map[string]interface{}{
		"id":          obj.Id,
		"description": obj.Description,
		"createdAt":   obj.CreatedAt,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

func AssertBidReviewConstraints(obj BidReview) error {
	return nil
}
