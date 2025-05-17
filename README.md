# gRPC Report Generation Service

This Go service exposes a gRPC API to generate reports for users and provides a health check endpoint. It also runs a cron job every 10 seconds to automatically generate reports for predefined users.

---

## Prerequisites

- Go 1.24 or later installed
- `protoc` (Protocol Buffers compiler) installed and added to your system PATH
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins installed (`go install` commands provided below)
- Git (to clone this repository)

---
_______________________________________________________________________________________________________________________________________________________________________________________________________________________

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/mgsquare/go-report-generation.git
   cd go-report-generation

2. Install protobuf Go plugins (if not installed):

  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

3.Generate the protobuf Go code:

  protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/report.proto

4.Download dependencies:

  go mod tidy

_______________________________________________________________________________________________________________________________________________________________________________________________________________________

## Running the service

Run the server with:

  go run main.go

  You should see logs indicating that the server has started and the cron job is generating reports every 10 seconds.

_______________________________________________________________________________________________________________________________________________________________________________________________________________________

## Testing the service

Using a gRPC Client

You can test the gRPC endpoints with Postman.

Import the provided Postman collection by using the following link
https://.postman.co/workspace/Lets-GOOOO~4c3e869f-3acc-4184-8a25-580dac631161/grpc-request/6828904f59acd04d44086667?action=share&creator=41825570&ctx=documentation

Update the server address to localhost:50051

Use the gRPC requests GenerateReport and HealthCheck from the collection.

for **GenerateReport**

request: {
  "userId": "test-user"
}

response: {
  "reportId": "d62da641-9818-40ab-b4e8-069dbb2cc64f",
  "error": ""
}

for **HealthCheck**

no request is necessary

response: {
"status": "ok" 
}


