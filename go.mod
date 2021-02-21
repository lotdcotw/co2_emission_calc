module github.com/farkow/co2e

go 1.14

replace github.com/farkow/etcd => ./third_party/etcd

require (
	github.com/bndr/gopencils v0.0.0-20161113114152-22e283ad7611
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/farkow/etcd v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/joho/godotenv v1.3.0
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	go.uber.org/zap v1.15.0
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.24.0
	gopkg.in/ini.v1 v1.57.0
)
