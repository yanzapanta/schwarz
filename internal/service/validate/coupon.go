package validate

import (
	"coupon_service/internal/api/entity"
	"errors"
	"time"
)

func ValidateCoupon(coupon entity.Coupon, validDateTime string) error {
	var err error
	if coupon.Discount <= 0 {
		err = errors.New("invalid discount")
	}
	if coupon.MinBasketValue <= 0 {
		err = errors.New("invalid minimum basket value")
	}
	validFrom, validFromErr := time.Parse(validDateTime, coupon.ValidFrom)
	if validFromErr != nil {
		err = errors.New("invalid valid from value")
	}
	validTo, validToErr := time.Parse(validDateTime, coupon.ValidTo)
	if validToErr != nil {
		err = errors.New("invalid valid until value")
	}
	if validTo.Unix() <= validFrom.Unix() {
		err = errors.New("invalid validity dates")
	}

	return err
}
