package router

import (
	useccase "myapp/Usecase"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	FiberApp *fiber.App
}

func NewApp(dataHandler *useccase.DataUsecase) *App {
	app := fiber.New()

	app.Get("/data", dataHandler.GetAllData)
	app.Get("/data/:id", dataHandler.GetData)
	app.Post("/dataAdd", dataHandler.AddData)
	app.Put("/dataUpdate/:id", dataHandler.UpdateData)
	app.Delete("/dataDelete/:id", dataHandler.DeleteData)
	return &App{FiberApp: app}
}

func (a *App) Start(port string) error {
	return a.FiberApp.Listen(":" + port)
}
