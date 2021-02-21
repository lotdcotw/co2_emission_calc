# CO2 Emission Calculator

## Introduction
This repo calculates CO2 emission between two cities depending on different vehicles.    

Server application can be communicated with any client via rpc calls.  
Client application directly communicates with the server with the given parameters.  
So that,
- The server is reusable and new clients can be developed independently.  
- The server handles the calculation. Every client will be synced with any data or methodolgy change.
- A user interface can be developed easily on top of it independently.

You can quickly start with 'Quick Start' section below.
However, make sure you got all requirements e.g. Go.

### Quick Start with Docker
1. Extract ZIP file
2. cd into the folder
3. Create .env file from .env.example
4. Put your ORS token
5. ```make reqs```
6. ```make docker```

### Quick Start without Docker
1. Extract ZIP file
2. cd into the folder
3. Create .env file from .env.example
4. Put your ORS token
5. ```make reqs```
6. ```make build```
7. ```make test```
8. ```make run-server```
9. ```make run-client```

### How-to-Use
1. Run the server or build & run via docker-compose
2. Run the client with parameters

The client application requires 3 parameters;
- ```start``` as start point (```default```: Munich)
- ```end``` as destination point (```default```: Berlin)
- ```tm``` or ```transportation-method``` as transportation method (```default```: medium-diesel-car)
- ```server``` as the server address (```default```: localhost:8080)

```
./bin/client -start Hamburg -end Berlin -tm medium-car-diesel
```

**List of Transportation Methods**
-	small-diesel-car
- small-petrol-car
- small-plugin-hybrid-car
- small-electric-car
- medium-diesel-car
- medium-petrol-car
- medium-plugin-hybrid-car
- medium-electric-car
- large-diesel-car
- large-petrol-car
- large-plugin-hybrid-car
- large-electric-car
- bus
- train

### Acceptance Criteria Helper
After you run the server, run the following commands to test functional requirements;
```
./test/ac.sh
```
  
## Project Information
Languages:
- Go
- Shell
- Makefile

Standards:
- Standard Go project layout is used: https://github.com/golang-standards/project-layout  
  
Storage (_as key-value_):
- ```etcd``` (is a strongly consistent, distributed key-value store that provides a reliable way to store data that needs to be accessed by a distributed system)
There is initial data for CO2e per passenger per km in different transportation methods.  
Still possible to use id without etcd.  
Static data file is under ```data``` folder.  
  
CI/CD:
- Docker
- Docker Compose

Package Management:
- Go Modules (```go mod```)

Third-party Packages and Services:
- ```protof``` to generate RPC service code automatically
- ```protoc-gen-go``` as protocol compiler plugin for Go
- Uber's ```zap``` as logging middleware
- ```Openrouteservice``` as Geocode and distance API
- ```gopencils``` to make API calls quickly
- ```godotenv``` to read environment variables
- ```etcd``` as key-value storage  
For version >=3, original version github.com/etcd-io/etcd is having dependency problems.  
I have a fork in my account github.com/farkow/etcd to make it run until the problem is resolved / the solution is merged.
- ```gopkg.in/ini``` to read INI file in-case-of not using ```etcd```
- Average emission values: https://www.gov.uk/government/publications/greenhouse-gas-reporting-conversion-factors-2019


## Tech Stack
- go
- protoc
- etcd
- docker & docker-compose

## Make

### Configure
- (Optional) Copy .env.example file as .env file and configure your ports
- Create an Openrouteservice account and save your token  
At least one of the method below is required.
  - In ORS_TOKEN environment variable
  - In .env file

### Requirements
Install all modules with the following command;
```
make reqs
```
Normally, git actions should not be there.
Since the project delivery is asked as a ZIP file, this command will handle everything for you.

### Build
All executables will be built under ```bin``` folder with the following command.
```
make build
```
Executables can be called independently to test the application.  
You are **not** obliged to build docker containers.

```pkg/api/co2e.pb.go``` file is not removed so you can skip auto-code generation and its prerequisites like protoc-go-gen executable in PATH.  

### Proto
Run the following command to auto generate service code from proto file under ```api/proto```
```
make proto
```

