package repository

import (
	"encoding/json"
	"fmt"

	model "myapp/Model"
	"os"
)

type DataRepository interface {
	GetAllData() ([]model.Drinkmodel, error)
}

type dataRepository struct {
	DataFile string
}

func NewDataRepository(dataFile string) DataRepository {
	return &dataRepository{DataFile: dataFile}
}

func DataUpdate(data []model.Drinkmodel) error {
	file, err := os.Create("drinks.json")
	if err != nil {
		fmt.Println("ไม่สามารถสร้างไฟล์ JSON:", err)

	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("ไม่สามารถเขียนข้อมูล JSON:", err)
		return err

	}
	return nil
}

func (r *dataRepository) GetAllData() ([]model.Drinkmodel, error) {
	file, err := os.Open(r.DataFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	var data []model.Drinkmodel

	if err := decoder.Decode(&data); err != nil {
		fmt.Println(err)
	}
	return data, nil
}
