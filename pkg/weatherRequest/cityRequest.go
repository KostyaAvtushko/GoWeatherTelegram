package weatherRequest

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

type Weather struct {
	Temp              float64
	RealTemp          float64
	WeatherStatus     string
	DescWeatherStatus string
	WindSpeed         float64
	Status            string
}

func GetWeatherCity(city string) *Weather {

	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=de45be9a3632b0b8dbdfcea3fe09004b")
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	weatherJson := string(body)
	if gjson.Get(weatherJson, "cod").String() == "200" {
		tempJson, err := strconv.ParseFloat(gjson.Get(weatherJson, "main.temp").String(), 64)
		tempJsonC := math.Round(tempJson - 273.15)

		windSpeedJson, err := strconv.ParseFloat(gjson.Get(weatherJson, "wind.speed").String(), 64)

		weatherStatusJson := gjson.Get(weatherJson, "weather.0.main").String()

		descWeatherStatusJson := gjson.Get(weatherJson, "weather.0.description").String()

		flTempJson, err := strconv.ParseFloat(gjson.Get(weatherJson, "main.feels_like").String(), 64)
		flTempJsonC := math.Round(flTempJson - 273.15)

		if err != nil {
			log.Panic(err)
		}
		fmt.Println(tempJsonC, flTempJsonC, windSpeedJson, weatherStatusJson, descWeatherStatusJson)
		return &Weather{
			tempJsonC,
			flTempJsonC,
			weatherStatusJson,
			descWeatherStatusJson,
			windSpeedJson,
			"ok",
		}
	}
	return &Weather{0, 0, "", "", 0, "No city found"}
}
