proto-generate:
	protoc  --go_out=ui/grpc/. --go-grpc_out=ui/grpc/. ui/grpc/*.proto

test-http:
	k6 run test-http.js

test-grpc:
	k6 run test-grpc.js