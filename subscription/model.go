package subscription

type Subscription struct {
	ID        string  `json:"uuid"`
	ProductID string  `json:"productID"`
	PlanName  string  `json:"planName"`
	Duration  int     `json:"duration"`
	Price     float64 `json:"price"`
}
