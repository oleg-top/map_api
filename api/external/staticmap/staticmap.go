package staticmap

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"net/http"
	"os"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/openstreetmap"
	sm "github.com/flopp/go-staticmaps"
	"github.com/fogleman/gg"
	"github.com/golang/geo/s2"
)

type IStaticMapRenderer interface {
	SetSize(int, int)
	SetLocations([]string)
	SetPath(string)
	SaveImage() (string, error)
}

type goSMR struct {
	Color color.RGBA
	Path  string
	Ctx   *sm.Context
}

type geoapifyGeometry struct {
	Coordinates [][][]float64 `json:"coordinates"`
}

type geoapifyObject struct {
	Geometry geoapifyGeometry `json:"geometry"`
}

type geoapifyResponse struct {
	Features []geoapifyObject `json:"features"`
}

func (smr *goSMR) SetSize(width, height int) {
	smr.Ctx.SetSize(width, height)
}

func (smr *goSMR) SetLocations(locations []string) {
	var from, to *geo.Location
	var clrMarker, clrPath color.RGBA

	for i := 0; i < len(locations)-1; i++ {
		from, _ = openstreetmap.Geocoder().Geocode(locations[i])
		to, _ = openstreetmap.Geocoder().Geocode(locations[i+1])

		if from == nil || to == nil {
			continue
		}

		req, err := http.NewRequest(http.MethodGet, os.Getenv("GEOAPIFY_URL"), nil)
		if err != nil {
			continue
		}

		q := req.URL.Query()
		q.Add("apiKey", os.Getenv("GEOAPIFY_TOKEN"))
		q.Add("mode", "drive")
		q.Add("waypoints", fmt.Sprintf("%.6f,%.6f|%.6f,%.6f", from.Lat, from.Lng, to.Lat, to.Lng))
		req.URL.RawQuery = q.Encode()

		client := &http.Client{}

		res, err := client.Do(req)

		if err != nil {
			continue
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)

		if err != nil {
			continue
		}

		var gr geoapifyResponse

		if err := json.Unmarshal(body, &gr); err != nil {
			continue
		}

		var positions []s2.LatLng

		if len(gr.Features) > 0 {
			for _, latLng := range gr.Features[0].Geometry.Coordinates[0] {
				positions = append(positions, s2.LatLngFromDegrees(latLng[1], latLng[0]))
			}
		}

		fmt.Println(positions)

		if i == 0 {
			clrMarker = color.RGBA{0xff, 0xf4, 0x4f, 0xff}
			clrPath = color.RGBA{0x8f, 0, 0xff, 0xff}
		} else {
			clrMarker = color.RGBA{0xff, 0, 0, 0xff}
			clrPath = color.RGBA{0, 0, 0xff, 0xff}
		}

		smr.Ctx.AddObject(
			sm.NewMarker(
				s2.LatLngFromDegrees(from.Lat, from.Lng),
				clrMarker,
				14.0,
			),
		)
		smr.Ctx.AddObject(
			sm.NewMarker(
				s2.LatLngFromDegrees(to.Lat, to.Lng),
				clrMarker,
				14.0,
			),
		)
		smr.Ctx.AddObject(
			sm.NewPath(
				positions,
				clrPath,
				2,
			),
		)
	}
}

func (smr *goSMR) SetPath(path string) {
	smr.Path = path
}

func (smr *goSMR) SaveImage() (string, error) {
	img, err := smr.Ctx.Render()

	if err != nil {
		return "", err
	}

	if err = gg.SavePNG(smr.Path, img); err != nil {
		return "", err
	}

	return smr.Path, nil
}

func NewStaticMapRenderer() IStaticMapRenderer {
	ctx := sm.NewContext()
	return &goSMR{Ctx: ctx}
}
