package migration

import (
	"Final/internal/accountSys/consts"
	"Final/internal/accountSys/models"
	"Final/internal/accountSys/secured"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MigrateDatabase() (db *gorm.DB) {
	silentLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	dsn := secured.DB_CONFIG
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: silentLogger,
	})

	if err != nil {
		log.Fatalf(consts.CONNECT_DB_FAILED, err)
	}

	if err := db.AutoMigrate(
		&models.DLModel{},
		&models.SLModel{},
		&models.Voucher{},
		&models.VoucherItem{},
	); err != nil {
		log.Fatalf(consts.MIGRATE_DB_FAILED, err)
	}

	return db
}
