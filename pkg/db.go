package pkg

import (
	"log"

	"github.com/vandenbill/db-tx-clean-arch/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func PostgreConn() *gorm.DB {
	dsn := "host=localhost user=postgres password=root dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Panic(err)
	}

	return db
}

func AutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(domain.User{}); err != nil {
		log.Panic(err)
	}
}
