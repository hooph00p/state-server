package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Map struct {
	States []State
}

/**
 * Open Data File,
 * Scan file, unmarshal from JSON into State struct
 * for each line.
 */
func (m *Map) LoadStates(path string) (err error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		state := newState(scanner.Text())
		m.States = append(m.States, *state)
	}
	return
}

/**
 * Create a new State, then generate the geo.Polygon
 * for that state (so we can just call s.Bounds.Contains)
 */
func newState(line string) *State {
	state := State{}
	json.Unmarshal([]byte(line), &state)
	state.toPolygon()
	return &state
}

/**
 * Using golang-geo, check to see if the map contains a point. If it does, return
 * the State.
 */
func (m *Map) Contains(latitude, longitude float64) (s *State, err error) {
	for i := range m.States {
		if m.States[i].Contains(latitude, longitude) {
			s = &m.States[i]
			return
		}
	}
	err = fmt.Errorf("Point not found.")
	return
}
