package authen

import (
	"net/http"

	"example/GoapiI/orm"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterBody struct {
	Username string ` json:"username"   binding:"required"`
	Password string ` json:"password"  binding:"required"`
	Fullname string ` json:"fullname"  binding:"required"`
	Avatar   string ` json:"avatar"  binding:"required"`
}

func Register(c *gin.Context) {
	var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//check user id is exist
	// Get first matched record
	var userExist orm.User
	orm.Db.Where("username = ?", json.Username).First(&userExist)
	if userExist.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Username is existed"})
		return
	}

	//create user
	enablePass, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 8)
	user := orm.User{Username: json.Username, Password: string(enablePass),
		Fullname: json.Fullname, Avatar: json.Avatar}
	orm.Db.Create(&user)
	if user.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"userId": user.ID,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "message": "User cannot Register!!"})
	}
}

type LoginBody struct {
	Username string ` json:"username"   binding:"required"`
	Password string ` json:"password"  binding:"required"`
}

func Login(c *gin.Context) {
	var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
