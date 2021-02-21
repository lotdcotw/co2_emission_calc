package ors

// TimeRequest contains the payload for ORS matrix call
type TimeRequest struct {
	Locations [2][2]float64 `json:"locations"` // TODO do not limit to 2x2
	Metrics   []string      `json:"metrics"`
	Units     string        `json:"units"`
}

// TimeResponse is a response that has only the required piece of ORS matrix call
type TimeResponse struct {
	Distances [][]float64 `json:"distances"`
}

// Time returns the time between two locations in seconds
func Time(start [2]float64, end [2]float64) TimeResponse {
	var resp TimeResponse
	req := TimeRequest{
		Locations: [2][2]float64{
			start,
			end,
		}, // TODO do not limit to 2
		Metrics: []string{"distance"}, // TODO parameterize
		Units:   "km",
	}

	// convert struct to apipencil payload param
	payload := map[string]interface{}{
		"locations": req.Locations,
		"metrics":   req.Metrics,
		"units":     req.Units,
	}

	// make post request
	r := api.Res(epTime, &resp)
	r.SetHeader("Authorization", token) // preferably, in a middleware
	_, err := r.Post(payload)
	if err != nil {
		return TimeResponse{}
	}
	return resp
}
