package meteo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchWeather(city string) MeteoResponse {
	coords := fetchCoordinates(city)
	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%.4f&longitude=%.4f&daily=weather_code,sunrise,sunset,apparent_temperature_max,apparent_temperature_min,wind_speed_10m_max,daylight_duration,sunshine_duration,showers_sum,snowfall_sum,precipitation_hours,rain_sum,uv_index_max&hourly=wind_speed_10m,apparent_temperature,rain&models=meteofrance_seamless&timezone=Europe%%2FBerlin&forecast_days=1",
		coords.Longitude, coords.Latitude,
	)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data MeteoResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		panic(err)
	}
	return data
}

func fetchCoordinates(city string) Coordinate {
	url := fmt.Sprintf(
		"https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=en&format=json",
		city,
	)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data geocodingResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		panic(err)
	}
	if len(data.Results) == 0 {
		panic("no coordinates found for city: " + city)
	}
	return data.Results[0]
}
