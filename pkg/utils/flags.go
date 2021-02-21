package utils

import (
	"flag"
	"log"
	"testing"
)

// Flags keeps all service flags
var Flags flapp

type flapp struct {
	GRPCPort string
	EtcdPort string

	LogLevel      int    // LogLevel is global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
	LogTimeFormat string // LogTimeFormat is print time format for logger e.g. 2006-01-02T15:04:05Z07:00

	EnvFile string
}

func flags() {
	flag.StringVar(&Flags.GRPCPort, "grpc-port", "8080", "gRPC port to bind")
	flag.StringVar(&Flags.EtcdPort, "etcd-port", "2379", "etcd port")

	flag.IntVar(&Flags.LogLevel, "log-level", 0, "Global log level")
	flag.StringVar(&Flags.LogTimeFormat, "log-time-format", "", "Print time format for logger e.g. 2006-01-02T15:04:05Z07:00")

	flag.StringVar(&Flags.EnvFile, "env-file", "./env", "Environment file")

	var _ = func() bool {
		testing.Init()
		return true
	}()

	flag.Parse()
	log.Printf("Flags are parsed.\n")
}
