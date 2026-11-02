package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"errors"
	"fmt"
	"sync"
	"strings"
	"io"

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

func CountAssets(w http.ResponseWriter, r *http.Request) {

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

    total := len(assets)

    w.Header().Set("Content-Type", "application/json")

    response := map[string]int{
        "count": total,
    }

    if err := json.NewEncoder(w).Encode(response); err != nil {
        log.Println("Failed to encode count:", err)
    }
}



type AssetHandler struct{
    mu sync.RWMutex
    assets []model.Asset
    nextID int
    }

func NewAssetHandler() *AssetHandler{
        return &AssetHandler{
            assets: []model.Asset{
               {
                            ID:              "A001",
                            Name:            "Toyota Camry",
                            ManufactureYear: 2022,
                            EstimatedValue:  65000,
                            IsWorking:       true,
                        },
                        {
                            ID:              "A002",
                            Name:            "CAT 320",
                            ManufactureYear: 2018,
                            EstimatedValue:  185000,
                            IsWorking:       true,
                        },
            },
        nextID: 3,
        }
    }


func writeJSON(w http.ResponseWriter, status int , data any){
        w.Header().Set("Content-Type" , "application/json")
        w.WriteHeader(status)
        if err :=  json.NewEncoder(w).Encode(data); err!=nil{
        fmt.Println("JSON encoding error" , err)
        }
    }

func writeError(w http.ResponseWriter, status int , message string){
    writeJSON(w, status,  map[string] string){
        "error" :message,
        }
    }