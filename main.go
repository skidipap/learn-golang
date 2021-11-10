package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"backend/handler"
	"backend/user"
)

func main() {
	db, err := gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})

	if err != nil {
		panic("failed to connect")
	}

	fmt.Println("Succes connect to DB")

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)

	router.Run()

}
