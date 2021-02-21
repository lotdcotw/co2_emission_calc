package service

import (
	"context"
	"errors"
	"log"
	"math"

	"github.com/farkow/co2e/pkg/api"
	"github.com/farkow/co2e/pkg/ors"
)

func init() {
	getValues()
}

// serviceServer is implementation of api.ServiceServer proto interface
type serviceServer struct {
}

// NewServiceServer creates CO2E service
func NewServiceServer() api.ServiceServer {
	return &serviceServer{}
}

// Calculate calculates CO2 emission between two distances with the given transport method
func (s *serviceServer) Calculate(ctx context.Context, req *api.Request) (*api.Response, error) {
	// validate the given data
	err := validate(req)
	if err != nil {
		return nil, err
	}

	// get locations coordinates
	c := make(chan ors.SearchResponse)
	go findLocation(req.Start, 0, c) // find start location
	go findLocation(req.End, 1, c)   // find end location
	l1, l2 := <-c, <-c
	start := ors.Filter([]ors.SearchResponse{l1, l2}, 0)
	end := ors.Filter([]ors.SearchResponse{l1, l2}, 1)
	// full route is needed
	if len(start.Features) == 0 || len(end.Features) == 0 {
		return &api.Response{
			Co2E:  0,
			Error: "Given cities are not found. Check your ORS token or your values.",
		}, nil
	}

	// find coordinates
	findCoordinates(&start, &end)
	log.Printf("Start: %v\n", start.Coords)
	log.Printf("Destination: %v\n", end.Coords)

	// find the distance between 2 coordinates
	kms, err := findDistance(start.Coords, end.Coords)
	if err != nil {
		return &api.Response{
			Co2E:  0,
			Error: "Could not calculate the distance. Check your parameters",
		}, nil
	}
	log.Printf("Distance: %v\n", kms)
	log.Printf("Transportation Method: %v\n", req.TransportationMethod)
	log.Printf("CO2e Value: %v\n", emissions[req.TransportationMethod])

	// calculate the emission value in kgs
	emission := kms * emissions[req.TransportationMethod] / 1000 // convert to kgs // TODO add conversion methods
	emission = math.Round(emission*100) / 100                    // 2 decimal points after .

	log.Printf("Emission: %v\n", emission)

	return &api.Response{
		Co2E: float32(emission), // float32 has more range among clients for general public usage, so that api uses float32
	}, nil
}

// find location with a name and location pointer
func findLocation(name string, point int, c chan ors.SearchResponse) {
	result := ors.Search(name)
	result.Point = point
	c <- result
}

// find coordinates in ORS response
func findCoordinates(start *ors.SearchResponse, end *ors.SearchResponse) {
	// get first available city from search results || TODO make it flexible, do not limit search results
	for _, f := range (*start).Features {
		if f.Properties.Layer == "locality" {
			copy((*start).Coords[:], f.Geometry.Coordinates[:2])
			break
		}
	}
	for _, f := range (*end).Features {
		if f.Properties.Layer == "locality" {
			copy((*end).Coords[:], f.Geometry.Coordinates[:2])
			break
		}
	}
}

// find the distance between 2 coordinates
func findDistance(start [2]float64, end [2]float64) (float64, error) {
	// find time between two cities
	seconds := ors.Time(start, end)
	if len(seconds.Distances) != 2 || len(seconds.Distances[0]) != 2 {
		return 0, errors.New("Could not calculate the distance. Check your parameters")
	}

	kms := seconds.Distances[0][1]
	return kms, nil
}
