package adapters

import "github.com/akshay0074700747/user-service/entities"

func (user *UserAdapter) Signup(req entities.Clients) (entities.Clients, error) {

	var res entities.Clients
	query := "INSERT INTO clients (name,email,mobile,password) VALUES($1,$2,$3,$4) RETURNING id,name,email,mobile"

	return res, user.DB.Raw(query, req.Name, req.Email, req.Mobile, req.Password).Scan(&res).Error
}

func (user *UserAdapter) GetUser(id uint) (entities.Clients, error) {

	var res entities.Clients
	query := "SELECT * FROM clients WHERE id = $1"

	return res, user.DB.Raw(query, id).Scan(&res).Error
}

func (user *UserAdapter) GetAdmin(id uint) (entities.Admins, error) {

	var res entities.Admins
	query := "SELECT * FROM admins WHERE id = $1"

	return res, user.DB.Raw(query, id).Scan(&res).Error
}

func (user *UserAdapter) GetSuAdmin(id uint) (entities.SuAdmins, error) {

	var res entities.SuAdmins
	query := "SELECT * FROM su_admins WHERE id = $1"

	return res, user.DB.Raw(query, id).Scan(&res).Error
}

func (user *UserAdapter) LoginUser(email, password string) (entities.Clients, error) {

	var res entities.Clients
	query := "SELECT * FROM clients WHERE email = $1 AND password = $2"

	return res, user.DB.Raw(query, email, password).Scan(&res).Error
}

func (user *UserAdapter) LoginSuAdmin(email, password string) (entities.SuAdmins, error) {

	var res entities.SuAdmins
	query := "SELECT * FROM su_admins WHERE email = $1 AND password = $2"

	return res, user.DB.Raw(query, email, password).Scan(&res).Error
}

func (user *UserAdapter) LoginAdmin(email, password string) (entities.Admins, error) {

	var res entities.Admins
	query := "SELECT * FROM admins WHERE email = $1 AND password = $2"

	return res, user.DB.Raw(query, email, password).Scan(&res).Error
}

func (user *UserAdapter) GetAllUsers() ([]entities.Clients, error) {

	var res []entities.Clients
	query := "SELECT * FROM clients"

	return res, user.DB.Raw(query).Scan(&res).Error
}

func (user *UserAdapter) GetAllAdmins() ([]entities.Admins, error) {

	var res []entities.Admins
	query := "SELECT * FROM admins"

	return res, user.DB.Raw(query).Scan(&res).Error
}

func (user *UserAdapter) GetPassByEmail(email string, isAdmin, isSuAdmin bool) (string, error) {

	var query, password string

	if isAdmin {
		query = "SELECT password FROM admins WHERE email = $1"
		return password, user.DB.Raw(query, email).Scan(&password).Error
	} else if isSuAdmin {
		query = "SELECT password FROM su_admins WHERE email = $1"
		return password, user.DB.Raw(query, email).Scan(&password).Error
	}

	query = "SELECT password FROM clients WHERE email = $1"
	return password, user.DB.Raw(query, email).Scan(&password).Error

}

func (user *UserAdapter) AddAdmin(req entities.Admins) (entities.Admins, error) {

	var res entities.Admins
	query := "INSERT INTO admins (name,email,mobile,password) VALUES($1,$2,$3,$4) RETURNING id,name,email,mobile"

	return res, user.DB.Raw(query, req.Name, req.Email, req.Mobile, req.Password).Scan(&res).Error
}
