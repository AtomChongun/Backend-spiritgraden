package model

type Drinkmodel struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
	Mixed []Mixed `json:"mixed"`
}

type Mixed struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Volume      int     `json:"volume"`
	Price       float64 `json:"price"`
	IsAvailable bool    `json:"isAvailable"`
}
