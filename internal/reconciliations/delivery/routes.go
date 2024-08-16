package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/acte/internal/reconciliations/delivery/api"
)

func MapReconciliationRoutes(echo *echo.Group, h *api.ReconciliationHandler) {
	reconciliations := echo.Group("/reconciliation")
	reconciliations.GET("/check", h.Reconcile)
}
