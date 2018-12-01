package main

import (
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
    location := os.Getenv("LOCATION")
    client_id := os.Getenv("CLIENT_ID")
    secret := os.Getenv("SECRET")

    fc := NewFoursquareClient(client_id, secret)
    venues := fc.VenuesNearby(location)

    fmt.Fprintf(w, venues)
}

func main() {
    addr := fmt.Sprintf("127.0.0.1:%s", getEnv("PORT", "8080"))
    log.Infof("Starting server on %s", addr)
    http.HandleFunc("/", lunchIdeas)
    log.Fatal(http.ListenAndServe(addr, nil))
}
