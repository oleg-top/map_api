package restaurantParser

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/codingsince1985/geo-golang/openstreetmap"
)

type IRestaurantParser interface {
	SetLocation(location string)
	ShowRestaurants() []Restaurant
}

type Restaurant struct {
	Name     string             `json:"name"`
	Location RestaurantLocation `json:"location"`
}

type RestaurantLocation struct {
	FormattedAddress string `json:"formatted_address"`
}

type LatLng struct {
	Lat float64
	Lng float64
}

func (ll *LatLng) String() string {
	return fmt.Sprintf("%.6f,%.6f", ll.Lat, ll.Lng)
}

type fsqResults struct {
	Results []Restaurant `json:"results"`
}

type fsqRestaurantParser struct {
	BaseURL  string
	Location LatLng
	Token    string
	Category string
	Fields   string
	Limit    int
	SortBy   string
}

func (fsq *fsqRestaurantParser) SetLocation(location string) {
	loc, _ := openstreetmap.Geocoder().Geocode(location)

	if loc == nil {
		return
	}

	fsq.Location = LatLng{Lat: loc.Lat, Lng: loc.Lng}
}

func (fsq *fsqRestaurantParser) ShowRestaurants() []Restaurant {
	req, err := http.NewRequest(http.MethodGet, fsq.BaseURL, nil)

	if err != nil {
		return nil
	}

	req.Header.Add("Authorization", fsq.Token)

	q := req.URL.Query()

	q.Add("ll", fsq.Location.String())
	q.Add("categories", fsq.Category)
	q.Add("fields", fsq.Fields)
	q.Add("sort", fsq.SortBy)
	q.Add("limit", strconv.Itoa(fsq.Limit))

	req.URL.RawQuery = q.Encode()

	client := http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil
	}

	var restaurants fsqResults

	if err := json.Unmarshal(body, &restaurants); err != nil {
		return nil
	}

	return restaurants.Results
}

func NewRestaurantParser() IRestaurantParser {
	baseURL := os.Getenv("FOURSQUARE_URL_PLACES")
	token := os.Getenv("FOURSQUARE_TOKEN")
	category := os.Getenv("FOURSQUARE_RESTAURANT_CATEGORY")
	fields := "name,location"
	limit := 5
	sortBy := "DISTANCE"
	return &fsqRestaurantParser{
		BaseURL: baseURL, Token: token, Category: category, Fields: fields, Limit: limit, SortBy: sortBy,
	}
}
