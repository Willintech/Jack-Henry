package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"

    "github.com/<Jack-Henry>/weather-server/weather"
)

func main() {
    http.HandleFunc("/forecast", forecastHandler)
    fmt.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func forecastHandler(w http.ResponseWriter, r *http.Request) {
    latStr := r.URL.Query().Get("lat")
    lonStr := r.URL.Query().Get("lon")

    if latStr == "" || lonStr == "" {
        http.Error(w, "lat and lon query parameters are required", http.StatusBadRequest)
        return
    }

    lat, err := strconv.ParseFloat(latStr, 64)
    if err != nil {
        http.Error(w, "invalid latitude", http.StatusBadRequest)
        return
    }

    lon, err := strconv.ParseFloat(lonStr, 64)
    if err != nil {
        http.Error(w, "invalid longitude", http.StatusBadRequest)
        return
    }

    forecast, tempType, err := weather.GetForecast(lat, lon)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := map[string]string{
        "forecast": forecast,
        "temperature_type": tempType,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
