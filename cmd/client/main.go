package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/farkow/co2e/pkg/api"
)

func main() {
	// get configuration
	address := flag.String("server", "localhost:8080", "gRPC server in format host:port")
	start := flag.String("start", "Munich", "Start city")
	end := flag.String("end", "Berlin", "Destination city")
	tm := flag.String("tm", "", "Transportation method")
	tmLong := flag.String("transportation-method", "", "Transportation method (long)")
	flag.Parse()

	// validate parameters
	// default value of the transportation method
	tmVal := ""
	if len(*tm) == 0 && len(*tmLong) == 0 {
		tmVal = "medium-diesel-car"
	} else if len(*tm) == 0 && len(*tmLong) != 0 {
		tmVal = *tmLong
	} else if len(*tm) != 0 && len(*tmLong) == 0 {
		tmVal = *tm
	} else {
		log.Fatalf("Unexpected tranportation parameter failure.")
	}
	// TODO add more validation

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := api.NewServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Calculate
	req := api.Request{
		Start:                *start,
		End:                  *end,
		TransportationMethod: tmVal,
	}
	res, err := c.Calculate(ctx, &req)
	if err != nil {
		log.Fatalf("Calculation failed: %v", err)
	}
	log.Printf("RPC Response: <%+v>\n", res)
	if len(res.Error) == 0 {
		log.Printf("CO2 Emission Result: %.2f", res.Co2E)
	} else {
		log.Printf("An error occured: %s", res.Error)
	}
}
