package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: go run . <城市名>")
		return
	}

	city := os.Args[1]

	lat, lon, err := geocode(city)
	if err != nil {
		fmt.Println("出错了：", err)
		return
	}

	temp, code, err := GetWeather(lat, lon)
	if err != nil {
		fmt.Println("出错了：", err)
		return
	}

	desc := weatherText(code)
	today := time.Now().Local().Format("2006-01-02")

	fmt.Printf("%s %s 当前气温 %.1f℃，天气：%s\n", city, today, temp, desc)

}

func weatherText(code int) string {
	m := map[int]string{
		0:  "晴",
		1:  "基本晴",
		2:  "局部多云",
		3:  "多云",
		45: "雾",
		61: "小雨",
		63: "中雨",
		65: "大雨",
		80: "阵雨",
	}
	if v, ok := m[code]; ok {
		return v
	}
	return "未知"
}
