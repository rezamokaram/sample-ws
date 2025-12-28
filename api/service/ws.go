package service

import (
	"context"

	"github.com/rezamokaram/sample-ws/api/pb"
	producerPort "github.com/rezamokaram/sample-ws/internal/producer/port"
)

type StreamService struct {
	svc producerPort.Service
}

func NewStreamService(svc producerPort.Service) *StreamService {
	return &StreamService{
		svc: svc,
	}
}

func (s *StreamService) GetStream(ctx context.Context) (*pb.Stream, error) {
	client, err := s.svc.NewClient()
	if err != nil {
		return nil, err
	}

	return &pb.Stream{
		Channel: client.Channel,
	}, nil
}
