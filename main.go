package main

import (
	AuthenController "example/GoapiI/controller/authen"
	"example/GoapiI/orm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Register struct {
	Username string ` json:"username"   binding:"required"`
	Password string ` json:"password"  binding:"required"`
	Fullname string ` json:"fullname"  binding:"required"`
	Avatar   string ` json:"avatar"  binding:"required"`
}

type User struct {
	gorm.Model
	ID       uint
	Username string
	Password string
	Fullname string
	Avatar   string
}

func main() {
	orm.InitDB()

	r := gin.Default()
	r.POST("/register", AuthenController.Register)
	r.POST("/login", AuthenController.Login)

	r.Use(cors.Default())
	r.Run()
}
