package pkg

type Availability struct {
	DomainName    string  `json:"domainName"`
	Purchasable   bool    `json:"purchasable"`
	PurchasePrice float64 `json:"purchasePrice"`
	RenewalPrice  float64 `json:"renewalPrice,omitempty"`
	PurchaseType  string  `json:"purchaseType,omitempty"`
}

type CheckAvailabilityRequest struct {
	DomainNames []string `json:"domainNames"`
}

type CheckAvailabilityResponse struct {
	Results []Availability `json:"results"`
}
