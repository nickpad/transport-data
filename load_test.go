package transportdata

import (
	"testing"
)

func TestLoadStops(t *testing.T) {
	database := LoadStops("data/stops.txt")
	stopName := "Milsons Point Wharf"

	if database[stopName].Name != stopName {
		t.Errorf("Unexpected stop name: %#v", stopName)
	}
}
