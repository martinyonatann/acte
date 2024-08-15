package server

import (
	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/acte/internal/billings/delivery"
	"github.com/martinyonatann/acte/internal/billings/delivery/api"
	"github.com/martinyonatann/acte/internal/billings/repository"
	"github.com/martinyonatann/acte/internal/billings/usecase"
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
		billerRepo     = repository.NewBillingRepository(s.db)
		billerUC       = usecase.NewBillingUC(billerRepo)
		billerDelivery = api.NewBillingHandlers(billerUC)
	)

	delivery.MapRoutes(version, billerDelivery)

	return nil
}
