package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	envVarGrpcPort  = "GRPC_PORT"
	defaultGrpcPort = "8080"

	envVarEtcdPort  = "ETCD_PORT"
	defaultEtcdPort = "2379"
)

func init() {
	flags()

	if len(Flags.EnvFile) != 0 {
		if err := godotenv.Load(Flags.EnvFile); err != nil {
			fmt.Printf(".env file is not found, continuing on with default and environment values...\n")
		}
	}

	// grpc default
	grpcPort := get(envVarGrpcPort)
	if len(grpcPort) == 0 && len(Flags.GRPCPort) != 0 {
		grpcPort = Flags.GRPCPort
		fmt.Printf("Overriding GRPC Port environment variable (%s) with the flag...\n", grpcPort)
	}
	Flags.GRPCPort = grpcPort

	// etcd default
	etcdPort := get(envVarEtcdPort)
	if len(etcdPort) == 0 && len(Flags.EtcdPort) != 0 {
		etcdPort = Flags.EtcdPort
		fmt.Printf("Overriding Etcd Port environment variable (%s) with the flag...\n", etcdPort)
	}
	Flags.EtcdPort = etcdPort
}

func get(key string) string {
	return os.Getenv(key)
}
