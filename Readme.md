# Vets

Vets is a concept flutter app that allows farmers to search for veterinary officers in their area, checkout their services and schedule appointments .This is the backend for the flutter app ([https://github.com/ndirangug/vet_flutter](https://github.com/ndirangug/vet_flutter)). It is a grpc service built with go.
The grpc service runs CRUD operations against a postgresql database.
  
## Technologies Used 
- Golang
- GRPC
- Docker was  used to continerize the application and deploy it to Google Cloud Run.
- Makfile was used to automate some repetitive tasks on local dev environement


## Cloud Deployment
The Go grpc service is deployed to a Google Cloud Run instance and CI pipeline has been setup with Google Cloud Build and Github.A Heroku Postgres Instance has been used for the database.
### Test the grpc service on cloud run
The [grpcurl](https://github.com/fullstorydev/grpcurl) tool may be used to send requests to the hosted grpc service on cloud run
```bash
# server reflection
$ grpcurl  vet-backend-fybfguvuua-uc.a.run.app:443 list

# example get a veterinry officers within  given radius of a particular location
$ grpcurl  -d '{"location":{"lat":-1.2939599460360305, "long": 36.799532813964184}, "radius": 300}'  vet-backend-fybfguvuua-uc.a.run.app:443 vet_backend.VetsBackend/GetVeterinariansInLocation
```
The [grpui](https://github.com/fullstorydev/grpcurl) tool may also be used to graphicly explore the service and play with it
```bash
$ grpcui  vet-backend-fybfguvuua-uc.a.run.app:443 
```



## Test on localhost
### Prerequsites
- Go 1.15
- Postgresql 13

### Run Server

```bash
# Setup Database locally
$ make migrate
$ make seed

# start the grpc service
$ make serve
```

