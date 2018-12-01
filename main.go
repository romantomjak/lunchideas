package main

import (
    "os"
    "fmt"
    "net/http"
    "log"
)

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func lunchIdeas(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world!")
}

func main() {
    port := getEnv("PORT", "8080")
    addr := fmt.Sprintf(":%s", port)

    http.HandleFunc("/", lunchIdeas)
    log.Fatal(http.ListenAndServe(addr, nil))
}
