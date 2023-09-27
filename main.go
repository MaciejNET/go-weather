package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

type Location struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Current struct {
	TempC     float64   `json:"temp_c"`
	Condition Condition `json:"condition"`
}

type Forecast struct {
	Forecastday []Forecastday `json:"forecastday"`
}

type Forecastday struct {
	Hour []Hour `json:"hour"`
}

type Hour struct {
	TimeEpoch    int64     `json:"time_epoch"`
	TempC        float64   `json:"temp_c"`
	Condition    Condition `json:"condition"`
	ChanceOfRain float64   `json:"chance_of_rain"`
}

type Condition struct {
	Text string `json:"text"`
}

type Weather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
	Forecast Forecast `json:"forecast"`
}

func main() {
	city := "Warsaw"
	key := "768d1128494447d8ae9120518232708"
	numberOfHours := 12

	if len(os.Args) >= 2 {
		city = os.Args[1]
	}

	if len(os.Args) >= 3 {
		city = os.Args[1]

		num, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
			return
		}

		if num <= 0 || num > 24 {
			log.Fatal("Number of hours must be between 1 and 24")
			return
		}

		numberOfHours = num
	}

	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + key + "&q=" + city + "&days=2&aqi=no&alerts=no")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Servic is not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, current, forecastdays := weather.Location, weather.Current, weather.Forecast.Forecastday

	var hours []Hour
	for _, day := range forecastdays {
		hours = append(hours, day.Hour...)
	}

	color.Set(color.FgYellow)
	fmt.Printf(
		"%s, %s\n",
		location.Name,
		location.Country,
	)
	color.Unset()

	color.Set(color.FgGreen)
	fmt.Printf(
		"Now: %.0fC, %s\n",
		current.TempC,
		current.Condition.Text,
	)
	color.Unset()

	count := 0
	for _, hour := range hours {
		if count == numberOfHours {
			return
		}

		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()) {
			continue
		}

		message := fmt.Sprintf(
			"%s - %.0fC, %.0f%%, %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)

		if hour.ChanceOfRain < 40 {
			fmt.Print(message)
		} else {
			color.Red(message)
		}

		count++
	}
}
