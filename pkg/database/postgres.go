package database

import (
	"fmt"
	"log"

	"agnos/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(host, port, user, password, dbname string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	log.Println("Database connection successfully established.")

	err = db.AutoMigrate(
		&model.Hospital{},
		&model.Staff{},
		&model.Patient{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	SeedHospitals(db)

	log.Println("Database migration completed.")
	return db, nil
}