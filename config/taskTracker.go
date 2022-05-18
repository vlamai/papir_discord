package config

import (
	"github.com/vlamai/papir_discord/taskTracker"
	"github.com/vlamai/papir_jira/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewJira() (taskTracker.JiraTaskTracker, *grpc.ClientConn) {
	conn, err := grpc.Dial(getEnvVar("PAPIR_TASK_TRACKER_SERVER"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewTaskTrackerClient(conn)
	return taskTracker.JiraTaskTracker{Client: client}, conn
}
