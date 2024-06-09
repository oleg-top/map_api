package venueParser

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/codingsince1985/geo-golang/openstreetmap"
)

type IVenueParser interface {
	SetLocation(string)
	ShowVenues() []Venue
}

type LatLng struct {
	Lat float64
	Lng float64
}

func (ll *LatLng) String() string {
	return fmt.Sprintf("%.6f,%.6f", ll.Lat, ll.Lng)
}

type VenueLocation struct {
	FormattedAddress []string `json:"formattedAddress"`
}

type Venue struct {
	Name     string        `json:"name"`
	Location VenueLocation `json:"location"`
}

type fsqVenueResponse struct {
	Venues []Venue `json:"venues"`
}

type fsqVenueAnswer struct {
	Response fsqVenueResponse `json:"response"`
}

type fsqVenueParser struct {
	Version    string
	Location   LatLng
	Limit      int
	BaseURL    string
	OAuthToken string
}

func (fsq *fsqVenueParser) SetLocation(location string) {
	loc, _ := openstreetmap.Geocoder().Geocode(location)

	if loc == nil {
		return
	}

	fsq.Location = LatLng{Lat: loc.Lat, Lng: loc.Lng}
}

func (fsq *fsqVenueParser) ShowVenues() []Venue {
	req, err := http.NewRequest(http.MethodGet, fsq.BaseURL, nil)

	if err != nil {
		return nil
	}

	q := req.URL.Query()

	q.Add("ll", fsq.Location.String())
	q.Add("limit", strconv.Itoa(fsq.Limit))
	q.Add("oauth_token", fsq.OAuthToken)
	q.Add("v", fsq.Version)

	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

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

	var answer fsqVenueAnswer

	if err := json.Unmarshal(body, &answer); err != nil {
		return nil
	}

	return answer.Response.Venues
}

func NewVenueParser() IVenueParser {
	version := os.Getenv("FOURSQUARE_VERSION_VENUES")
	limit := 3
	baseURL := os.Getenv("FOURSQUARE_URL_VENUES")
	oAuthToken := os.Getenv("FOURSQUARE_ACCESS_TOKEN")
	return &fsqVenueParser{
		Version:    version,
		Limit:      limit,
		BaseURL:    baseURL,
		OAuthToken: oAuthToken,
	}
}
