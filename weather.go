package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherResponse struct {
	Current struct {
		Temperature2m float64 `json:"temperature_2m"` //当前气温
		WeatherCode   int     `json:"weather_code"`   //天气代码
	} `json:"current"`
}

func GetWeather(lat, lon float64) (float64, int, error) {
	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%.4f&longitude=%.4f&current=temperature_2m,weather_code",
		lat, lon)

	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("天气接口返回状态码 %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	var w WeatherResponse

	if err := json.Unmarshal(body, &w); err != nil {
		return 0, 0, err
	}

	return w.Current.Temperature2m, w.Current.WeatherCode, nil
}
