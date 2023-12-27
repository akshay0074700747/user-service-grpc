package db

import (
	"github.com/akshay0074700747/user-service/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(connectTo string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(connectTo), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entities.Clients{})
	db.AutoMigrate(&entities.Admins{})
	db.AutoMigrate(&entities.SuAdmins{})

	return db, nil
}
