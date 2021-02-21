package ors

// only needed ORS geolocation structures
type properties struct {
	Layer    string `json:"layer"`
	Name     string `json:"name"`
	Accuracy string `json:"accuracy"`
}
type geometry struct {
	Coordinates []float64 `json:"coordinates"`
}
type feature struct {
	Properties properties `json:"properties"`
	Geometry   geometry   `json:"geometry"`
}

// SearchResponse is a response that has only the required piece of ORS geo location call
type SearchResponse struct {
	Features []feature  `json:"features"`
	Point    int        `json:"-"`
	Coords   [2]float64 `json:"-"`
}

// Search returns coordinates of the given location
func Search(location string) SearchResponse {
	var resp SearchResponse
	_, err := api.Res(epSearch, &resp).Get(map[string]string{
		"api_key": token, // preferably, should be in a middleware
		"text":    location,
	})
	if err != nil {
		return SearchResponse{}
	}
	return resp
}

// Filter returns the requested point
func Filter(locations []SearchResponse, point int) SearchResponse {
	for _, v := range locations {
		if v.Point == point {
			return v
		}
	}
	return SearchResponse{}
}
