package ors

const (
	tokEnv = "ORS_TOKEN"
)

// preferably, this part should be in a JSON config file
const (
	apiURL   = "https://api.openrouteservice.org/"
	epSearch = "geocode/search"
	epTime   = "v2/matrix/driving-car"
)

var token = ""
