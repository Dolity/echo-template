package router

import (
	"myapp/handler"
	"myapp/service"

	"context"
	"log"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitRoutes(e *echo.Echo) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://dolity:dolity123456@cluster0.9jq9yah.mongodb.net/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("myapp")
	userService := service.NewUserService(db)
	userHandler := handler.NewUserHandler(userService)

	api := e.Group("/api/v1")
	api.POST("/users", userHandler.CreateUser)
	api.GET("/users/:id", userHandler.GetUserById)
}