package infra

import (
	"fmt"

	"github.com/yafiesetyo/poc-workflow/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta"
	return gorm.Open(postgres.Open(fmt.Sprintf(
		dsn,
		config.Cfg.DB.Host,
		config.Cfg.DB.User,
		config.Cfg.DB.Password,
		config.Cfg.DB.DBName,
		config.Cfg.DB.Port)), &gorm.Config{})
}