### Run Server
Use the following command and parameters to run the server.
```
make run-server
```

### Run Client
Use the following command and parameters to run client with example parameters.
```
make run-client
```
To run the client with your parameters;
```
./bin/client --start Hamburg --end Berlin --transportation-method medium-car-diesel
```

### Tests
Fist, please make sure you have OS_TOKEN in your PATH for testing purposes.  
Call the following commands to run all tests and create a coverate report under ```data```.
```
make test
``` 

### Local Deployment
The following softwares are needed;
- docker
- docker compose

The following command will build and run all containers.
```
make deploy-local
```

### Platforms Builds
```
make linux
make win
make mac
```
For more specific builds, you can visit https://github.com/golang/go/blob/master/src/go/build/syslist.go to see all possible options for GOOS and GOARCH.


### Client
This project also includes a client to be able to test the service quickly.  
Run the following command to find out CO2E value from Hamburg to Berlin.
```
make client
```

## Notes for Developers / Reviewers
- Because of time limitations;
  - I did not parameterize "distance" and "units" in ORS API calls
  - I have limited ORS matrix to 2x2
  - I only spend time on really necessary tests
  - I used some constant values for ORS API e.g. address and endpoints
  - I did not use clusters for etcd, only localhost setup, fixed address
  - I did not validate the input
  - I limited the location search results to first city only
  - I did not work on Jenkins or Travis scripts

