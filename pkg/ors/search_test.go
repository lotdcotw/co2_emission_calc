package ors

import "testing"

func TestSearch(t *testing.T) {
	initHelper()
	resp := Search("Munich")
	if len(resp.Features) == 0 {
		t.Errorf("Failed to communicate with Openrouteservice properly. Check and validate your token.")
	}
}
