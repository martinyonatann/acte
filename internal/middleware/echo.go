package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/martinyonatann/acte/config"
)

var configCORS = echoMiddleware.CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch},
}

func NewEcho(cfg config.Config) *echo.Echo {
	e := echo.New()
	e.Use(echoMiddleware.RecoverWithConfig(echoMiddleware.RecoverConfig{DisableStackAll: true}))
	e.Use(echoMiddleware.CORSWithConfig(configCORS))
	e.Debug = cfg.Server.Debug

	return e
}
