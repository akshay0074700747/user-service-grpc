package adapterinterfaces

import "github.com/akshay0074700747/user-service/entities"

type AdapterInterface interface {
	Signup(req entities.Clients) (entities.Clients, error)
	GetUser(id uint) (entities.Clients, error)
	GetAdmin(id uint) (entities.Admins, error)
	GetSuAdmin(id uint) (entities.SuAdmins, error)
	LoginUser(email, password string) (entities.Clients, error)
	LoginSuAdmin(email, password string) (entities.SuAdmins, error)
	LoginAdmin(email, password string) (entities.Admins, error)
	GetAllUsers() ([]entities.Clients, error)
	GetAllAdmins() ([]entities.Admins, error)
	GetPassByEmail(email string, isAdmin, isSuAdmin bool) (string, error)
	AddAdmin(req entities.Admins) (entities.Admins,error)
}
