
proto:
	protoc --proto_path=api/proto --go_out=plugins=grpc:pkg/api co2e.proto

reqs:
	git init
	git add .gitignore
	git add .
	git commit -m "Initial"
	git submodule init
	git submodule add https://github.com/farkow/etcd.git third_party/etcd
	git submodule update --recursive
	git submodule sync
	go mod vendor
	go mod verify

build:
	go mod vendor
	go build -o ./bin/ ./cmd/co2e
	go build -o ./bin/ ./cmd/client

run-server: build
	./bin/co2e -env-file ./.env

run-client: build
	./bin/client -start Hamburg -end Berlin -tm medium-diesel-car

test:
	go test ./... -coverprofile=./data/c.out
	go tool cover -html=./data/c.out

clean:
	rm ./bin/*

docker: docker-clean
	docker-compose build
	docker-compose up -d
	scripts/etcd.sh

docker-server:
	docker build -t farkow/co2e:latest .

docker-clean:
	docker-compose down

linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/co2e ./cmd/co2e

win:
	GOOS=windows GOARCH=amd64 go build -o ./bin/co2e.win.exe ./cmd/co2e

mac:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/co2e.mac ./cmd/co2e
