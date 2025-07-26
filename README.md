### Car Parts Shop API
This is a simple RESTful API that provides a list of car parts and their prices.

### Prerequisites

- Go 1.23.7 
- Gin 1.10.1
- Gorm 2.2.46
- Air 1.41.0
- Mockery 1.1.5
- Protobuf 3.21.12
- Protobuf-go-grpc 1.3.0
- gRPC 1.56.1
- gRPC-go 1.6.0

### Usage

1.Install dependencies
```bash
go mod download
```

2.Build application
```bash
make build
```

3.You can use air to run the application
```
air
```

## Execute test

1.Use go test to execute tests
```
go test -v ./...
```