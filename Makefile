generate:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protos/service.proto

reflect:
    grpcurl --plaintext localhost:50051 list




# test:
#     grpcurl --plaintext -d '{"farmerId": 8, "firstName":"George", "lastName":"Ndirangu", "email":"ndirangu.mepawa@gmail.com", "phone":"", "address" :{"lat":-23.87876, "long": 36.678678676, "address": "nairobi"}, "photo": []}'  localhost:50051 vet_backend.VetsBackend/UpdateFarmer
#     grpcurl --plaintext -d '{"vetId":1}'  localhost:50051 vet_backend.VetsBackend/GetVeterinarian



