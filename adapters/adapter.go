package adapters

import "github.com/akshay0074700747/user-service/entities"

func (user *UserAdapter) Adduser(req entities.Users) (entities.Users, error) {

	var res entities.Users
	query := "INSERT INTO users (name,is_admin) VALUES($1,$2) RETURNING id,name,is_admin "

	return res, user.DB.Raw(query, req.Name, req.IsAdmin).Scan(&res).Error
}

func (user *UserAdapter) GetUser(id uint) (entities.Users, error) {

	var res entities.Users
	query := "SELECT * FROM users WHERE id = $1"

	return res, user.DB.Raw(query, id).Scan(&res).Error
}

func (user *UserAdapter) GetAllUsers() ([]entities.Users, error) {

	var res []entities.Users
	query := "SELECT * FROM users"

	return res, user.DB.Raw(query).Scan(&res).Error
}
