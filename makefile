http:
	server http

proto-generate:
	protoc  --go_out=ui/grpc/. --go-grpc_out=ui/grpc/. ui/grpc/*.proto