package meteo

import (
	"fmt"
	"time"
)

func BuildAdvice(weather MeteoResponse) (string, string) {
	advice := &clothingAdvice{temperatureLevel: 1}

	//msg := buildWeatherCodeMessage(weather)
	msg := buildTemperatureMessage(weather, advice)
	msg += buildRainMessage(weather, advice)
	msg += buildWindMessage(weather)
	msg += buildDaylightMessage(weather, advice)
	msg += buildSnowMessage(weather)
	msg += buildUVMessage(weather, advice)

	msg += "Je te conseil donc de t'habiller comme ça !"

	imageRef := getClothingImage(*advice)
	return msg, imageRef
}

func buildWeatherCodeMessage(weather MeteoResponse) string {
	label := getWeatherCodeLabel(weather.Daily.WeatherCode[0])
	return fmt.Sprintf("Au programme aujourd'hui : %s ! ☁️\n", label)
}

func buildTemperatureMessage(weather MeteoResponse, advice *clothingAdvice) string {
	max := weather.Daily.TemperatureMax[0]
	min := weather.Daily.TemperatureMin[0]
	msg := fmt.Sprintf("Côté thermomètre, il fera entre %.1f°C et %.1f°C. 🌡️\n", min, max)

	switch {
	case max >= 28:
		msg += "Ouah, c'est la canicule ! Un t-shirt léger et hop, on reste au frais ! ☀️🥵\n"
		advice.temperatureLevel = 2
		advice.isSunny = true
	case max >= 18:
		//msg += "Il va faire tout doux ! Un petit pull suffira pour jouer dehors ! 😎\n"
		advice.temperatureLevel = 2
	case min <= 5:
		//msg += "Brrr, ça caille ! Sortez les gros manteaux, les bonnets et les gants ! ❄️🧥\n"
		advice.temperatureLevel = 0
	default:
		//msg += "Une petite veste ou une polaire et vous serez parfaits ! 😊\n"
		advice.temperatureLevel = 1
	}

	msg += fmt.Sprintf("À 8h il fera %.1f°C et à 18h encore %.1f°C.\n",
		weather.Hourly.Temperature[8], weather.Hourly.Temperature[18])
	return msg
}

func buildRainMessage(weather MeteoResponse, advice *clothingAdvice) string {
	rainSum := weather.Daily.RainSum[0]
	if rainSum == 0 {
		return ""
		//return "Génial, pas une goutte d'eau à l'horizon ! 🌈\n"
	}

	var msg string
	if weather.Daily.ShowersSum[0] > 2 { // Seuil d'averses significatives
		msg = "Attention, il va y avoir de grosses rincées ! Sortez les parapluies et les bottes ! ☔️\n"
		advice.isRainy = true
	} else if weather.Daily.PrecipitationHours[0] > 4 {
		msg = "La pluie va jouer les prolongations toute la journée... Le k-way est obligatoire ! 🌧️\n"
		advice.isRainy = true
	} else {
		msg = "Juste quelques petites gouttes, même pas peur ! 💧\n"
	}
	return msg
}

func buildWindMessage(weather MeteoResponse) string {
	speed := weather.Daily.WindSpeed[0]
	var msg string
	switch {
	case speed < 20:
		msg = ""
		// msg = "Pas de vent aujourd'hui, c'est très calme. 🍃\n"
	case speed < 45:
		msg = "Ça souffle un peu, les feuilles s'amusent ! 🌬️\n"
	default:
		msg = "Ouh là là ! Ça décoiffe ! Tenez bien vos chapeaux, mes p'tits primates ! 💨\n"
	}
	return msg
}

func buildDaylightMessage(weather MeteoResponse, advice *clothingAdvice) string {
	sunrise, _ := time.Parse("2006-01-02T15:04", weather.Daily.Sunrise[0])
	sunset, _ := time.Parse("2006-01-02T15:04", weather.Daily.Sunset[0])
	sunshineH := weather.Daily.SunshineDuration[0] / 3600

	msg := fmt.Sprintf("Le soleil se lève à %s et se couchera à %s.\n",
		sunrise.Format("15h04"), sunset.Format("15h04"))

	if sunshineH > 7 {
		advice.isSunny = true
		//msg += "Le soleil va briller de mille feux ! Profitez-en pour gambader ! 🌻\n"
	} else {
		advice.isSunny = false
		//msg += "Le soleil se cache derrière les nuages, mais gardez le sourire ! ☁️\n"
	}
	return msg
}

func buildSnowMessage(weather MeteoResponse) string {
	if snow := weather.Daily.SnowfallSum[0]; snow > 0 {
		return fmt.Sprintf("Oh il va neiger, %.0f cm de neige!\nC'est joli mais fais attention en vélo!\n", snow)
	}
	return ""
}

func buildUVMessage(weather MeteoResponse, advice *clothingAdvice) string {

	if len(weather.Daily.UVMax) == 0 {
		return ""
	}

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
