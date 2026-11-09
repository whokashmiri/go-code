
package main

import (
    "context"
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"

    "example.com/asset-api/internal/config"
    "example.com/asset-api/internal/handler"
    "example.com/asset-api/internal/middleware"
)

func main() {
    cfg, err := config.Load()
    if err != nil {
        log.Fatal("Configuration error: ", err)
    }

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

    // Keep any additional routes from previous lessons.

    appHandler := middleware.Logging(
        middleware.RequestID(
            middleware.Recovery(mux),
        ),
    )

    server := &http.Server{
        Addr:              fmt.Sprintf(":%d", cfg.Port),
        Handler:           appHandler,
        ReadHeaderTimeout: 5 * time.Second,
        ReadTimeout:       cfg.ReadTimeout,
        WriteTimeout:      cfg.WriteTimeout,
        IdleTimeout:       60 * time.Second,
    }

    signalCtx, stop := signal.NotifyContext(
        context.Background(),
        os.Interrupt,
    )
    defer stop()

    serverErrors := make(chan error, 1)

    go func() {
        log.Printf(
            "Asset API started | env=%s | port=%d",
            cfg.Environment,
            cfg.Port,
        )

        serverErrors <- server.ListenAndServe()
    }()

    // Wait for either a signal or a server failure.
    select {
    case <-signalCtx.Done():
        log.Println("Shutdown signal received")

    case err := <-serverErrors:
        if !errors.Is(err, http.ErrServerClosed) {
            log.Printf("HTTP server failed: %v", err)
            return
        }

        return
    }

    // Restore normal interrupt behavior.
    stop()

    shutdownCtx, cancel := context.WithTimeout(
        context.Background(),
        cfg.ShutdownTimeout,
    )
    defer cancel()

    log.Println("Starting graceful shutdown")

    if err := server.Shutdown(shutdownCtx); err != nil {
        log.Printf("Graceful shutdown failed: %v", err)

        if closeErr := server.Close(); closeErr != nil {
            log.Printf("Force close error: %v", closeErr)
        }
    }

    log.Println("Server stopped")
}

func healthHandler(
    w http.ResponseWriter,
    r *http.Request,
) {
    w.Header().Set("Content-Type", "application/json")

    response := map[string]string{
        "status": "ok",
    }

    if err := json.NewEncoder(w).Encode(response); err != nil {
        log.Printf("Health response error: %v", err)
    }
}
