postgres:
	export DATABASE_URL="host=localhost user=vets_backend password=vets dbname=vets port=5432 sslmode=disable TimeZone=Africa/Nairobi" && \
	sudo service postgresql start

push_db:
	export DATABASE_URL=postgres://rhdktlxtufaaoe:22971c5d488da6a8e27011251b24f2d99eff8af796deeaaec49b13efb557330c@ec2-52-1-115-6.compute-1.amazonaws.com:5432/da5kla8im8n293 && \
	heroku pg:reset DATABASE_URL --app vets-backend
	heroku pg:push vets DATABASE_URL --app vets-backend

serve:
	export DATABASE_URL=postgres://rhdktlxtufaaoe:22971c5d488da6a8e27011251b24f2d99eff8af796deeaaec49b13efb557330c@ec2-52-1-115-6.compute-1.amazonaws.com:5432/da5kla8im8n293 && \
	export PORT=:50051 && \
	go run main.go serve

seed:
	make postgres && \
	export DATABASE_URL="host=localhost user=vets_backend password=vets dbname=vets port=5432 sslmode=disable TimeZone=Africa/Nairobi" && \
	go run main.go seed && \
	xterm -e make push_db

migrate:
	make postgres && \
	export DATABASE_URL="host=localhost user=vets_backend password=vets dbname=vets port=5432 sslmode=disable TimeZone=Africa/Nairobi" && \
	go run main.go migrate && \
	xterm -e make push_db

# generate:
# 	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/service.proto


# reflect:
#     grpcurl --plaintext localhost:50051 list

# push_db:
# 	heroku pg:push vets DATABASE_URL --app vets-backend

# pg_credentials:
# 	heroku pg:credentials:url DATABASE --app vets-backend


# test:
#     grpcurl --plaintext -d '{"farmerId": 8, "firstName":"George", "lastName":"Ndirangu", "email":"ndirangu.mepawa@gmail.com", "phone":"", "address" :{"lat":-23.87876, "long": 36.678678676, "address": "nairobi"}, "photo": []}'  localhost:50051 vet_backend.VetsBackend/UpdateFarmer
#     grpcurl --plaintext -d '{"vetId":1}'  localhost:50051 vet_backend.VetsBackend/GetVeterinarian
#     grpcurl --plaintext -d '{"time":62135596800, "location":{"lat":-23.87876, "long": 36.678678676}, "farmerId" :8, "VeterinaryId": 1, "services": [{"serviceId":1, "units":10}, {"serviceId":1, "units":20}]}'  localhost:50051 vet_backend.VetsBackend/ScheduleSession
#     grpcurl --plaintext -d '{"time":"2021-04-01 16:24:17.876", "location":{"lat":-23.87876, "long": 36.678678676}, "farmerId" :8, "VeterinaryId": 1, "services": [{"serviceId":1, "units":10}, {"serviceId":1, "units":20}]}'  localhost:50051 vet_backend.VetsBackend/ScheduleSession
#     grpcurl --plaintext -d '{"vetId": 1, "firstName":"George", "lastName":"Wachira", "email":"ndirangu.mepawa@gmail.com", "phone":"", "address" :{"lat":-23.87876, "long": 36.678678676, "address": "nairobi"}, "photo": []}'  localhost:50051 vet_backend.VetsBackend/UpdateVeterian
#     grpcurl --plaintext -d '{"location":{"lat":-1.2939599460360305, "long": 36.799532813964184}, "radius": 300}'  localhost:50051 vet_backend.VetsBackend/GetVeterinariansInLocation
#     grpcurl --plaintext -d '{"location":{"lat":-1.2939599460360305, "long": 36.799532813964184}, "radius": 300}' vet-backend-fybfguvuua-uc.a.run.app:443 vet_backend.VetsBackend/GetVeterinariansInLocation
#     grpcurl  -d '{"location":{"lat":-1.2939599460360305, "long": 36.799532813964184}, "radius": 300}'  vet-backend-fybfguvuua-uc.a.run.app:443 vet_backend.VetsBackend/GetVeterinariansInLocation
#	grpcurl  vet-backend-fybfguvuua-uc.a.run.app:50051 list
#	grpcui  vet-backend-fybfguvuua-uc.a.run.app:50051
