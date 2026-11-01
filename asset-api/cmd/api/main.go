// package main
//
// import ("fmt"
//      "example.com/asset-api/internal/model"
//       "example.com/asset-api/internal/service"
//     )
//
//
// func main() {
// asset := model.Asset{
//     ID: "Asset-001",
//     Name: "CAT-321 Ex",
//     ManufactureYear:2018,
//     EstimatedValue:185000,
//     IsWorking: true,
//     }
//
// age, err := service.CalculateAge(asset)
//
// if err != nil{
//
//     fmt.Println("Error" , err)
//     return
//     }
//
// inspection , err := service.InspectionType(asset)
//
//     if err !=nil{
//      fmt.Println("Error" , err)
//         return
//     }
//
//
// fmt.Println("====Asset Summary======")
// fmt.Println("ID:", asset.ID)
//     fmt.Println("Name:", asset.Name)
//     fmt.Println("Age:", age)
//     fmt.Printf("Value: %.2f SAR\n", asset.EstimatedValue)
//     fmt.Println("Inspection:", inspection)
//
// }

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"example.com/asset-api/internal/handler"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	//     fmt.Fprintln(w, "Server is healthy")
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	response := map[string]string{
		"status":  "ok",
		"message": "Asset API is running",
	}
	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Println("Response Encoding Error", "err")
	}

}

func main() {

	mux := http.NewServerMux()

	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("GET /api/assets", handler.GetAssets)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Println("Asset API running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Server error:", err)
	}

}
