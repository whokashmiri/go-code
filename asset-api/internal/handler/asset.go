package handler
//
// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// // 	"errors"
// 	"fmt"
// 	"sync"
// 	"strings"
// 	"io"
//
// 	"example.com/asset-api/internal/model"
// )
//
//
// type CreateAssetRequest struct {
//     Name            string  `json:"name"`
//     ManufactureYear int     `json:"manufactureYear"`
//     EstimatedValue  float64 `json:"estimatedValue"`
//     IsWorking       bool    `json:"isWorking"`
// }
//
//
// func GetAssets(w http.ResponseWriter, r *http.Request) {
//
// 	assets := []model.Asset{
// 		{
// 			ID:              "A001",
// 			Name:            "Camry",
// 			ManufactureYear: 2022,
// 			EstimatedValue:  650000,
// 			IsWorking:       true,
// 		},
// 		{
// 			ID:              "A002",
// 			Name:            "I20",
// 			ManufactureYear: 2020,
// 			EstimatedValue:  250000,
// 			IsWorking:       true,
// 		},
// 	}
//
// 	w.Header().Set(
// 		"Content-Type",
// 		"application/json",
// 	)
//
// 	if err := json.NewEncoder(w).Encode(assets); err != nil {
// 		log.Println("Failed to encode the assets", err)
// 	}
//
// }
//
// func CountAssets(w http.ResponseWriter, r *http.Request) {
//
//     assets := []model.Asset{
//         {
//             ID:              "A001",
//             Name:            "Camry",
//             ManufactureYear: 2022,
//             EstimatedValue:  650000,
//             IsWorking:       true,
//         },
//         {
//             ID:              "A002",
//             Name:            "I20",
//             ManufactureYear: 2020,
//             EstimatedValue:  250000,
//             IsWorking:       true,
//         },
//     }
//
//     total := len(assets)
//
//     w.Header().Set("Content-Type", "application/json")
//
//     response := map[string]int{
//         "count": total,
//     }
//
//     if err := json.NewEncoder(w).Encode(response); err != nil {
//         log.Println("Failed to encode count:", err)
//     }
// }
//
//
//
// type AssetHandler struct{
//     mu sync.RWMutex
//     assets []model.Asset
//     nextID int
//     }
//
// func NewAssetHandler() *AssetHandler{
//         return &AssetHandler{
//             assets: []model.Asset{
//                {
//                             ID:              "A001",
//                             Name:            "Toyota Camry",
//                             ManufactureYear: 2022,
//                             EstimatedValue:  65000,
//                             IsWorking:       true,
//                         },
//                         {
//                             ID:              "A002",
//                             Name:            "CAT 320",
//                             ManufactureYear: 2018,
//                             EstimatedValue:  185000,
//                             IsWorking:       true,
//                         },
//             },
//         nextID: 3,
//         }
//     }
//
//
//
// func writeJSON(
//     w http.ResponseWriter,
//     status int,
//     data any,
// ) {
//     w.Header().Set("Content-Type", "application/json")
//     w.WriteHeader(status)
//
//     if err := json.NewEncoder(w).Encode(data); err != nil {
//         fmt.Println("JSON encoding error:", err)
//     }
// }
//
// func writeError(
//     w http.ResponseWriter,
//     status int,
//     message string,
// ) {
//     writeJSON(w, status, map[string]string{
//         "error": message,
//     })
// }
//
//
//
// func (h *AssetHandler) GetAssets(
//     w http.ResponseWriter,
//     r *http.Request,
// ) {
//     h.mu.RLock()
//
//     assets := make([]model.Asset, len(h.assets))
//     copy(assets, h.assets)
//
//     h.mu.RUnlock()
//
//     writeJSON(w, http.StatusOK, assets)
// }
//
//
//
// func (h *AssetHandler) GetAssetByID(
//     w http.ResponseWriter,
//     r *http.Request,
// ) {
//     id := r.PathValue("id")
//
//     h.mu.RLock()
//     defer h.mu.RUnlock()
//
//     for _, asset := range h.assets {
//         if asset.ID == id {
//             writeJSON(w, http.StatusOK, asset)
//             return
//         }
//     }
//
//     writeError(
//         w,
//         http.StatusNotFound,
//         "asset not found",
//     )
// }
//
//
//
// func (h *AssetHandler) CreateAsset(
//     w http.ResponseWriter,
//     r *http.Request,
// ) {
//     r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
//
//     decoder := json.NewDecoder(r.Body)
//     decoder.DisallowUnknownFields()
//
//     var input CreateAssetRequest
//
//     if err := decoder.Decode(&input); err != nil {
//         writeError(w, http.StatusBadRequest, "invalid JSON body")
//         return
//     }
//
//     // Ensure only one JSON value was provided.
//     var extra any
//     if err := decoder.Decode(&extra); err != io.EOF {
//         writeError(w, http.StatusBadRequest, "expected one JSON object")
//         return
//     }
//
//     input.Name = strings.TrimSpace(input.Name)
//
//     if input.Name == "" {
//         writeError(w, http.StatusBadRequest, "name is required")
//         return
//     }
//
//     if input.ManufactureYear < 1 ||
//         input.ManufactureYear > 2026 {
//         writeError(w, http.StatusBadRequest, "invalid manufacturing year")
//         return
//     }
//
//     if input.EstimatedValue < 0 {
//         writeError(w, http.StatusBadRequest, "estimated value cannot be negative")
//         return
//     }
//
//     h.mu.Lock()
//
//     asset := model.Asset{
//         ID:              fmt.Sprintf("A%03d", h.nextID),
//         Name:            input.Name,
//         ManufactureYear: input.ManufactureYear,
//         EstimatedValue:  input.EstimatedValue,
//         IsWorking:       input.IsWorking,
//     }
//
//     h.nextID++
//     h.assets = append(h.assets, asset)
//
//     h.mu.Unlock()
//
//     writeJSON(w, http.StatusCreated, asset)
// }



