package datasources

import (
	"context"
	"fmt"

	"github.com/martinyonatann/acte/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMySQLDB(ctx context.Context, cfg config.DatabaseConfig) (*gorm.DB, error) {
	otps := &gorm.Config{SkipDefaultTransaction: true}

	if !cfg.Debug {
		otps = &gorm.Config{SkipDefaultTransaction: true, Logger: logger.Default.LogMode(logger.Silent)}
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Asia%%2FJakarta", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), otps)
	if err != nil {
		panic(err)
	}

	return db, nil
}
