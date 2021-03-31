generate:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protos/service.proto

reflect:
    grpcurl --plaintext localhost:50051 list




# test:
#     grpcurl --plaintext -d '{"farmerId": "id", "email": "f@gmail.com"}'  localhost:50051 vet_backend.VetsBackend/GetFarmer


