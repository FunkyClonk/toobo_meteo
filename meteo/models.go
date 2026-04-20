package meteo

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
		Temperature []float64 `json:"apparent_temperature"`
		Wind        []float64 `json:"wind_speed_10m"`
		Rain        []float64 `json:"rain"`
	} `json:"hourly"`
}

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type geocodingResponse struct {
	Results []Coordinate `json:"results"`
}

type clothingAdvice struct {
	temperatureLevel int8 // 0=cold, 1=mid, 2=hot
	isSunny          bool
	isRainy          bool
}
