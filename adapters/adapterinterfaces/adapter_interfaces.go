package adapterinterfaces

import "github.com/akshay0074700747/user-service/entities"

type AdapterInterface interface {
	Adduser(req entities.Users) (entities.Users, error)
	GetUser(id uint) (entities.Users, error)
	GetAllUsers() ([]entities.Users, error)
}