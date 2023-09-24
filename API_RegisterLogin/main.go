package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	AuthController "hitgub.com/Mickeythitiwut/Api_registerlogin/controller/auth"
	UsersController "hitgub.com/Mickeythitiwut/Api_registerlogin/controller/user"
	"hitgub.com/Mickeythitiwut/Api_registerlogin/middleware"
	"hitgub.com/Mickeythitiwut/Api_registerlogin/orm"
)

// binding from json
type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}

type User struct {
	gorm.Model
	Username string
	Password string
	Fullname string
	Avatar   string
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	orm.InitDB()

	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)
	authorized := r.Group("/users", middleware.JWTAuthen())
	authorized.GET("/readall", UsersController.ReadAll)
	authorized.GET("/profile", UsersController.Profile)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
