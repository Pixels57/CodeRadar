package main

import (
	"fmt"
	"log"
	"os"

	"server/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to the database")
	}

	router := gin.Default()

	userHandler := handlers.NewUserHandler(db)

	router.POST("/users/create", userHandler.CreateUser)
	router.DELETE("/users/delete/:id", userHandler.DeleteUser)
	router.GET("/users", userHandler.GetAllUsers)
	router.GET("/users/skills/:skill", userHandler.GetUsersBySkill)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Server is running on port 8080")
	}

}
