package server

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/mgsquare/grpc-report-generation/pb"
)

type ReportServer struct {
	pb.UnimplementedReportServiceServer
	mu      sync.Mutex
	reports map[string]string
}

func NewReportServer() *ReportServer {
	return &ReportServer{
		reports: make(map[string]string),
	}
}

func (s *ReportServer) GenerateReport(ctx context.Context, req *pb.GenerateReportRequest) (*pb.GenerateReportResponse, error) {
	log.Printf("[%s] GenerateReport called for user: %s\n", time.Now().Format(time.RFC3339), req.UserId)

	reportID := uuid.New().String()

	s.mu.Lock()
	s.reports[reportID] = req.UserId
	s.mu.Unlock()

	log.Printf("[%s] Report generated: %s for user %s\n", time.Now().Format(time.RFC3339), reportID, req.UserId)

	return &pb.GenerateReportResponse{
		ReportId: reportID,
		Error:    "",
	}, nil
}

func (s *ReportServer) HealthCheck(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	log.Printf("[%s] HealthCheck called\n", time.Now().Format(time.RFC3339))
	return &pb.HealthCheckResponse{
		Status: "ok",
	}, nil
}
