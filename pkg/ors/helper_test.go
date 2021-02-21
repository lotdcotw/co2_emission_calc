package ors

import "os"

func initHelper() {
	token = os.Getenv("ORS_TOKEN")
}
