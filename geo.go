package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type GenResponse struct {
	Results []struct {
		Latitude  float64 `json:"latitude"`  // 纬度，查天气要用
		Longitude float64 `json:"longitude"` // 经度
		Name      string  `json:"name"`      // 匹配到的城市名（用来回显）
		Country   string  `json:"country"`   // 国家，用来区分重名城市
	} `json:"results"`
}

func geocode(city string) (float64, float64, error) {
	q := url.QueryEscape(city)
	apiURL := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=zh", q)

	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("地理编码接口返回状态码 %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	var g GenResponse

	if err := json.Unmarshal(body, &g); err != nil {
		return 0, 0, err
	}

	if len(g.Results) == 0 {
		return 0, 0, fmt.Errorf("找不到城市：%s", city)
	}

	r := g.Results[0]

	return r.Latitude, r.Longitude, nil

}
