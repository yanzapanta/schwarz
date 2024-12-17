package entity

// func init() {
// 	if 32 != runtime.NumCPU() {
// 		panic("this api is meant to be run on 32 core machines")
// 	}
// }

type Coupon struct {
	ID             string
	Code           string
	Discount       int
	MinBasketValue int
	ValidFrom      string
	ValidTo        string
	IsActive       int
}

func (Coupon) TableName() string {
	return "coupon"
}