### Example output of Docker & Client
The following output is taken from my home machine;
```
make docker
docker-compose down
Stopping co2e_etcd_1   ... done
Stopping co2e_server_1 ... done
Removing co2e_etcd_1   ... done
Removing co2e_server_1 ... done
Removing network co2e_co2e
Removing network co2e_default
docker-compose build
etcd uses an image, skipping
Building server
Step 1/12 : FROM golang:1.14.3 AS builder
 ---> 7e5e8028e8ec
Step 2/12 : ADD . /co2e
 ---> baf2d3d38bb2
Step 3/12 : WORKDIR /co2e
 ---> Running in 926d4779b58b
Removing intermediate container 926d4779b58b
 ---> e56774a609fd
Step 4/12 : RUN go mod download && go mod vendor
 ---> Running in c1a3f3f74db7
Removing intermediate container c1a3f3f74db7
 ---> 7203060fa372
Step 5/12 : RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/co2e
 ---> Running in f1714dfd47fa
Removing intermediate container f1714dfd47fa
 ---> f63cee2853f1
Step 6/12 : FROM alpine:latest
 ---> f70734b6a266
Step 7/12 : RUN apk --no-cache add ca-certificates
 ---> Using cache
 ---> d094a1f068e0
Step 8/12 : WORKDIR /root/
 ---> Using cache
 ---> 07808275e112
Step 9/12 : COPY --from=builder /co2e/.env .
 ---> Using cache
 ---> 2ed1d95b6929
Step 10/12 : COPY --from=builder /co2e/data/co2e.ini /root/data/co2e.ini
 ---> 95a050085823
Step 11/12 : COPY --from=builder /co2e/app .
 ---> 6fe516bf73eb
Step 12/12 : CMD ["./app"]
 ---> Running in 88e343460416
Removing intermediate container 88e343460416
 ---> cd77980a096e
Successfully built cd77980a096e
Successfully tagged co2e_server:latest
docker-compose up -d
Creating network "co2e_co2e" with driver "bridge"
Creating network "co2e_default" with the default driver
Creating co2e_server_1 ... done
Creating co2e_etcd_1   ... done
scripts/etcd.sh
etcd Version: 3.4.9
Git SHA: 54ba95891
Go Version: go1.12.17
Go OS/Arch: linux/amd64
etcdctl version: 3.4.9
API version: 3.4
127.0.0.1:2379 is healthy: successfully committed proposal: took = 1.313033544s
carsmall carmedium carlarge bus train
INI__carsmall__diesel INI__carsmall__petrol INI__carsmall__pluginhybrid INI__carsmall__electric INI__carmedium__diesel INI__carmedium__petrol INI__carmedium__pluginhybrid INI__carmedium__electric INI__carlarge__diesel INI__carlarge__petrol INI__carlarge__pluginhybrid INI__carlarge__electric INI__bus__generic INI__train__generic
Storing variables in carsmall:
carsmall_diesel = 142
OK
carsmall_electric = 50
OK
carsmall_petrol = 154
OK
carsmall_pluginhybrid = 73
OK
Storing variables in carmedium:
carmedium_diesel = 171
OK
carmedium_electric = 58
OK
carmedium_petrol = 192
OK
carmedium_pluginhybrid = 110
OK
Storing variables in carlarge:
carlarge_diesel = 209
OK
carlarge_electric = 73
OK
carlarge_petrol = 282
OK
carlarge_pluginhybrid = 126
OK
Storing variables in bus:
bus_generic = 27
OK
Storing variables in train:
train_generic = 6
OK
farkow@home:~/co2e$ ./bin/client -start Munich -end Berlin -tm medium-diesel-car
2020/05/30 02:59:03 RPC Response: <co2e:100.13>
2020/05/30 02:59:03 CO2 Emission Result: 100.13
farkow@home:~/co2e$ docker ps
CONTAINER ID        IMAGE               COMMAND                 CREATED             STATUS              PORTS                              NAMES
4a58816198a7        co2e_server         "./app"                 19 seconds ago      Up 15 seconds       0.0.0.0:8080->8080/tcp             co2e_server_1
86eb05a18081        bitnami/etcd:3      "/entrypoint.sh etcd"   19 seconds ago      Up 15 seconds       0.0.0.0:2379-2380->2379-2380/tcp   co2e_etcd_1
farkow@home:~/co2e$ docker logs 4a58816198a7
2020/05/30 00:58:54 Flags are parsed.
.env file is not found, continuing on with default and environment values...
{"level":"warn","ts":"2020-05-30T00:58:57.201Z","caller":"clientv3/retry_interceptor.go:62","msg":"retrying of unary invoker failed","target":"endpoint://client-888bba67-800d-46b8-b22f-efd8a9dd6c89/localhost:2379","attempt":0,"error":"rpc error: code = DeadlineExceeded desc = latest balancer error: all SubConns are in TransientFailure, latest connection error: connection error: desc = \"transport: Error while dialing dial tcp 127.0.0.1:2379: connect: connection refused\""}
2020/05/30 00:58:57 [ETCD ERROR] context deadline exceeded
2020/05/30 00:58:57 Loading default emission values...
{"level":"warn","ts":1590800337.2019384,"msg":"time format for logger is not provided - use zap default"}
Running server on GRPC: 8080
{"level":"info","ts":1590800337.2022727,"msg":"Starting gRPC server..."}
{"level":"info","ts":1590800343.2978976,"msg":"Start: [11.544467 48.152126]"}
{"level":"info","ts":1590800343.2979543,"msg":"Destination: [13.40732 52.52045]"}
{"level":"info","ts":1590800343.3807096,"msg":"Distance: 585.53"}
{"level":"info","ts":1590800343.3807511,"msg":"Transportation Method: carmedium_diesel"}
{"level":"info","ts":1590800343.3807652,"msg":"CO2e Value: 171"}
{"level":"info","ts":1590800343.380776,"msg":"Emission: 100.13"}
{"level":"info","ts":1590800343.3808115,"msg":"finished unary call with code OK","grpc.start_time":"2020-05-30T00:59:02Z","grpc.request.deadline":"2020-05-30T00:59:07Z","system":"grpc","span.kind":"server","grpc.service":"api.Service","grpc.method":"Calculate","peer.address":"192.168.192.1:42764","grpc.code":"OK","grpc.time_ms":400.8630065917969}
{"level":"info","ts":1590800343.3835776,"msg":"transport: loopyWriter.run returning. connection error: desc = \"transport is closing\"","system":"grpc","grpc_log":true}
farkow@home:~/co2e$ ./bin/client -start Munich -end Berlin -tm large-diesel-car
2020/05/30 02:59:39 RPC Response: <co2e:122.38>
2020/05/30 02:59:39 CO2 Emission Result: 122.38
farkow@home:~/co2e$ 
```
