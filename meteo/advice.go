package meteo

import (
	"fmt"
	"time"
)

func BuildAdvice(weather MeteoResponse) (string, string) {
	advice := &clothingAdvice{temperatureLevel: 1}

	msg := buildWeatherCodeMessage(weather)
	msg += buildTemperatureMessage(weather, advice)
	msg += buildRainMessage(weather, advice)
	msg += buildWindMessage(weather)
	msg += buildDaylightMessage(weather)
	msg += buildSnowMessage(weather)
	msg += buildUVMessage(weather, advice)

	imageRef := getClothingImage(*advice)
	return msg, imageRef
}

func buildWeatherCodeMessage(weather MeteoResponse) string {
	label := getWeatherCodeLabel(weather.Daily.WeatherCode[0])
	return fmt.Sprintf("Aujourd'hui on aura : %s\n", label)
}

func buildTemperatureMessage(weather MeteoResponse, advice *clothingAdvice) string {
	max := weather.Daily.TemperatureMax[0]
	min := weather.Daily.TemperatureMin[0]
	msg := fmt.Sprintf("Température: %.1f - %.1f°C\n", min, max)
	msg += "Avec cette température je te conseil de t'habiller comme ça!\n"

	switch {
	case max > 20:
		msg += "☀️ Un petit pull grand max, il fera chaud aujourd'hui 😎\n"
		advice.temperatureLevel = 2
		advice.isSunny = max > 30
	case min < 2:
		msg += "❄️ Brr, mets un bon manteau, un bonnet et des gants, il fera froid aujourd'hui 🧥🧣\n"
		advice.temperatureLevel = 0
	default:
		msg += "🌤️ Une bonne polaire suffira, il fera bon aujourd'hui 🙂\n"
		advice.temperatureLevel = 1
	}

	msg += fmt.Sprintf("En particulier, à 8h il fera %.1f°C et à 18h il fera %.1f°C\n",
		weather.Hourly.Temperature[8], weather.Hourly.Temperature[18])
	return msg
}

func buildRainMessage(weather MeteoResponse, advice *clothingAdvice) string {
	rainSum := weather.Daily.RainSum[0]
	var msg string

	if rainSum == 0 {
		msg = "Pas de pluie aujourd'hui, youpi!\n"
	} else if weather.Daily.ShowersSum[0] > 0 {
		msg = "De grosses averses sont à prévoir\nJe te conseille de prendre ton parapluie et un kway!\n"
		advice.isRainy = true
	} else if weather.Daily.PrecipitationHours[0] > 5 {
		msg = "Il pleuvera un peu toute la journée, un kway suffira pour te balader tranquillement\n"
		advice.isRainy = true
	} else {
		msg = "Très peu de pluie aujourd'hui, youpi!\n"
	}

	if rain8AM := weather.Hourly.Rain[8]; rain8AM > 0 {
		msg += fmt.Sprintf("Il pleuvera %.1f mm ce matin à 8h, pense à ton kway à l'allée\n", rain8AM)
	}
	if rain6PM := weather.Hourly.Rain[18]; rain6PM > 0 {
		msg += fmt.Sprintf("Il pleuvera %.1f mm ce soir à 18h, pense à ton kway au retour\n", rain6PM)
	}
	return msg
}

func buildWindMessage(weather MeteoResponse) string {
	speed := weather.Daily.WindSpeed[0]
	var msg string
	switch {
	case speed < 40:
		msg = "Pas trop de vent aujourd'hui, une journée calme :)\n"
	case speed < 50:
		msg = "Un peu de vent aujourd'hui, mais rien d'insurmontable!\n"
	case speed < 60:
		msg = "Beaucoup de vent aujourd'hui, fais attention si tu prends le vélo!\n"
	default:
		msg = "Oula, c'est la tempête aujourd'hui, fais attention quand tu sors!\n"
	}

	if wind8AM := weather.Hourly.Wind[8]; wind8AM > 10 {
		msg += fmt.Sprintf("Pas mal de vent ce matin à 8h : %.1f km/h, attention à l'allée!\n", wind8AM)
	}
	if wind6PM := weather.Hourly.Wind[18]; wind6PM > 10 {
		msg += fmt.Sprintf("Pas mal de vent ce soir à 18h : %.1f km/h, attention au retour!\n", wind6PM)
	}
	return msg
}

func buildDaylightMessage(weather MeteoResponse) string {
	sunrise, err := time.Parse("2006-01-02T15:04", weather.Daily.Sunrise[0])
	if err != nil {
		panic(err)
	}
	sunset, err := time.Parse("2006-01-02T15:04", weather.Daily.Sunset[0])
	if err != nil {
		panic(err)
	}

	daylightH := weather.Daily.DayLightDuration[0] / 3600
	sunshineH := weather.Daily.SunshineDuration[0] / 3600

	msg := fmt.Sprintf("Le soleil se lèvera à %s et se couchera à %s\nSoit %.0fh de lumière et %.0fh d'ensoleillement\n",
		sunrise.Format("15:04"), sunset.Format("15:04"), daylightH, sunshineH)

	if sunshineH < 7 {
		msg += "Une journée nuageuse mais pas de quoi se décourager!\n"
	} else {
		msg += "Une journée ensoleillée! Profites-en bien avec tes proches :)\n"
	}
	return msg
}

func buildSnowMessage(weather MeteoResponse) string {
	if snow := weather.Daily.SnowfallSum[0]; snow > 0 {
		return fmt.Sprintf("Oh il va neiger, %.0f cm de neige!\nC'est joli mais fais attention en voiture!\n", snow)
	}
	return ""
}

func buildUVMessage(weather MeteoResponse, advice *clothingAdvice) string {
	if uv := weather.Daily.UVMax[0]; uv > 3 {
		advice.isSunny = true
		return "Le soleil va taper aujourd'hui, prends tes lunettes et ta crème solaire!\n"
	}
	return ""
}

func getClothingImage(advice clothingAdvice) string {
	base := "img/clothing/"
	switch {
	case advice.temperatureLevel == 2 && advice.isSunny:
		return base + "hot_sunny.png"
	case advice.temperatureLevel == 2:
		return base + "hot.png"
	case advice.temperatureLevel == 1 && advice.isRainy:
		return base + "mid_rainy.png"
	case advice.temperatureLevel == 1:
		return base + "mid.png"
	case advice.temperatureLevel == 0 && advice.isRainy:
		return base + "cold_rainy.png"
	default:
		return base + "cold.png"
	}
}
