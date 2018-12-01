package main

import (
    "math/rand"
    "time"
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

    rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
    venue := venues[rand.Intn(len(venues))]

    fmt.Fprintf(w, "How about %v?", venue.Name)
}

func main() {
    addr := fmt.Sprintf("127.0.0.1:%s", getEnv("PORT", "8080"))
    log.Infof("Starting server on %s", addr)
    http.HandleFunc("/", lunchIdeas)
    log.Fatal(http.ListenAndServe(addr, nil))
}
