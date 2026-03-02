package meteo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

func GetAdvices(dataMeteo MeteoResponse) string {
	advices := getAdviceWeatherCode(dataMeteo)
	advices += getAdviceTemperature(dataMeteo)
	advices += getAdviceRain(dataMeteo)
	advices += getAdviceWind(dataMeteo)
	advices += getAdviceDayTime(dataMeteo)
	advices += getAdviceSnow(dataMeteo)
	advices += getAdviceUV(dataMeteo)
	return advices
}

func getAdviceUV(dataMeteo MeteoResponse) string {
	advice := ""
	if dataMeteo.Daily.UVMax[0] != 0 {
		fmt.Printf("Indice UV max: %.1f \n", dataMeteo.Daily.UVMax[0])
		uvMax := dataMeteo.Daily.UVMax[0]
		if uvMax > 3 {
			advice = "Le soleil va tapper aujourd'hui, prend tes lunettes et ta crème solaire!\n"
		}

	}
	return advice
}

func getAdviceSnow(dataMeteo MeteoResponse) string {
	snowfallSum := dataMeteo.Daily.SnowfallSum[0]
	var advice string
	if snowfallSum > 0 {
		advice = fmt.Sprintf("Oh il va neiger, %1.f cm de neige !\n C'est jolie mais fait attention en voiture!\n", snowfallSum)
	}
	return advice
}

func getAdviceDayTime(dataMeteo MeteoResponse) string {
	srTime, err := time.Parse("2006-01-02T15:04", dataMeteo.Daily.Sunrise[0])
	if err != nil {
		panic(err)
	}
	sunrise := srTime.Format("15:04")
	ssTime, err := time.Parse("2006-01-02T15:04", dataMeteo.Daily.Sunset[0])
	if err != nil {
		panic(err)
	}
	sunset := ssTime.Format("15:04")
	dayLightDurationHours := float64(dataMeteo.Daily.DayLightDuration[0]) / 3600
	fmt.Println(dayLightDurationHours) // 2
	sunshineDurationHours := float64(dataMeteo.Daily.SunshineDuration[0]) / 3600
	fmt.Println(sunshineDurationHours) // 2
	advice := fmt.Sprintf("Le soleil se levera à %s et se couchera à %s\nDonc %.0fh de lumière du jour et %.0fh d'ensolleiment\n", sunrise, sunset, dayLightDurationHours, sunshineDurationHours)
	if sunshineDurationHours < 7 {
		advice += "Une journée nuageuse mais pas de quoi se décourager!\n"
	} else {
		advice += "Une journée ensolleilé! Profite en bien avec tes proches :)\n"
	}
	return advice
}
func getAdviceWind(dataMeteo MeteoResponse) string {
	windSpeed := dataMeteo.Daily.WindSpeed[0]
	var advice string
	switch {
	case windSpeed < 40:
		advice = "Pas trop de vent aujourd'hui, une journée calme :)\n"
	case windSpeed < 50:
		advice = "Un peu de vent aujourd'hui, mais rien d'insurmontable!\n"
	case windSpeed < 60:
		advice = "Beaucoup de vent aujourd'hui, fait attention si tu prend le vélo!\n"
	case windSpeed > 60:
		advice = "Oula, c'est la tempete aujourd'hui, fait attention quand tu sors!\n"
	}
	return advice
}

func getAdviceWeatherCode(dataMeteo MeteoResponse) string {
	weatherSignification := getWeatherCodeTraduction(dataMeteo.Daily.WeatherCode[0])
	advice := fmt.Sprintf("Aujourd'hui on aura : %s\n", weatherSignification)
	return advice
}

func getAdviceRain(dataMeteo MeteoResponse) string {
	rainSum := dataMeteo.Daily.RainSum[0]
	var advice string
	if rainSum == 0 {
		advice = "Pas de pluie aujourd'hui, youpi!\n"
	} else {
		showerSum := dataMeteo.Daily.ShowersSum[0]
		if showerSum > 0 {
			advice = "De grosses averses sont a prévoires\n Je te conseil de prendre ton parapluie et du kway pour te protéger\n"
		} else {
			precipitationHours := dataMeteo.Daily.PrecipitationHours[0]
			if precipitationHours > 5 {
				advice = "Il pleuvera un peu toute la journée, tu peut prendre un kway pour te balader en toute tranquillité\n"
			} else {
				advice = "Très peu de pluie aujourd'hui, youpi!\n"
			}
		}
	}
	return advice
}

func getAdviceTemperature(dataMeteo MeteoResponse) string {
	advice := fmt.Sprintf("Température: %.1f - %.1f°C\n", dataMeteo.Daily.TemperatureMin[0], dataMeteo.Daily.TemperatureMax[0])
	advice += "Avec cette température je te conseil je t'habiller comme ca!\n"
	temperatureMax := dataMeteo.Daily.TemperatureMax[0]
	temperatureMin := dataMeteo.Daily.TemperatureMin[0]
	if temperatureMax > 20 {
		advice += "☀️ Un petit pull grand max, il fera chaud aujourd'hui 😎\n"
	} else if temperatureMin < 2 {
		advice += "❄️ Brr, mets un bon manteau, un bonnet et des gants, il fera froid aujourd'hui 🧥🧣\n"
	} else {
		advice += "🌤️ Une bonne polaire suffira, il fera bon aujourd'hui 🙂\n"
	}
	return advice
}
