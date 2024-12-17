package entity

import (
	_ "github.com/gin-gonic/gin"
)

type Basket struct {
	Value                 int     `json:"value,omitempty"`
	AppliedDiscount       int     `json:"applied_discount,omitempty"`
	ApplicationSuccessful bool    `json:"application_successful,omitempty"`
	Total                 float64 `json:"total,omitempty"`
}
