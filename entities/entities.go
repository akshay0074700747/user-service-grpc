package entities

type Clients struct {
	Id       uint `gorm:"primaryKey"`
	Name     string
	Email    string
	Mobile   string
	Password string
}

type Admins struct {
	Id       uint `gorm:"primaryKey"`
	Name     string
	Email    string
	Mobile   string
	Password string
}

type SuAdmins struct {
	Id       uint `gorm:"primaryKey"`
	Name     string
	Email    string
	Mobile   string
	Password string
}
