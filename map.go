package main

import (
	"bufio"
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
 * Using golang-geo, check to see if the map contains a point. If it does, return
 * the State.
 */
func (m *Map) Contains(latitude, longitude float64) (states []string, err error) {
	for i := range m.States {
		if m.States[i].Contains(latitude, longitude) {
			states = append(states, m.States[i].Name)
		}
	}
	if len(states) == 0 {
		err = fmt.Errorf("No State Found.")
	}
	return
}
