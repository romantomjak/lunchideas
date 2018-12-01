package main

import (
    "encoding/json"
    "os"
    "fmt"
    "net/http"

    log "github.com/sirupsen/logrus"
)

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func lunchIdeas(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    location := os.Getenv("LOCATION")
    client_id := os.Getenv("CLIENT_ID")
    secret := os.Getenv("SECRET")

    fc := NewFoursquareClient(client_id, secret)
    venues := fc.VenuesNearby(location)

    json, err := json.Marshal(venues)
    if err != nil {
        log.Error("Failed to serialize JSON response")
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(json)
}

func main() {
    addr := fmt.Sprintf("127.0.0.1:%s", getEnv("PORT", "8080"))
    log.Infof("Starting server on %s", addr)
    http.HandleFunc("/", lunchIdeas)
    log.Fatal(http.ListenAndServe(addr, nil))
}
