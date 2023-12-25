package entities

type Users struct {
	Id   uint `gorm:"primaryKey"`
	Name string
	IsAdmin bool
}
