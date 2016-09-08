package main

import (
	"encoding/json"

	geo "github.com/kellydunn/golang-geo"
)

const (
	LATITUDE  int = 1 // position in the 2D Array
	LONGITUDE     = 0
)

type State struct {
	Name   string      `json:"state"`
	Border [][]float64 `json:"border"`
	Bounds *geo.Polygon
}

/**
 * Create a new State by unmarshalling into JSON,
 * then generate the geo.Polygon
 * for that state (so we can just call s.Bounds.Contains)
 */
func newState(line string) *State {
	state := State{}
	json.Unmarshal([]byte(line), &state)
	state.toPolygon()
	return &state
}

/**
 * Migrate to a golang-geo Polygon for ease of use.
 */
func (s *State) toPolygon() {
	var points []*geo.Point
	for i := range s.Border {
		latitude := s.Border[i][LATITUDE]
		longitude := s.Border[i][LONGITUDE]
		point := geo.NewPoint(latitude, longitude)
		points = append(points, point)
	}
	s.Bounds = geo.NewPolygon(points)
}

/**
 * Using golang-geo, check to see if the state polygon contains the point.
 */
func (s *State) Contains(latitude, longitude float64) bool {
	p := geo.NewPoint(latitude, longitude)
	return s.Bounds.Contains(p)
}
