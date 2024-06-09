package hotelParser

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/codingsince1985/geo-golang/openstreetmap"
)

type IHotelParser interface {
	SetLocation(string)
	SetCheckIn(time.Time)
	SetCheckOut(time.Time)
	SetLimit(int)
	ShowHotels() []Hotel
}

type Hotel struct {
	Name     string
	Stars    int
	Location string
	PriceAVG float64
}

type travelPayoutsHP struct {
	BaseURL  string
	City     string
	CheckIn  string
	Checkout string
	Limit    int
}

type responseGeo struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lon"`
}

type responseLocation struct {
	Country string      `json:"country"`
	Geo     responseGeo `json:"geo"`
	City    string      `json:"name"`
}

type responseHotel struct {
	Location  responseLocation `json:"location"`
	PriceAVG  float64          `json:"priceAvg"`
	HotelName string           `json:"hotelName"`
	Stars     int              `json:"stars"`
}

func (tphp *travelPayoutsHP) SetLocation(location string) {
	loc, _ := openstreetmap.Geocoder().Geocode(location)

	if loc == nil {
		return
	}

	addr, _ := openstreetmap.Geocoder().ReverseGeocode(loc.Lat, loc.Lng)

	if addr == nil {
		return
	}

	tphp.City = addr.City
}

func (tphp *travelPayoutsHP) SetCheckIn(checkIn time.Time) {
	tphp.CheckIn = checkIn.Format("2006-01-02")
}

func (tphp *travelPayoutsHP) SetCheckOut(checkOut time.Time) {
	tphp.Checkout = checkOut.Format("2006-01-02")
}

func (tphp *travelPayoutsHP) SetLimit(limit int) {
	tphp.Limit = limit
}

func (tphp *travelPayoutsHP) ShowHotels() []Hotel {
	req, err := http.NewRequest(http.MethodGet, tphp.BaseURL, nil)

	if err != nil {
		return nil
	}

	q := req.URL.Query()

	q.Add("location", tphp.City)
	q.Add("currency", "rub")
	q.Add("checkIn", tphp.CheckIn)
	q.Add("checkOut", tphp.Checkout)
	q.Add("limit", strconv.Itoa(tphp.Limit))

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

	var resHotel []responseHotel

	if err := json.Unmarshal(body, &resHotel); err != nil {
		return nil
	}

	hotels := make([]Hotel, 0)

	for _, rh := range resHotel {
		hotel := Hotel{
			Name:     rh.HotelName,
			Stars:    rh.Stars,
			PriceAVG: rh.PriceAVG,
		}

		addr, _ := openstreetmap.Geocoder().ReverseGeocode(rh.Location.Geo.Lat, rh.Location.Geo.Lng)

		if addr == nil {
			hotel.Location = rh.Location.City
		} else {
			hotel.Location = addr.FormattedAddress
		}

		hotels = append(hotels, hotel)
	}

	return hotels
}

func NewHotelParser() IHotelParser {
	return &travelPayoutsHP{BaseURL: os.Getenv("TRAVELPAYOUTS_URL")}
}
