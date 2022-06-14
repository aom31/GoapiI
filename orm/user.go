package orm

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint
	Username string
	Password string
	Fullname string
	Avatar   string
}
