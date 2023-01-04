package toolkit

import (
	"strconv"
	"time"

	owm "github.com/briandowns/openweathermap"
)

func GetWeather(city string) string {
	w, err := owm.NewCurrent("C", "zh_cn", "458a9af0d4e1661e8be8ab5d771b6613")
	if err != nil {
		return ""
	}
	err = w.CurrentByName(city)
	if err != nil {
		return ""
	}
	temp := strconv.FormatFloat(w.Main.Temp, 'f', 2, 64)
	weather := MergeText(w.Name+":",
		w.Weather[0].Description+" "+temp+"°C",
		"风速："+strconv.FormatFloat(w.Wind.Speed, 'f', 2, 64)+"m/s",
		"湿度："+strconv.Itoa(w.Main.Humidity)+"%",
		"气压："+strconv.FormatFloat(w.Main.Pressure, 'f', 2, 64)+"hPa",
		"能见度："+strconv.Itoa(w.Visibility)+"m",
		"云："+strconv.Itoa(w.Clouds.All)+"%",
		"今日日出："+unixTimetoTime(w.Sys.Sunrise),
		"今日日落："+unixTimetoTime(w.Sys.Sunset),
	)
	return weather
}

func unixTimetoTime(unixTime int) string {
	time := time.Unix(int64(unixTime), 0)
	return time.Format("15:04:05")
}
