package transportdata

import (
	"testing"
)

func TestLoadStops(t *testing.T) {
	database := Database{}
	LoadStops("data/stops.txt", database)
	stopId := "20611"
	expectedName := "Milsons Point Wharf"

	actualName := database[stopId].Name

	if expectedName != actualName {
		t.Errorf("Unexpected stop name: %#v", actualName)
	}
}
