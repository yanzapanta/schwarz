package service

import (
	ce "coupon_service/internal/api/entity"
	"coupon_service/internal/service/entity"
	"coupon_service/internal/service/validate"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const dateTimeFormat = "2006-01-02T15:04:05Z07:00"
const validDateTime = "2006-01-02 15:04:05"

type Repository interface {
	FindByCode(string) (*ce.Coupon, error)
	Save(entity.Coupon) (*ce.Coupon, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) ApplyCoupon(basket entity.Basket, code string) (b *entity.Basket, e error) {
	b = &basket
	// check if the coupon code exists
	coupon, err := s.repo.FindByCode(code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid code")
		}
		return nil, err
	}
	// check if the coupon is within the validity date
	validFrom, _ := time.Parse(dateTimeFormat, coupon.ValidFrom)
	validTo, _ := time.Parse(dateTimeFormat, coupon.ValidTo)
	now := time.Now()
	if now.Before(validFrom) || now.After(validTo) || coupon.IsActive == 0 {
		return nil, errors.New("invalid code mm")
	}
	// check if the basket value has reached the minimum value
	if b.Value < coupon.MinBasketValue {
		return nil, errors.New("basket value didn't reach the minimum value")
	}

	b.AppliedDiscount = coupon.Discount
	b.ApplicationSuccessful = true
	b.Total = float64(basket.Value - coupon.Discount)

	return b, nil
}

func (s Service) CreateCoupon(coupon ce.Coupon) (*ce.Coupon, error) {
	// check if the coupon code already exists
	existingCoupon, err := s.repo.FindByCode(coupon.Code)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingCoupon != nil {
		return nil, errors.New("coupon code already exists")
	}
	// validate requests
	if err := validate.ValidateCoupon(coupon, validDateTime); err != nil {
		return nil, err
	}

	newCoupon := entity.Coupon{
		Discount:       coupon.Discount,
		Code:           coupon.Code,
		MinBasketValue: coupon.MinBasketValue,
		ValidFrom:      coupon.ValidFrom,
		ValidTo:        coupon.ValidTo,
		ID:             uuid.NewString(),
		IsActive:       1,
	}

	savedCoupon, err := s.repo.Save(newCoupon)
	if err != nil {
		return nil, err
	}

	return savedCoupon, nil
}

func (s Service) GetCoupons(codes []string) (map[string]ce.Coupon, error) {
	coupons := make(map[string]ce.Coupon, len(codes))
	var e error = nil
	// check the request
	if len(codes) == 0 {
		e = errors.New("codes are required")
		return coupons, e
	}

	for _, code := range codes {
		coupon, err := s.repo.FindByCode(code)
		var newCoupon ce.Coupon
		if err != nil {
			newCoupon.Status = err.Error()
		} else {
			newCoupon = validateCoupon(coupon)
		}
		coupons[code] = newCoupon
	}

	return coupons, e
}

// check if coupon is valid, invalid or expired
func validateCoupon(coupon *ce.Coupon) (newCoupon ce.Coupon) {
	validFrom, _ := time.Parse(dateTimeFormat, coupon.ValidFrom)
	validTo, _ := time.Parse(dateTimeFormat, coupon.ValidTo)
	now := time.Now()

	if coupon.IsActive == 0 || now.Before(validFrom) {
		newCoupon.Status = "invalid coupon"
	} else if now.After(validTo) {
		newCoupon.Status = "expired coupon"
	} else {
		newCoupon = *coupon
	}
	return newCoupon
}
