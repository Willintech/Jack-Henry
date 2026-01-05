package weather 

import (
    "encoding/json"
    "errors"
    "fmt"
    "net/http"
)

// API responses
type PointsResponse struct {
    Properties struct {
        Forecast string `json:"forecast"`
    } `json:"properties"`
}

type ForecastResponse struct {
    Properties struct {
        Periods []struct {
            Name        string  `json:"name"`
            DetailedForecast string `json:"detailedForecast"`
            Temperature int     `json:"temperature"`
        } `json:"periods"`
    } `json:"properties"`
}

// GetForecast fetches today's forecast and characterizes temperature
func GetForecast(lat, lon float64) (string, string, error) {
    // Step 1: Get the forecast URL from /points
    pointsURL := fmt.Sprintf("https://api.weather.gov/points/%f,%f", lat, lon)
    resp, err := http.Get(pointsURL)
    if err != nil {
        return "", "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", "", errors.New("failed to fetch forecast points")
    }

    var points PointsResponse
    if err := json.NewDecoder(resp.Body).Decode(&points); err != nil {
        return "", "", err
    }

    forecastURL := points.Properties.Forecast
    if forecastURL == "" {
        return "", "", errors.New("no forecast URL found for coordinates")
    }

    // Step 2: Get the forecast
    forecastResp, err := http.Get(forecastURL)
    if err != nil {
        return "", "", err
    }
    defer forecastResp.Body.Close()

    if forecastResp.StatusCode != http.StatusOK {
        return "", "", errors.New("failed to fetch forecast")
    }

    var forecast ForecastResponse
    if err := json.NewDecoder(forecastResp.Body).Decode(&forecast); err != nil {
        return "", "", err
    }

    if len(forecast.Properties.Periods) == 0 {
        return "", "", errors.New("no forecast periods available")
    }

    // Step 3: Take the first period (Today)
    today := forecast.Properties.Periods[0]
    tempType := characterizeTemperature(today.Temperature)

    return today.DetailedForecast, tempType, nil
}

// characterizeTemperature maps temperature to hot/cold/moderate
func characterizeTemperature(temp int) string {
    switch {
    case temp >= 85:
        return "hot"
    case temp <= 60:
        return "cold"
    default:
        return "moderate"
    }
}