import (
        "fmt"
        "bytes"
        "encoding/json"
        "errors"
        "io"
        "net/http"
        "strings"
        "sync"


        "example.com/asset-api/internal/model"


    )

func decodeUpdateRequest(w http.ResponseWriter , r *http.Request,)(UpdateAssetRequest , error){
    var input UpdateAssetRequest

    body , err := io.ReadAll(
        http.MaxBytesReader(w, r.Body, 1<<20),
        )

    if err != nil{
        return input , err
        }
    trimmed :=  body.TrimSpace(body)

    if len(trimmed) == 0 || trimmed[0] != '{' {
        return input, errors.New("expected a JSON object")
        }

    var fields map[string].json.RawMessage

    if err := json.Unmarshal(body , &fields); err != nil{
    return input ,err
    }


    for key , raw :=  range feilds{
        if bytes.Equal(bytes.TrimSpace(raw), []byte("null")){
        return input , fmt.Errorf("feild %s cannot be null" , key)
        }
    }

    decoder := json.NewEncoder(bytes.NewEncoder(body))
    decoder.DisallowUnknownFields()

    if  err:= decoder.Decode(&input); err != nil {
     return input , err
    }

    var extra any

    if err := decoder.Decode(&extra); err != io.EOF{
        return input, errors.New("expected one JSON object")
    }

    return input , nil

    }


func (h *AssetHandler) ReplaceAsset(w http.ResponseWriter , r *Request,){
    id := r.PathValue("id")
    input , err : decodeUpdateRequest(w , r)
    if err != nil {
        writeError(w, http.StatusBadRequest, "invalid JSON :"+err.Error())
        return
    }

    if input.name== nil ||
        input.ManufactureYear == nil ||
        input.EstimatedValue == nil ||
        input.IsWorking ==nil {
        writeError(w, http.StatusBadRequest , "All fields are required")
        return
        }

        name := strings.TrimSpace(*input.Name)

        if name == "" {
        writeError(w.http.StatusBadRequest , "name is required")
        return
        }
        if *input.ManufactureYear < 1 || *input.ManufactureYear > 2026{
            writeError(w, http.StatusBadRequest, "invalid manufacture year" )
            return
        }
     if *input.EstimatedValue < 0 {
            writeError(w, http.StatusBadRequest, "invalid estimated value")
            return
        }

        replacement := model.Asset{
        ID:              id,
        Name:            name,
        ManufactureYear: *input.ManufactureYear,
        EstimatedValue:  *input.EstimatedValue,
        IsWorking:       *input.IsWorking,
        }

         h.mu.Lock()

         for i:= range h.assets{
         if h.assets[i].ID == id{
            h.assets[i] = replacement
            h.mu.Unlock()

            writeJSON(w, http.StatusOK, replacement)
            return

         }
         }

    }