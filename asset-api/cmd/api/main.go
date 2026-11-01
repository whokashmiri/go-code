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

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /health", healthHandler)
    mux.HandleFunc("GET /api/assets", handler.GetAssets)
    mux.HandleFunc("GET /api/assets/count", handler.CountAssets)

    server := &http.Server{
        Addr:              ":8080",
        Handler:           mux,
        ReadHeaderTimeout: 5 * time.Second,
    }

    log.Println("Asset API running on http://localhost:8080")

    if err := server.ListenAndServe(); err != nil {
        log.Fatal("Server error:", err)
    }
}

func healthHandler(
    w http.ResponseWriter,
    r *http.Request,
) {
    w.Header().Set(
        "Content-Type",
        "application/json",
    )

    response := map[string]string{
        "status":  "ok",
        "message": "Asset API is running",
    }

    if err := json.NewEncoder(w).Encode(response); err != nil {
        log.Println("Failed to encode health response:", err)
    }
}
