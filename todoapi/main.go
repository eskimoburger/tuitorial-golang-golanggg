package main

import (
	"fmt"
	"todoapi/config"
	"todoapi/config/db"
	"todoapi/controller"
	"todoapi/model"
	"todoapi/router"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	err := godotenv.Load("local.env")

	if err != nil {
		//log.Fatal("Error loading .env file")
		fmt.Println("Error loading .env file %v \n", err)
	}
	cfg := config.NewConfig()

	storage, err := db.InitDb(cfg.ConnDB)

	if err != nil {
		panic("failed to connect to db")
	}

	storage.AutoMigrate(&model.Todo{})
	controller.NewStorage(storage)

	// value := os.Getenv("PORT")

	e := echo.New()

	router.InitRouter(e)

	e.Logger.Fatal(e.Start(cfg.Port))
}
