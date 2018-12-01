package main

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "time"

    log "github.com/sirupsen/logrus"
)

type FoursquareClient struct {
    clientId string
    secret string
    baseURL string
}

type response struct {
    Response json.RawMessage `json:"response"`
}

type groupList struct {
    Groups []group `json:"groups"`
}

type groupItem struct {
    Venue Venue `json:"venue"`
}

type group struct {
    Type  string `json:"type"`
    Name  string `json:"name"`
    Items []groupItem `json:"items"`
}

type Venue struct {
    Name string `json:"name"`
}

func NewFoursquareClient(clientId string, secret string) *FoursquareClient {
    return &FoursquareClient{
        clientId: clientId,
        secret: secret,
        baseURL: "https://api.foursquare.com/v2/",
    }
}

func (fc *FoursquareClient) VenuesNearby(location string) []Venue {
    body, err := fc.query("venues/explore/?near=" + location + "&query=lunch")
    if err != nil {
        log.Error("Failed to query Foursquare Venues")
        return nil
    }

    r := response{}
    err = json.Unmarshal(body, &r)
    if err != nil {
        log.Error("Failed to deserialize API response")
        return nil
    }

    gl := &groupList{}
    err = json.Unmarshal(r.Response, &gl)
    if err != nil {
        log.Error("Failed to deserialize API response")
        return nil
    }

    venues := []Venue{}
    for i := range gl.Groups[0].Items {
        venues = append(venues, gl.Groups[0].Items[i].Venue)
    }
    return venues
}

func (fc *FoursquareClient) query(reqPath string) (response []byte, err error) {
    req, err := http.NewRequest(http.MethodGet, fc.baseURL + reqPath, nil)
    if err != nil {
        log.Errorf("Foursquare API error: %v", err)
        return []byte{}, err
    }

    q := req.URL.Query()
    q.Add("client_id", fc.clientId)
    q.Add("client_secret", fc.secret)
    q.Add("v", "20181201") // foursquare API "version"
    req.URL.RawQuery = q.Encode()
    req.Header.Set("User-Agent", "LunchIdeasBot/1.0 (+https://lunchideas.herokuapp.com)")

    client := http.Client{
        Timeout: time.Second * 2,
    }
    res, err := client.Do(req)
    if err != nil {
        log.Errorf("Foursquare API error: %v", err)
        return []byte{}, nil
    }

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Errorf("Foursquare API error: %v", err)
        return []byte{}, nil
    }

    return body, nil
}
