package controllers

import (
	model "myapp/Model"
	usecase "myapp/Usecase"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	data []model.Drinkmodel
}

func HomeController(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}

func GetDataDrink(c *fiber.Ctx, h *usecase.DataUsecase) error {
	return h.GetAllData(c)
}

func GetData(c *fiber.Ctx, h *usecase.DataUsecase) error {
	return h.GetData(c)
}

func AddData(c *fiber.Ctx, h *usecase.DataUsecase) error {
	return h.AddData(c)
}

func UpdateData(c *fiber.Ctx, h *usecase.DataUsecase) error {
	return h.UpdateData(c)
}

func DeleteData(c *fiber.Ctx, h *usecase.DataUsecase) error {
	return h.DeleteData(c)
}
