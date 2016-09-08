package main

import (
	"log"
	"testing"

	"net/http"

	geo "github.com/kellydunn/golang-geo"
)

// Test the README case
func TestContains(t *testing.T) {
	app := Application{}
	app.Load(STATES_FILE)

	ContainingState, longitude, latitude := "Pennsylvania", -77.036133, 40.513799

	states, _ := app.Map.Contains(latitude, longitude)

	if len(states) == 0 {
		log.Fatal("Failed: No states found.")
	}

	if len(states) > 1 {
		log.Fatal("Only one state expected.")
	}

	state := states[0]

	if state != ContainingState {
		log.Fatal("Failed: Did not find PA.")
	}
}

// Test a point out of known bounds
func TestDoesntContain(t *testing.T) {
	app := Application{}
	app.Load(STATES_FILE)

	longitude, latitude := -9999.99, 9999.99

	_, err := app.Map.Contains(latitude, longitude)

	if err == nil {
		log.Fatal("Failed: There's nothing at (-9999.99, 9999.99)!")
	}

}

// Test a point on two state lines
func TestContainsOnStateLine(t *testing.T) {
	app := Application{}
	app.Load(STATES_FILE)

	PA, WV := "Pennsylvania", "West Virginia"
	y1, x1, y2, x2 := -77.475793, 39.719623, -80.524269, 39.721209

	p1, p2 := geo.NewPoint(x1, y1), geo.NewPoint(x2, y2)
	mid := p1.MidpointTo(p2)

	states, _ := app.Map.Contains(mid.Lat(), mid.Lng())

	if len(states) == 0 {
		log.Fatal("Failed: No states found.")
	}

	if len(states) != 2 {
		log.Fatal("Failed: Expected PA and WV")
	}

	if states[0] != PA {
		log.Fatal("Failed: Did not detect that point is on State Line")
	}

	if states[1] != WV {
		log.Fatal("Failed: Did not detect that point is on State Line")
	}
}

// Test the web call (200)
func TestPost200(t *testing.T) {
	go runWebServer()

	resp, _ := http.Post("http://localhost:8080/?longitude=-77.036133&latitude=40.513799", "application/json", nil)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Unexpected Status Code: ", resp.StatusCode)
	}
}

// Test a HTTP Request the server doesn't have a route for
func TestGet404(t *testing.T) {
	go runWebServer()

	resp, _ := http.Get("http://localhost:8080/DNE")

	if resp.StatusCode != http.StatusNotFound {
		log.Fatal("Unexpected Status Code: ", resp.StatusCode)
	}
}

// Test a bad point request, check status code
func TestPost202(t *testing.T) {
	go runWebServer()

	resp, _ := http.Post("http://localhost:8080/?longitude=-9999.99&latitude=-9999.99", "application/json", nil)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		log.Fatal("Unexpected Status Code: ", resp.StatusCode)
	}
}

// Test a malformed point request, check status code
func TestBadRequest(t *testing.T) {
	go runWebServer()

	resp, _ := http.Post("http://localhost:8080/?longitude='hey'", "application/json", nil)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		log.Fatal("Unexpected Status Code: ", resp.StatusCode)
	}
}

// TODO: move to @before
func runWebServer() {
	app := Application{}
	app.Load(STATES_FILE)
	app.Run()
}
