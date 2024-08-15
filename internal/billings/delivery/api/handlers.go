package api

import (
	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/acte/internal/billings"
	"github.com/martinyonatann/acte/internal/billings/dtos"
	"github.com/martinyonatann/acte/pkg/app_error"
	"github.com/martinyonatann/acte/pkg/pagination"
	"github.com/martinyonatann/acte/pkg/response"
)

type BillingHandler struct {
	uc billings.BillingUC
}

func NewBillingHandlers(uc billings.BillingUC) *BillingHandler {
	return &BillingHandler{uc: uc}
}

func (h *BillingHandler) CreateLoan(c echo.Context) error {
	var request dtos.CreateLoanRequest
	if err := c.Bind(&request); err != nil {
		return response.ErrorBuilder(app_error.BadRequest(err)).Send(c)
	}

	if err := h.uc.CreateLoan(c.Request().Context(), request); err != nil {
		return response.ErrorBuilder(err).Send(c)
	}

	return response.SuccessBuilder(nil).Send(c)
}

func (h *BillingHandler) DetailLoan(c echo.Context) error {
	var request struct {
		ID int64 `param:"id"`
	}
	if err := c.Bind(&request); err != nil {
		return response.ErrorBuilder(app_error.BadRequest(err)).Send(c)
	}

	loan, err := h.uc.DetailLoan(c.Request().Context(), request.ID)
	if err != nil {
		return response.ErrorBuilder(err).Send(c)
	}

	return response.SuccessBuilder(loan).Send(c)
}

func (h *BillingHandler) InquireRepaymentAmount(c echo.Context) error {
	var request dtos.InquireRepaymentAmountRequest
	if err := c.Bind(&request); err != nil {
		return response.ErrorBuilder(app_error.BadRequest(err)).Send(c)
	}

	if err := request.Validate(); err != nil {
		return response.ErrorBuilder(app_error.BadRequest(err)).Send(c)
	}

	inquiry, err := h.uc.InquireRepaymentAmount(c.Request().Context(), request)
	if err != nil {
		return response.ErrorBuilder(err).Send(c)
	}

	return response.SuccessBuilder(inquiry).Send(c)
}

func (h *BillingHandler) LoanRepayment(c echo.Context) error {
	var request dtos.LoanRepaymentRequest
	if err := c.Bind(&request); err != nil {
		return response.ErrorBuilder(app_error.BadRequest(err)).Send(c)
	}

	if err := request.Validate(); err != nil {
		return response.ErrorBuilder(app_error.BadRequest(err)).Send(c)
	}

	if err := h.uc.LoanRepayment(c.Request().Context(), request); err != nil {
		return response.ErrorBuilder(err).Send(c)
	}

	return response.SuccessBuilder(nil).Send(c)
}

func (h *BillingHandler) ListDelinquentLoan(c echo.Context) error {
	var params pagination.PaginationRequest
	if err := c.Bind(&params); err != nil {
		return response.ErrorBuilder(app_error.BadRequest(err)).Send(c)
	}

	loans, err := h.uc.ListDelinquentLoan(c.Request().Context(), &params)
	if err != nil {
		return response.ErrorBuilder(err).Send(c)
	}

	return response.SuccessBuilder(loans, params).Send(c)
}
