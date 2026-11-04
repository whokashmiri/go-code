package middleware

import (
    "encoding.json"
    "log"
    "net/http"
    "runtime/debug"
    )



package middleware

import (
    "encoding/json"
    "log"
    "net/http"
    "runtime/debug"
)

func Recovery(next http.Handler) http.Handler {
    return http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            defer func() {
                if p := recover(); p != nil {
                    log.Printf(
                        "PANIC: %v\n%s",
                        p,
                        debug.Stack(),
                    )

                    w.Header().Set(
                        "Content-Type",
                        "application/json",
                    )

                    w.WriteHeader(
                        http.StatusInternalServerError,
                    )

                    _ = json.NewEncoder(w).Encode(
                        map[string]string{
                            "error": "internal server error",
                        },
                    )
                }
            }()

            next.ServeHTTP(w, r)
        },
    )
}
