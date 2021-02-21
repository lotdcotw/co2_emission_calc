package ors

import "testing"

var coords = [2][2]float64{
	{9.70093, 48.477473},
	{9.207916, 49.153868},
}

func TestTime(t *testing.T) {
	initHelper()
	resp := Time(coords[0], coords[1])
	if len(resp.Distances) == 0 {
		t.Errorf("Failed to communicate with Openrouteservice properly. Check and validate your token.")
	}
}
