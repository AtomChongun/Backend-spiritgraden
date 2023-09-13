package useccase

import (
	model "myapp/Model"
	repository "myapp/Respository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type DataUsecase struct {
	DataRepo repository.DataRepository
}

func NewDataHandler(dataRepo repository.DataRepository) *DataUsecase {
	return &DataUsecase{DataRepo: dataRepo}
}

func (h *DataUsecase) GetAllData(c *fiber.Ctx) error {
	data, err := h.DataRepo.GetAllData()
	if err != nil {
		return err
	}

	return c.JSON(data)
}

func (h *DataUsecase) GetData(c *fiber.Ctx) error {
	data, err := h.DataRepo.GetAllData()
	if err != nil {
		return err
	}
	id := c.Params("id")
	intId, err := strconv.Atoi(id)
	{
		if err != nil {
			return c.Status(400).JSON(map[string]string{
				"error": "id must be an integer",
			})
		}
		for _, d := range data {
			if d.Id == intId {
				return c.JSON(d)
			}
		}
		return c.Status(404).JSON(map[string]string{
			"error": "drink not found",
		})
	}
}

func (h *DataUsecase) AddData(c *fiber.Ctx) error {
	drinks, err := h.DataRepo.GetAllData()
	if err != nil {
		return err
	}
	drinknew := new(model.Drinkmodel)
	if err := c.BodyParser(drinknew); err != nil {
		return err
	}
	drinknew.Id = len(drinks) + 1
	drinks = append(drinks, *drinknew)
	repository.DataUpdate(drinks)
	return c.JSON(drinknew)
}

func (h *DataUsecase) UpdateData(c *fiber.Ctx) error {
	drinks, err := h.DataRepo.GetAllData()
	if err != nil {
		return err
	}
	id := c.Params("id")
	intId, err := strconv.Atoi(id)
	{
		if err != nil {
			return c.Status(400).JSON(map[string]string{
				"error": "id must be an integer",
			})
		}
		for i, d := range drinks {
			if d.Id == intId {
				drinknew := new(model.Drinkmodel)
				if err := c.BodyParser(drinknew); err != nil {
					return err
				}
				drinks[i] = *drinknew
				repository.DataUpdate(drinks)
				return c.JSON(drinknew)
			}
		}
		return c.Status(404).JSON(map[string]string{
			"error": "drink not found",
		})
	}
}

func (h *DataUsecase) DeleteData(c *fiber.Ctx) error {
	data, err := h.DataRepo.GetAllData()
	if err != nil {
		return err
	}
	id := c.Params("id")
	for i, drinknew := range data {
		if strconv.Itoa(drinknew.Id) == id {
			data = append(data[:i], data[i+1:]...)
			repository.DataUpdate(data)
			return c.JSON(drinknew)
		}
	}
	return c.Status(404).JSON(map[string]string{
		"error": "drink not found",
	})
}
