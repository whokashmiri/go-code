
package config

import (
    "fmt"
    "os"
    "strconv"
    "time"
)

type Config struct {
    Port            int
    Environment     string
    ReadTimeout     time.Duration
    WriteTimeout    time.Duration
    ShutdownTimeout time.Duration
}

func Load() (Config, error) {
    cfg := Config{
        Port:            8080,
        Environment:     "development",
        ReadTimeout:     10 * time.Second,
        WriteTimeout:    15 * time.Second,
        ShutdownTimeout: 10 * time.Second,
    }

    if value, ok := os.LookupEnv("PORT"); ok {
        port, err := strconv.Atoi(value)
        if err != nil || port < 1 || port > 65535 {
            return cfg, fmt.Errorf(
                "invalid PORT: %q",
                value,
            )
        }

        cfg.Port = port
    }

    if value, ok := os.LookupEnv("APP_ENV"); ok {
        if value != "development" &&
            value != "production" {
            return cfg, fmt.Errorf(
                "invalid APP_ENV: %q",
                value,
            )
        }

        cfg.Environment = value
    }

    if value, ok := os.LookupEnv("READ_TIMEOUT"); ok {
        duration, err := time.ParseDuration(value)
        if err != nil || duration <= 0 {
            return cfg, fmt.Errorf(
                "invalid READ_TIMEOUT: %q",
                value,
            )
        }

        cfg.ReadTimeout = duration
    }

    if value, ok := os.LookupEnv("WRITE_TIMEOUT"); ok {
        duration, err := time.ParseDuration(value)
        if err != nil || duration <= 0 {
            return cfg, fmt.Errorf(
                "invalid WRITE_TIMEOUT: %q",
                value,
            )
        }

        cfg.WriteTimeout = duration
    }

    if value, ok := os.LookupEnv("SHUTDOWN_TIMEOUT"); ok {
        duration, err := time.ParseDuration(value)
        if err != nil || duration <= 0 {
            return cfg, fmt.Errorf(
                "invalid SHUTDOWN_TIMEOUT: %q",
                value,
            )
        }

        cfg.ShutdownTimeout = duration
    }

    return cfg, nil
}
