generate:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protos/service.proto

reflect:
    grpcurl --plaintext localhost:50051 list




# test:
#     grpcurl --plaintext -d '{"firstName":"", "lastName":"", "email":"", "phone":"", "address" :{"lat":-23.87876, "long": 36.678678676, "address": "nairobi"}, "photo": []}'  localhost:50051 vet_backend.VetsBackend/UpdateFarmer



