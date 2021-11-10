package main

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"github.com/gin-gonic/gin"

	"backend/user"
	"backend/handler"
)

func main() {
	db, err := gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})

	if err != nil {
		fmt.Println("Succes connect to DB")
		panic("failed to connect")
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()


}
