package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/asset-api/internal/model"
)

func GetAssets(w http.ResponseWriter, r *http.Request) {

	assets := []model.Asset{
		{
			ID:              "A001",
			Name:            "Camry",
			ManufactureYear: 2022,
			EstimatedValue:  650000,
			IsWorking:       true,
		},
		{
			ID:              "A002",
			Name:            "I20",
			ManufactureYear: 2020,
			EstimatedValue:  250000,
			IsWorking:       true,
		},
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	if err := json.NewEncoder(w).Encode(assets); err != nil {
		log.Println("Failed to encode the assets", err)
	}

}
