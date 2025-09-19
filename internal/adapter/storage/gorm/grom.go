package gormOrm

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Trx *gorm.DB

func Init(dsn string, maxIdleConns int, maxOpenConns int, connMaxLifetime int) error {
	var err error
	Trx, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := Trx.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)

	return nil
}
