package database

import (
	"fmt"
	backend "s-store"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase(config backend.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.DbUser,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("Open Database: %w", err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("Connect Database faild: %w", err)
	}

	sqlDb.SetMaxOpenConns(10)
	sqlDb.SetMaxIdleConns(5)
	sqlDb.SetConnMaxIdleTime(time.Hour)

	if err := sqlDb.Ping(); err != nil {
		sqlDb.Close()
		return nil, fmt.Errorf("Ping Database failed: %w", err)
	}

	return db, nil
}
