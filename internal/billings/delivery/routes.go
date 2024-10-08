package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/acte/internal/billings/delivery/api"
)

func MapBillingRoutes(echo *echo.Group, h *api.BillingHandler) {
	billers := echo.Group("/billing")
	billers.POST("/create", h.CreateLoan)
	billers.GET("/:id/detail", h.DetailLoan)
	billers.POST("/:id/inquiry", h.InquireRepaymentAmount)
	billers.POST("/:id/payment", h.LoanRepayment)
	billers.GET("/delinquent/list", h.ListDelinquentLoan)
}
