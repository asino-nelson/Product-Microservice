package product

import "time"

type Product struct {
	ID          string `json:"uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt   time.Duration `json:"createdAt"`
	UpdatedAt   time.Duration `json:"updatedAt"`
}

type DigitalProduct struct {
	Product
	FileSize int `json:"fileSize"`
	DownloadLink string `json:"downloadLink"`
}

type PhysicalProduct struct {
	Product
	Weight float64 `json:"weight"`
	Dimensions string `json:"dimensions"`
}

type SubscriptionProduct struct {
	Product
	SubscriptionPeriod string `json:"subscriptionPeriod"`
	RenewalPrice float64 `json:"renewalPrice"`
}
