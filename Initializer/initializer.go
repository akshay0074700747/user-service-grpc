package initializer

import (
	"github.com/akshay0074700747/user-service/adapters"
	"github.com/akshay0074700747/user-service/services"
	"gorm.io/gorm"
)

func Initialize(db *gorm.DB) *services.UserServiceServer {
	adapter := adapters.NewUserAdapter(db)
	service := services.NewUserServiceServer(adapter)

	return service
}