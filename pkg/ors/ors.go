package ors

import (
	"os"

	"github.com/bndr/gopencils"
)

// api client
var api *gopencils.Resource

func init() {
	// get API token from env vars
	token = os.Getenv(tokEnv)

	// create ORS api
	api = gopencils.Api(apiURL)
}
