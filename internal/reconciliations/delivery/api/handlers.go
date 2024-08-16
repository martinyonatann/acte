package api

import (
	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/acte/internal/reconciliations"
	"github.com/martinyonatann/acte/internal/reconciliations/dtos"
	"github.com/martinyonatann/acte/pkg/app_error"
	"github.com/martinyonatann/acte/pkg/response"
)

type ReconciliationHandler struct {
	uc reconciliations.ReconciliationUC
}

func NewReconciliationHandlers(uc reconciliations.ReconciliationUC) *ReconciliationHandler {
	return &ReconciliationHandler{uc: uc}
}

func (h *ReconciliationHandler) Reconcile(c echo.Context) error {
	var request dtos.ReconcileRequest
	if err := c.Bind(&request); err != nil {
		return response.ErrorBuilder(app_error.BadRequest(err)).Send(c)
	}

	if err := request.Validate(); err != nil {
		return response.ErrorBuilder(app_error.BadRequest(err)).Send(c)
	}

	reconcileData, err := h.uc.TransactionReconciliation(c.Request().Context(), request)
	if err != nil {
		return response.ErrorBuilder(err).Send(c)
	}

	return response.SuccessBuilder(reconcileData).Send(c)
}
