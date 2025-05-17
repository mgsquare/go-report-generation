package main

import (
	"log"
	"net"

	"github.com/mgsquare/grpc-report-generation/cron"
	"github.com/mgsquare/grpc-report-generation/pb"
	"github.com/mgsquare/grpc-report-generation/server"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	reportServer := server.NewReportServer()

	grpcServer := grpc.NewServer()
	pb.RegisterReportServiceServer(grpcServer, reportServer)

	userIDs := []string{"user1", "user2", "user3"}

	cron.StartReportGenerationCron(reportServer, userIDs)

	log.Println("gRPC server started on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
