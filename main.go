package main

import (
	repository "myapp/Respository"
	router "myapp/Router"
	useccase "myapp/Usecase"
)

func main() {
	dataFile := "drinks.json"
	dataRepo := repository.NewDataRepository(dataFile)
	DataUsecase := useccase.NewDataHandler(dataRepo)

	app := router.NewApp(DataUsecase)

	if err := app.Start("8080"); err != nil {
		panic(err)
	}
}
