package main

import (
	"context"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/martinyonatann/acte/config"
	"github.com/martinyonatann/acte/internal/server"
	"github.com/samber/lo"
)

func main() {
	cfg := lo.Must(config.LoadConfig("config"))

	_ = lo.Must(time.LoadLocation(cfg.Server.TimeZone))

	if err := server.NewServer(context.Background(), cfg).Run(); err != nil {
		log.Warn(err)
	}
}
