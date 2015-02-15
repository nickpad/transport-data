package graph

import "testing"

func TestGraphStopNode(t *testing.T) {
	stop := StopNode{"Central"}
	if stop.Name != "Central" {
		t.Errorf("Expected Central but got: %#v", stop.Name)
	}
}
