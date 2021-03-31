package service

import (
	"context"
	"fmt"

	"github.com/ndirangug/vets-backend/db"
	"github.com/ndirangug/vets-backend/logger"
	"github.com/ndirangug/vets-backend/models"
	"github.com/ndirangug/vets-backend/protos"
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

func (s *BackendService) GetVeterinarians(request *protos.VetRequest, stream protos.VetsBackend_GetVeterinariansServer) error {
	vets := []*protos.Veterinary{}

	for _, vet := range vets {

		if err := stream.Send(vet); err != nil {
			return err
		}
	}

	return nil
}

func (s *BackendService) UpdateVeterian(ctx context.Context, vet *protos.Veterinary) (*protos.Veterinary, error) {
	result := &protos.Veterinary{}

	return result, nil
}

func (s *BackendService) UpdateFarmer(ctx context.Context, farmer *protos.Farmer) (*protos.Farmer, error) {
	dbFarmer := &models.Farmer{FirstName: farmer.FirstName, LastName: farmer.LastName, Email: farmer.Email, Phone: farmer.Phone, Address: farmer.Address.Address, Longitude: farmer.Address.Long, Latitude: farmer.Address.Lat}

	result := s.db.Conn.Create(&dbFarmer)

	farmer.FarmerId = int32(dbFarmer.ID)

	return farmer, result.Error
}

func (s *BackendService) GetFarmer(ctx context.Context, request *protos.FarmerRequest) (*protos.Farmer, error) {
	result := &protos.Farmer{}

	return result, nil
}

func (s *BackendService) ScheduleSession(ctx context.Context, req *protos.TreatmentSession) (*protos.TreatmentSession, error) {
	result := &protos.TreatmentSession{}

	return result, nil
}
