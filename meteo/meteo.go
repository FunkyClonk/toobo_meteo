package meteo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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
	Hourly struct {
		TemperatureHourly []float64 `json:"apparent_temperature"`
		WindHourly        []float64 `json:"wind_speed_10m"`
		RainHourly        []float64 `json:"rain"`
	} `json:"hourly"`
}

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ApiResponse struct {
	Results []Coordinate `json:"results"`
}

type Advice struct {
	Temperature int8
	Sunny       bool
	Rainny      bool
}

func CallMeteo(city string) MeteoResponse {
	coordinates := GetCoordinatesFromString(city)
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%.4f&longitude=%.4f&daily=weather_code,sunrise,sunset,apparent_temperature_max,apparent_temperature_min,wind_speed_10m_max,daylight_duration,sunshine_duration,showers_sum,snowfall_sum,precipitation_hours,rain_sum,uv_index_max&hourly=wind_speed_10m,apparent_temperature,rain&models=meteofrance_seamless&timezone=Europe%%2FBerlin&forecast_days=1", coordinates.Longitude, coordinates.Latitude)
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

func getImageFromAdvice(advice Advice) string {
	base := "img/clothing/"

	switch {
	// Hot (2)
	case advice.Temperature == 2 && advice.Sunny && !advice.Rainny:
		return base + "hot_sunny.png"
	case advice.Temperature == 2 && !advice.Sunny && advice.Rainny:
		return base + "hot.png"
	case advice.Temperature == 2 && advice.Sunny && advice.Rainny:
		return base + "hot_sunny.png"
	case advice.Temperature == 2 && !advice.Sunny && !advice.Rainny:
		return base + "hot.png"

	// Mid (1)
	case advice.Temperature == 1 && advice.Sunny && !advice.Rainny:
		return base + "mid.png"
	case advice.Temperature == 1 && !advice.Sunny && advice.Rainny:
		return base + "mid_rainy.png"
	case advice.Temperature == 1 && advice.Sunny && advice.Rainny:
		return base + "mid_rainy.png"
	case advice.Temperature == 1 && !advice.Sunny && !advice.Rainny:
		return base + "mid.png"

	// Cold (0)
	case advice.Temperature == 0 && advice.Sunny && !advice.Rainny:
		return base + "cold.png"
	case advice.Temperature == 0 && !advice.Sunny && advice.Rainny:
		return base + "cold_rainy.png"
	case advice.Temperature == 0 && advice.Sunny && advice.Rainny:
		return base + "cold_rainny.png"
	case advice.Temperature == 0 && !advice.Sunny && !advice.Rainny:
		return base + "cold.png"
	}

	// fallback (au cas où valeur inattendue)
	return base + "unknown.png"
}

func GetAdvices(dataMeteo MeteoResponse) (string, string) {
	advice := Advice{
		Temperature: 1,
		Sunny:       false,
		Rainny:      false,
	}

	messageAdvices := getAdviceWeatherCode(dataMeteo)
	messageAdvices += getAdviceTemperature(dataMeteo, &advice)
	messageAdvices += getAdviceRain(dataMeteo, &advice)
	messageAdvices += getAdviceWind(dataMeteo)
	messageAdvices += getAdviceDayTime(dataMeteo)
	messageAdvices += getAdviceSnow(dataMeteo)
	messageAdvices += getAdviceUV(dataMeteo, &advice)
	imageRef := getImageFromAdvice(advice)
	return messageAdvices, imageRef
}

func getAdviceUV(dataMeteo MeteoResponse, adviceStruct *Advice) string {
	advice := ""
	if dataMeteo.Daily.UVMax[0] != 0 {
		uvMax := dataMeteo.Daily.UVMax[0]
		if uvMax > 3 {
			advice = "Le soleil va tapper aujourd'hui, prend tes lunettes et ta crème solaire!\n"
			adviceStruct.Sunny = true
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
	sunshineDurationHours := float64(dataMeteo.Daily.SunshineDuration[0]) / 3600
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
	wind8AM := dataMeteo.Hourly.WindHourly[8]
	if wind8AM > 10 {
		advice += fmt.Sprintf("Il y aura pas mal de vent ce matin à 8h, %.1f km/h quand même! Attention à l'allé!\n", wind8AM)
	}
	wind6PM := dataMeteo.Hourly.WindHourly[18]
	if wind6PM > 10 {
		advice += fmt.Sprintf("Il y aura pas mal de vent ce soir à 18h, %.1f km/h quand même! Attention au retour!\n", wind6PM)
	}
	return advice
}

func getAdviceWeatherCode(dataMeteo MeteoResponse) string {
	weatherSignification := getWeatherCodeTraduction(dataMeteo.Daily.WeatherCode[0])
	advice := fmt.Sprintf("Aujourd'hui on aura : %s\n", weatherSignification)
	return advice
}

func getAdviceRain(dataMeteo MeteoResponse, adviceStruct *Advice) string {
	rainSum := dataMeteo.Daily.RainSum[0]
	var advice string
	if rainSum == 0 {
		advice = "Pas de pluie aujourd'hui, youpi!\n"
	} else {
		showerSum := dataMeteo.Daily.ShowersSum[0]
		if showerSum > 0 {
			advice = "De grosses averses sont a prévoires\n Je te conseil de prendre ton parapluie et du kway pour te protéger\n"
			adviceStruct.Rainny = true
		} else {
			precipitationHours := dataMeteo.Daily.PrecipitationHours[0]
			if precipitationHours > 5 {
				advice = "Il pleuvera un peu toute la journée, tu peut prendre un kway pour te balader en toute tranquillité\n"
				adviceStruct.Rainny = true
			} else {
				advice = "Très peu de pluie aujourd'hui, youpi!\n"
			}
		}
	}
	rain8AM := dataMeteo.Hourly.RainHourly[8]
	rain6PM := dataMeteo.Hourly.RainHourly[18]
	if rain8AM > 0 {
		advice += fmt.Sprintf("Il pleuvera %.1f mm ce matin à 8h, tu pourras mettre ton kway à l'allée\n", rain8AM)
	}
	if rain6PM > 0 {
		advice += fmt.Sprintf("Il pleuvera %.1f mm ce soir à 18h, tu pourras mettre ton kway au retour\n", rain6PM)
	}
	return advice
}

func getAdviceTemperature(dataMeteo MeteoResponse, adviceStruct *Advice) string {
	advice := fmt.Sprintf("Température: %.1f - %.1f°C\n", dataMeteo.Daily.TemperatureMin[0], dataMeteo.Daily.TemperatureMax[0])
	advice += "Avec cette température je te conseil je t'habiller comme ca!\n"
	temperatureMax := dataMeteo.Daily.TemperatureMax[0]
	temperatureMin := dataMeteo.Daily.TemperatureMin[0]
	if temperatureMax > 20 {
		advice += "☀️ Un petit pull grand max, il fera chaud aujourd'hui 😎\n"
		adviceStruct.Temperature = 2
		//While UV isn't working, use temp to get the sun
		adviceStruct.Sunny = temperatureMax > 30
	} else if temperatureMin < 2 {
		advice += "❄️ Brr, mets un bon manteau, un bonnet et des gants, il fera froid aujourd'hui 🧥🧣\n"
		adviceStruct.Temperature = 0
	} else {
		advice += "🌤️ Une bonne polaire suffira, il fera bon aujourd'hui 🙂\n"
		adviceStruct.Temperature = 1
	}
	advice += fmt.Sprintf("En particulier, A 8h il fera %.1f°C et à 18h il fera %.1f°C\n", dataMeteo.Hourly.TemperatureHourly[8], dataMeteo.Hourly.TemperatureHourly[18])
	return advice
}

func GetCoordinatesFromString(city string) Coordinate {
	url := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=en&format=json", city)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	if len(data.Results) == 0 {
		panic("no results found")
	}

	coord := data.Results[0]

	return coord
}
