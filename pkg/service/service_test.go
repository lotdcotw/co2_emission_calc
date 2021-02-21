package service

import (
	"testing"

	"github.com/farkow/co2e/pkg/ors"
)

var (
	testLocation = "Munich"
	testCoord1   = [2]float64{9.70093, 48.477473}
	testCoord2   = [2]float64{9.207916, 49.153868}
)

func TestFindLocation(t *testing.T) {
	c := make(chan ors.SearchResponse)
	go findLocation(testLocation, 0, c)
	l := <-c
	if len(l.Features) == 0 {
		t.Errorf("Failed to search a location via ORS")
	}
}

func TestFindCoordinates(t *testing.T) {
	l1 := ors.SearchResponse{
		Coords: testCoord1,
	}
	l2 := ors.SearchResponse{
		Coords: testCoord2,
	}
	findCoordinates(&l1, &l2)
}

func TestFindDistance(t *testing.T) {
	km, err := findDistance(testCoord1, testCoord2)
	if err != nil {
		t.Error(err)
	}
	if km == 0 {
		t.Errorf("Unable to calculate the distance via ORS")
	}
}
