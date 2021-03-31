package service

import (
	"context"
	"fmt"
	"time"

	"github.com/ndirangug/vets-backend/db"
	"github.com/ndirangug/vets-backend/logger"
	"github.com/ndirangug/vets-backend/models"
	"github.com/ndirangug/vets-backend/protos"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

// Backend implements  grpc methods auto-generated from .proto file
type BackendService struct {
	logger *logger.TinyLogger
	protos.UnimplementedVetsBackendServer
	db *db.DbConnection
}

// TinyErpGrpc returns a new server
func NewBackendService(logger *logger.TinyLogger) *BackendService {
	conn, err := db.NewDbConnection()

	if err != nil {
		logger.Panic("failed to connect to database. Err: %s", err)
	}

	return &BackendService{logger: logger, db: conn}
}

func (s *BackendService) TestHello(ctx context.Context, request *protos.TestHelloRequest) (*protos.TestHelloResponse, error) {
	return &protos.TestHelloResponse{Response: fmt.Sprintf("Hello %s", request.Name)}, nil
}

func (s *BackendService) GetVeterinarian(ctx context.Context, request *protos.VetRequest) (*protos.Veterinary, error) {
	dbVet := &models.Veterinary{}
	var services []*protos.VetService

	result := s.db.Conn.First(&dbVet, request.VetId)

	services = fetchServices(s, int(request.VetId))
	vet := &protos.Veterinary{VetId: int32(dbVet.ID), Title: dbVet.Title, FirstName: dbVet.FirstName, LastName: dbVet.LastName, Email: dbVet.Email, Phone: dbVet.Phone, Address: &protos.Location{Lat: dbVet.Latitude, Long: dbVet.Longitude, Address: dbVet.Address}, Summary: "", Services: services}

	return vet, result.Error
}

func (s *BackendService) GetVeterinarians(request *protos.VetRequest, stream protos.VetsBackend_GetVeterinariansServer) error {
	var dbVets []models.Veterinary

	result := s.db.Conn.Where("first_name = ?", request.SearchQuery).Find(&dbVets)

	if result.Error != nil {
		s.logger.Info("Record not found, %s", result.Error)
	}

	for _, dbVet := range dbVets {
		var services []*protos.VetService
		services = fetchServices(s, int(dbVet.ID))
		vet := &protos.Veterinary{VetId: int32(dbVet.ID), Title: dbVet.Title, FirstName: dbVet.FirstName, LastName: dbVet.LastName, Email: dbVet.Email, Phone: dbVet.Phone, Address: &protos.Location{Lat: dbVet.Latitude, Long: dbVet.Longitude, Address: dbVet.Address}, Summary: "", Services: services}
		if err := stream.Send(vet); err != nil {
			return err
		}
	}

	return nil
}

func (s *BackendService) GetVeterinariansInLocation(location *protos.Location, stream protos.VetsBackend_GetVeterinariansInLocationServer) error {
	var dbVets []models.Veterinary

	s.db.Conn.Where("latitude < ?", location.Lat).Find(&dbVets) //TODO USE SOMETHING LIKE DISTANCE VECTOR API HERE

	for _, dbVet := range dbVets {
		var services []*protos.VetService
		services = fetchServices(s, int(dbVet.ID))
		vet := &protos.Veterinary{VetId: int32(dbVet.ID), Title: dbVet.Title, FirstName: dbVet.FirstName, LastName: dbVet.LastName, Email: dbVet.Email, Phone: dbVet.Phone, Address: &protos.Location{Lat: dbVet.Latitude, Long: dbVet.Longitude, Address: dbVet.Address}, Summary: "", Services: services}
		if err := stream.Send(vet); err != nil {
			return err
		}
	}

	return nil
}

func (s *BackendService) CreateVeterian(ctx context.Context, vet *protos.Veterinary) (*protos.Veterinary, error) {
	dbVet := &models.Veterinary{FirstName: vet.FirstName, LastName: vet.LastName, Email: vet.Email, Phone: vet.Phone, Address: vet.Address.Address, Longitude: vet.Address.Long, Latitude: vet.Address.Lat}

	result := s.db.Conn.Create(&dbVet)

	vet.VetId = int32(dbVet.ID)

	return vet, result.Error
}

func (s *BackendService) UpdateVeterian(ctx context.Context, vet *protos.Veterinary) (*protos.Veterinary, error) {
	dbVet := &models.Veterinary{Model: gorm.Model{ID: uint(vet.VetId), UpdatedAt: time.Now()}, FirstName: vet.FirstName, LastName: vet.LastName, Email: vet.Email, Phone: vet.Phone, Address: vet.Address.Address, Longitude: vet.Address.Long, Latitude: vet.Address.Lat}

	result := s.db.Conn.UpdateColumns(&dbVet)

	vet.VetId = int32(dbVet.ID)

	return vet, result.Error
}

func (s *BackendService) CreateFarmer(ctx context.Context, farmer *protos.Farmer) (*protos.Farmer, error) {
	dbFarmer := &models.Farmer{FirstName: farmer.FirstName, LastName: farmer.LastName, Email: farmer.Email, Phone: farmer.Phone, Address: farmer.Address.Address, Longitude: farmer.Address.Long, Latitude: farmer.Address.Lat}

	result := s.db.Conn.Create(&dbFarmer)

	farmer.FarmerId = int32(dbFarmer.ID)

	return farmer, result.Error
}

func (s *BackendService) UpdateFarmer(ctx context.Context, farmer *protos.Farmer) (*protos.Farmer, error) {
	dbFarmer := &models.Farmer{Model: gorm.Model{ID: uint(farmer.FarmerId), UpdatedAt: time.Now()}, FirstName: farmer.FirstName, LastName: farmer.LastName, Email: farmer.Email, Phone: farmer.Phone, Address: farmer.Address.Address, Longitude: farmer.Address.Long, Latitude: farmer.Address.Lat}

	result := s.db.Conn.UpdateColumns(&dbFarmer)

	farmer.FarmerId = int32(dbFarmer.ID)

	return farmer, result.Error
}

func (s *BackendService) GetFarmer(ctx context.Context, request *protos.FarmerRequest) (*protos.Farmer, error) {
	var dbFarmer models.Farmer

	s.db.Conn.Where("email = ?", request.Email).First(&dbFarmer)
	farmer := &protos.Farmer{FarmerId: int32(dbFarmer.ID), FirstName: dbFarmer.FirstName, LastName: dbFarmer.LastName, Email: dbFarmer.Email, Phone: dbFarmer.Phone, Address: &protos.Location{Lat: dbFarmer.Latitude, Long: dbFarmer.Longitude, Address: dbFarmer.Address}}
	return farmer, nil
}

func (s *BackendService) ScheduleSession(ctx context.Context, req *protos.TreatmentSessionRequest) (*protos.TreatmentSession, error) {
	vetServiceSessions := []*models.VetServiceSession{}

	for _, service := range req.Services {
		vetServiceSessions = append(vetServiceSessions, &models.VetServiceSession{VetServiceID: uint(service.ServiceId), Units: uint(service.Units)})
	}
	dbSession := &models.Session{Date: req.Time.AsTime(), Latitude: req.Location.Lat, Longitude: req.Location.Long, FarmerID: uint(req.FarmerId), VeterinaryID: uint(req.VeterinaryId), VetServiceSessions: vetServiceSessions}

	result := s.db.Conn.Create(&dbSession)

	session := &protos.TreatmentSession{SessionId: int32(dbSession.ID), Time: timestamppb.New(dbSession.Date), Location: &protos.Location{Lat: dbSession.Latitude, Long: dbSession.Longitude}, FarmerId: uint32(dbSession.FarmerID), VeterinaryId: uint32(dbSession.VeterinaryID)}
	return session, result.Error
}

func fetchServices(s *BackendService, vetId int) []*protos.VetService {
	var services []*protos.VetService
	var dbVetServices []models.VetService
	s.db.Conn.Where("veterinary_id = ?", vetId).Find(&dbVetServices)

	for _, dbVetService := range dbVetServices {
		var dbService models.Service
		s.db.Conn.First(&dbService, dbVetService.ServiceID)
		services = append(services, &protos.VetService{ServiceId: int32(dbVetService.ID), Title: dbService.Name, Description: dbService.Description, Unit: dbVetService.Unit, CostPerUnit: dbVetService.ChargePerUnit})
	}

	return services
}
