package service

import (
	"context"
	"fmt"

	"github.com/ndirangug/vets-backend/logger"
	"github.com/ndirangug/vets-backend/protos"
)

// Backend implements  grpc methods auto-generated from .proto file
type BackendService struct {
	logger *logger.TinyLogger
	protos.UnimplementedBebaBackendServer
}

// TinyErpGrpc returns a new server
func NewBackendService(logger *logger.TinyLogger) *BackendService {
	return &BackendService{logger: logger}
}

func (s *BackendService) TestHello(ctx context.Context, request *protos.TestHelloRequest) (*protos.TestHelloResponse, error) {
	return &protos.TestHelloResponse{Response: fmt.Sprintf("Hello %s", request.Name)}, nil
}

func (s *BackendService) GetDrivers(EmptyMessage *protos.EmptyMessage, stream protos.BebaBackend_GetDriversServer) error {
	drivers := []*protos.Driver{
		{IdNumber: "6", FirstName: "George", LastName: "Ndirangu"},
		{IdNumber: "7", FirstName: "Ndisho", LastName: "Heyy"},
	}

	for _, driver := range drivers {

		if err := stream.Send(driver); err != nil {
			return err
		}
	}

	return nil

}
