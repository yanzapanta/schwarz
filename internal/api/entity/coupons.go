package entity

type Coupon struct {
	ID             string `json:"id,omitempty"`
	Discount       int    `json:"discount,omitempty"`
	Code           string `json:"code,omitempty"`
	MinBasketValue int    `json:"min_basket_value,omitempty"`
	ValidFrom      string `json:"valid_from,omitempty"`
	ValidTo        string `json:"valid_to,omitempty"`
	IsActive       int    `json:"-"`
	Status         string `json:"status,omitempty"`
}
