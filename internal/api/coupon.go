package api

import (
	"coupon_service/internal/api/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// applies coupon
func (a *API) Apply(c *gin.Context) {
	apiReq := entity.ApplicationRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	basket, err := a.svc.ApplyCoupon(apiReq.Basket, apiReq.Code)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, basket)
}

// creates a coupon
func (a *API) Create(c *gin.Context) {
	apiReq := entity.Coupon{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	coupon, err := a.svc.CreateCoupon(apiReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, coupon)
}

// retrieves coupons by code
func (a *API) Get(c *gin.Context) {
	apiReq := entity.CouponRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	coupons, err := a.svc.GetCoupons(apiReq.Codes)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, coupons)
}
