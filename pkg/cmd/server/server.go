package server

import (
	"context"
	"fmt"

	"github.com/farkow/co2e/pkg/logger"
	"github.com/farkow/co2e/pkg/protocol/grpc"
	"github.com/farkow/co2e/pkg/service"
	"github.com/farkow/co2e/pkg/utils"
)

// RunServer runs gRPC server
func RunServer() error {
	ctx := context.Background()

	// initialize logger
	if err := logger.Init(utils.Flags.LogLevel, utils.Flags.LogTimeFormat); err != nil {
		return fmt.Errorf("Failed to initialize logger: %v", err)
	}

	fmt.Printf("Running server on GRPC: %s\n", utils.Flags.GRPCPort)

	api := service.NewServiceServer()

	return grpc.RunServer(ctx, api, utils.Flags.GRPCPort)
}
