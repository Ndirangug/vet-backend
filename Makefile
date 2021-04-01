generate:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protos/service.proto

reflect:
    grpcurl --plaintext localhost:50051 list




# test:
#     grpcurl --plaintext -d '{"farmerId": 8, "firstName":"George", "lastName":"Ndirangu", "email":"ndirangu.mepawa@gmail.com", "phone":"", "address" :{"lat":-23.87876, "long": 36.678678676, "address": "nairobi"}, "photo": []}'  localhost:50051 vet_backend.VetsBackend/UpdateFarmer
#     grpcurl --plaintext -d '{"vetId":1}'  localhost:50051 vet_backend.VetsBackend/GetVeterinarian
#     grpcurl --plaintext -d '{"time":62135596800, "location":{"lat":-23.87876, "long": 36.678678676}, "farmerId" :8, "VeterinaryId": 1, "services": [{"serviceId":1, "units":10}, {"serviceId":1, "units":20}]}'  localhost:50051 vet_backend.VetsBackend/ScheduleSession
#     grpcurl --plaintext -d '{"time":"2021-04-01 16:24:17.876", "location":{"lat":-23.87876, "long": 36.678678676}, "farmerId" :8, "VeterinaryId": 1, "services": [{"serviceId":1, "units":10}, {"serviceId":1, "units":20}]}'  localhost:50051 vet_backend.VetsBackend/ScheduleSession
#     grpcurl --plaintext -d '{"vetId": 1, "firstName":"George", "lastName":"Wachira", "email":"ndirangu.mepawa@gmail.com", "phone":"", "address" :{"lat":-23.87876, "long": 36.678678676, "address": "nairobi"}, "photo": []}'  localhost:50051 vet_backend.VetsBackend/UpdateVeterian
#     grpcurl --plaintext -d '{"location":{"lat":-1.2939599460360305, "long": 36.799532813964184}, "radius": 300}'  localhost:50051 vet_backend.VetsBackend/GetVeterinariansInLocation

