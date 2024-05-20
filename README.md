
# Go Meat Challenges

This repository contains several Go modules designed to solve specific challenges related to string encoding/decoding, path summation in a triangle array, and an API for counting meat types from text data. Each module is independently testable and maintainable.

## Project Structure

```
Go Meat Challenges/
│
├── cmd/
│ ├── encode_decode_numbers/
│ │ ├── encode_decode_numbers.go # Main application logic
│ │ └── encode_decode_numbers_test.go # Test cases for the application
│ ├── max_path_sum/
│ │ ├── files/
│ │ │ └── hard.json # Triangle data for path summation
│ │ ├── max_path_sum.go # Main application logic
│ │ └── max_path_sum_test.go # Test cases for the application
│ └── meat_counter_api/
│ ├── meatcounter/
│ │ ├── meat_counter_grpc.pb.go # GRPC generated file
│ │ └── meat_counter.pb.go # Protobuf generated file
│ ├── proto/
│ │ └── meat_counter.proto # Protobuf definition file
│ ├── meat_counter_api.go # Main application logic for the API
│ └── meat_counter_api_test.go # Test cases for the API
│
├── go.mod # Go modules file
├── go.sum # Go checksum file
├── LICENSE # The license file for the project
└── README.md # This file
```

## Technologies and Libraries

- **Go**: Primary programming language.
- **gRPC**: For the meat counter API.
- **Protobuf**: Used for defining gRPC services and generating Go files.

## Requirements

To run this project, you will need:

- Go (version 1.15 or higher recommended)
- Protobuf Compiler (for generating Go files from `.proto` files, if changes are made to Protobuf definitions)

## Running Tests

Each module can be tested independently by navigating to the corresponding directory under `cmd/` and running the Go test command. Here are the commands to run tests for each module:

### Encode-Decode Numbers

```bash
cd cmd/encode_decode_numbers
go test
```

### Maximum Path Sum

```bash
cd cmd/max_path_sum
go test
```

### Meat Counter API

Before running tests for the Meat Counter API, ensure you have generated the necessary Protobuf files.

```bash
cd cmd/meat_counter_api
go test
```

### Test All
```bash
go test ./cmd/encode_decode_numbers
go test ./cmd/max_path_sum
go test ./cmd/meat_counter_api
```

For regenerating Protobuf files, use:

```bash
protoc --go_out=. --go-grpc_out=. proto/meat_counter.proto
```
