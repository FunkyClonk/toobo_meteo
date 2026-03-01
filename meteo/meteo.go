package meteo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"os"
)

type MeteoResponse struct {
	Daily struct {
		Sunrise            []string  `json:"sunrise"`
		Sunset             []string  `json:"sunset"`
		TemperatureMax     []float64 `json:"apparent_temperature_max"`
		TemperatureMin     []float64 `json:"apparent_temperature_min"`
		WeatherCode        []int     `json:"weather_code"`
		WindSpeed          []float64 `json:"wind_speed_10m_max"`
		DayLightDuration   []float64 `json:"daylight_duration"`
		SunshineDuration   []float64 `json:"sunshine_duration"`
		ShowersSum         []float64 `json:"showers_sum"`
		SnowfallSum        []float64 `json:"snowfall_sum"`
		PrecipitationHours []float64 `json:"precipitation_hours"`
		RainSum            []float64 `json:"rain_sum"`
		UVMax              []float64 `json:"uv_index_max"`
	} `json:"daily"`
}

func CallMeteo() MeteoResponse {
	url := os.Getenv("METEO_URL")

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data MeteoResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	return data
}

func PrintMeteo(dataMeteo MeteoResponse) {
	fmt.Printf("Température max: %.1f°C\n", dataMeteo.Daily.TemperatureMax[0])
	fmt.Printf("Température min: %.1f°C\n", dataMeteo.Daily.TemperatureMin[0])
	fmt.Printf("Code météo: %d \n", dataMeteo.Daily.WeatherCode[0])
	fmt.Printf("Vitesse du vent max: %.1f km/h\n", dataMeteo.Daily.WindSpeed[0])
	fmt.Printf("Lever du soleil: %s\n", dataMeteo.Daily.Sunrise[0])
	fmt.Printf("Coucher du soleil: %s\n", dataMeteo.Daily.Sunset[0])
	fmt.Printf("Durée de la lumière du jour: %.0f min\n", dataMeteo.Daily.DayLightDuration[0])
	fmt.Printf("Durée d'ensoleillement: %.0f min\n", dataMeteo.Daily.SunshineDuration[0])
	fmt.Printf("Pluie totale: %.1f mm\n", dataMeteo.Daily.RainSum[0])
	fmt.Printf("Neige totale: %.1f cm\n", dataMeteo.Daily.SnowfallSum[0])
	fmt.Printf("Heures de précipitation: %.1f h\n", dataMeteo.Daily.PrecipitationHours[0])
	fmt.Printf("Somme des averses: %.1f mm\n", dataMeteo.Daily.ShowersSum[0])
	if dataMeteo.Daily.UVMax[0] != 0 {
		fmt.Printf("Indice UV max: %.1f \n", dataMeteo.Daily.UVMax[0])
	}
}

func GetAdviceTemperature(dataMeteo MeteoResponse) string {
	var advice string
	temperatureMax := dataMeteo.Daily.TemperatureMax[0]
	temperatureMin := dataMeteo.Daily.TemperatureMin[0]
	if temperatureMax > 20 {
		advice = "☀️ Un petit pull grand max, il fera chaud aujourd'hui 😎\n"
	} else if temperatureMin < 2 {
		advice = "❄️ Brr, mets un bon manteau, il fera froid aujourd'hui 🧥🧣\n"
	} else {
		advice = "🌤️ Une bonne polaire suffira, il fera bon aujourd'hui 🙂\n"
	}
	return advice
}

func getAdviceRain(dataMeteo MeteoResponse) string {
	var advice string
	temperatureMax := dataMeteo.Daily.TemperatureMax[0]
	temperatureMin := dataMeteo.Daily.TemperatureMin[0]
	if temperatureMax > 20 {
		advice = "☀️ Un petit pull grand max, il fera chaud aujourd'hui 😎\n"
	} else if temperatureMin < 2 {
		advice = "❄️ Brr, mets un bon manteau, il fera froid aujourd'hui 🧥🧣\n"
	} else {
		advice = "🌤️ Une bonne polaire suffira, il fera bon aujourd'hui 🙂\n"
	}
	return advice
}
