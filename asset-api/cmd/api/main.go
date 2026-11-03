
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

    assetHandler := handler.NewAssetHandler()

    mux.HandleFunc("GET /health", healthHandler)

    mux.HandleFunc(
        "GET /api/assets",
        assetHandler.GetAssets,
    )

    mux.HandleFunc(
        "GET /api/assets/{id}",
        assetHandler.GetAssetByID,
    )

    mux.HandleFunc(
        "POST /api/assets",
        assetHandler.CreateAsset,
    )
mux.HandleFunc(
    "PUT /api/assets/{id}",
    assetHandler.ReplaceAsset,
)

mux.HandleFunc(
    "PATCH /api/assets/{id}",
    assetHandler.PatchAsset,
)

mux.HandleFunc(
    "DELETE /api/assets/{id}",
    assetHandler.DeleteAsset,
)

    server := &http.Server{
        Addr:              ":8080",
        Handler:           mux,
        ReadHeaderTimeout: 5 * time.Second,
    }

    log.Println("API running at http://localhost:8080")

    if err := server.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}

func healthHandler(
    w http.ResponseWriter,
    r *http.Request,
) {
    w.Header().Set("Content-Type", "application/json")

    json.NewEncoder(w).Encode(map[string]string{
        "status": "ok",
    })
}
