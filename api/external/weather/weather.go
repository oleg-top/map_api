package weatherParser

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/codingsince1985/geo-golang/openstreetmap"
)

type IWeatherParser interface {
	SetLocation(string)
	SetPeriod(time.Time, time.Time)
	ShowWeather() []Weather
}

type Weather struct {
	Date      string  `json:"date"`
	MaxTemp   float64 `json:"max_temp"`
	MinTemp   float64 `json:"min_temp"`
	Condition string  `json:"condition"`
}

type weatherAPIParser struct {
	BaseURL  string
	Token    string
	Dates    []time.Time
	Location LatLng
}

type weatherAPIDayCondition struct {
	Text string `json:"text"`
}

type weatherAPIDay struct {
	MaxTemp   float64                `json:"maxtemp_c"`
	MinTemp   float64                `json:"mintemp_c"`
	Condition weatherAPIDayCondition `json:"condition"`
}

type weatherAPIForecastDay struct {
	Date string        `json:"date"`
	Day  weatherAPIDay `json:"day"`
}

type weatherAPIForecast struct {
	ForecastDays []weatherAPIForecastDay `json:"forecastday"`
}

type weatherAPIResponse struct {
	Forecast weatherAPIForecast `json:"forecast"`
}

type LatLng struct {
	Lat float64
	Lng float64
}

func (wp *weatherAPIParser) SetLocation(location string) {
	loc, _ := openstreetmap.Geocoder().Geocode(location)

	if loc == nil {
		return
	}

	wp.Location = LatLng{Lat: loc.Lat, Lng: loc.Lng}
}

func (wp *weatherAPIParser) SetPeriod(fromDate, toDate time.Time) {
	var dates []time.Time

	for last := fromDate; last.Before(toDate); last = last.Add(24 * time.Hour) {
		dates = append(dates, last)
	}

	dates = append(dates, toDate)

	wp.Dates = dates
}

func (wp *weatherAPIParser) ShowWeather() []Weather {
	var result []Weather

	for _, date := range wp.Dates {
		var endpoint string

		difference := math.Ceil(time.Until(date).Hours() / 24)

		if difference > 14 {
			endpoint = "future.json"
		} else if difference > 0 {
			endpoint = "forecast.json"
		} else {
			endpoint = "history.json"
		}

		req, err := http.NewRequest(http.MethodGet, wp.BaseURL+endpoint, nil)

		if err != nil {
			return nil
		}

		q := req.URL.Query()

		q.Add("key", wp.Token)
		q.Add("q", fmt.Sprintf("%.6f,%.6f", wp.Location.Lat, wp.Location.Lng))
		q.Add("lang", "ru")
		q.Add("dt", date.Format("2006-01-02"))

		req.URL.RawQuery = q.Encode()

		client := &http.Client{}

		res, err := client.Do(req)

		if err != nil || res.StatusCode != 200 {
			return nil
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)

		if err != nil {
			return nil
		}

		var wpResponse weatherAPIResponse

		if err := json.Unmarshal(body, &wpResponse); err != nil {
			continue
		}

		if len(wpResponse.Forecast.ForecastDays) > 0 {
			result = append(result, Weather{
				Date:      wpResponse.Forecast.ForecastDays[0].Date,
				MaxTemp:   wpResponse.Forecast.ForecastDays[0].Day.MaxTemp,
				MinTemp:   wpResponse.Forecast.ForecastDays[0].Day.MinTemp,
				Condition: wpResponse.Forecast.ForecastDays[0].Day.Condition.Text,
			})
		}
	}

	return result
}

func NewWeatherParser() IWeatherParser {
	return &weatherAPIParser{BaseURL: os.Getenv("WEATHERAPI_URL"), Token: os.Getenv("WEATHERAPI_TOKEN")}
}
