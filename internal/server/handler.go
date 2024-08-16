package server

import (
	"github.com/labstack/echo/v4"

	billingDelivery "github.com/martinyonatann/acte/internal/billings/delivery"
	billingHandler "github.com/martinyonatann/acte/internal/billings/delivery/api"
	billingRepository "github.com/martinyonatann/acte/internal/billings/repository"
	billingUC "github.com/martinyonatann/acte/internal/billings/usecase"

	reconciliationDelivery "github.com/martinyonatann/acte/internal/reconciliations/delivery"
	reconciliationHandler "github.com/martinyonatann/acte/internal/reconciliations/delivery/api"
	reconciliationUC "github.com/martinyonatann/acte/internal/reconciliations/usecase"
)

func (s *Server) mapHandlers() error {
	var (
		serviceName = "acte"
		version     = s.echo.Group("api/v1/" + serviceName)
	)

	version.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, "welcome to acte service 👋")
	})

	var (
		billerRepo     = billingRepository.NewBillingRepository(s.db)
		billerUC       = billingUC.NewBillingUC(billerRepo)
		billerHandlers = billingHandler.NewBillingHandlers(billerUC)

		reconciliationUC       = reconciliationUC.NewreconciliationUC()
		reconciliationHandlers = reconciliationHandler.NewReconciliationHandlers(reconciliationUC)
	)

	billingDelivery.MapBillingRoutes(version, billerHandlers)
	reconciliationDelivery.MapReconciliationRoutes(version, reconciliationHandlers)

	return nil
}
