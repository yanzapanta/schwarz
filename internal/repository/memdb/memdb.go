package memdb

import (
	ce "coupon_service/internal/api/entity"
	"coupon_service/internal/config"
	"coupon_service/internal/service/entity"

	"github.com/jinzhu/copier"
)

type Config struct{}

type repository interface {
	FindByCode(string) (*entity.Coupon, error)
	Save(entity.Coupon) (*ce.Coupon, error)
}

type Repository struct {
	entries map[string]entity.Coupon
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) FindByCode(code string) (*ce.Coupon, error) {
	if len(code) == 0 {
		return nil, nil
	}
	var dbCoupon entity.Coupon
	if err := config.DB().Where("code = ?", code).First(&dbCoupon).Error; err != nil {
		return nil, err
	}

	var coupon ce.Coupon
	copier.Copy(&coupon, dbCoupon)

	return &coupon, nil
}

func (r *Repository) Save(newCoupon entity.Coupon) (*ce.Coupon, error) {
	if err := config.DB().Save(&newCoupon).Error; err != nil {
		return nil, err
	}
	var coupon ce.Coupon
	copier.Copy(&coupon, newCoupon)

	return &coupon, nil
}
