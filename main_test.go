package main

import (
	"log"
	"testing"

	"net/http"

	geo "github.com/kellydunn/golang-geo"
)

func TestContains(t *testing.T) {
	app := Application{}
	app.Load(STATES_FILE)

	ContainingState, longitude, latitude := "Pennsylvania", -77.036133, 40.513799

	state, _ := app.Map.Contains(latitude, longitude)

	if state.Name != ContainingState {
		log.Fatal("Failed: States not the same.")
	}
}

func TestDoesntContain(t *testing.T) {
	app := Application{}
	app.Load(STATES_FILE)

	longitude, latitude := -9999.99, 9999.99

	_, err := app.Map.Contains(latitude, longitude)

	if err == nil {
		log.Fatal("Failed: There's nothing at (-9999.99, 9999.99)!")
	}

}

func TestContainsOnStateLine(t *testing.T) {
	app := Application{}
	app.Load(STATES_FILE)

	ContainingState, y1, x1, y2, x2 := "Pennsylvania", -77.475793, 39.719623, -80.524269, 39.721209 // Penn JSON

	p1, p2 := geo.NewPoint(x1, y1), geo.NewPoint(x2, y2)
	mid := p1.MidpointTo(p2)

	state, _ := app.Map.Contains(mid.Lat(), mid.Lng())

	if state.Name != ContainingState {
		log.Fatal("Failed: Did not detect that point is on State Line")
	}
}

func TestBorder(t *testing.T) {
	// TODO: Test a spot that rests exactly on two state borders.
	// {"state": "New Hampshire", "border": [[-72.279917, 42.720467], [-71.087509, 45.301469], [-70.81388, 42.867065], [-72.279917, 42.720467]]}
	// {"state": "Massachusetts", "border": [[-71.319328, 41.772195], [-73.49884, 42.07746], [-70.898111, 42.886877], [-69.91778, 41.767653], [-71.319328, 41.772195]]}
}

func TestPost(t *testing.T) {
	go runWebServer()

	resp, _ := http.Post("http://localhost:8080/?longitude=-77.036133&latitude=40.513799", "application/json", nil)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Unexpected Status Code: ", resp.StatusCode)
	}
}

func Test404(t *testing.T) {
	go runWebServer()

	resp, _ := http.Get("http://localhost:8080/DNE")

	if resp.StatusCode != http.StatusNotFound {
		log.Fatal("Unexpected Status Code: ", resp.StatusCode)
	}
}

func Test202(t *testing.T) {
	go runWebServer()

	resp, _ := http.Post("http://localhost:8080/?longitude=-9999.99&latitude=-9999.99", "application/json", nil)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		log.Fatal("Unexpected Status Code: ", resp.StatusCode)
	}
}

func TestBadRequest(t *testing.T) {
	go runWebServer()

	resp, _ := http.Post("http://localhost:8080/?longitude='hey'", "application/json", nil)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		log.Fatal("Unexpected Status Code: ", resp.StatusCode)
	}
}

func runWebServer() {
	app := Application{}
	app.Load(STATES_FILE)
	app.Run()
}
