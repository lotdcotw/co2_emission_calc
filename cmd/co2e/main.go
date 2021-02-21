package main

import (
	"fmt"
	"os"

	"github.com/farkow/co2e/pkg/cmd/server"
)

func main() {
	if err := server.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
