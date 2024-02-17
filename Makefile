stubs:
	protoc -I ./api --go_out ./internal/infra/grpc --go_opt paths=source_relative --go-grpc_out ./internal/infra/grpc --go-grpc_opt paths=source_relative ./api/payment-processing.proto