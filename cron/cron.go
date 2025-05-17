package cron

import (
	"context"
	"log"
	"time"

	"github.com/mgsquare/grpc-report-generation/pb"
	"github.com/mgsquare/grpc-report-generation/server"
	"github.com/robfig/cron/v3"
)

func StartReportGenerationCron(srv *server.ReportServer, userIDs []string) *cron.Cron {
	c := cron.New(cron.WithSeconds())

	_, err := c.AddFunc("@every 10s", func() {
		for _, userID := range userIDs {

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			resp, err := srv.GenerateReport(ctx, &pb.GenerateReportRequest{UserId: userID})
			if err != nil {
				log.Printf("[%s] Failed to generate report for user %s: %v\n", time.Now().Format(time.RFC3339), userID, err)
				continue
			}
			log.Printf("[%s] Cron generated report %s for user %s\n", time.Now().Format(time.RFC3339), resp.ReportId, userID)
		}
	})

	if err != nil {
		log.Fatalf("Failed to add cron job: %v", err)
	}

	c.Start()
	log.Println("Cron job started: generating reports every 10 seconds")

	return c
}
