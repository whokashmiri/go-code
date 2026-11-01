package model

type Asset struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	ManufactureYear int     `json:"manufactureYear"`
	EstimatedValue  float64 `json:"estimatedValue"`
	IsWorking       bool    `json:"isWorking"`
}
