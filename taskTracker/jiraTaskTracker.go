package taskTracker

import (
	"context"
	"github.com/vlamai/papir_jira/proto/pb"
)

type JiraTaskTracker struct {
	Client pb.TaskTrackerClient
}

func (j JiraTaskTracker) GetTaskName(taskID string) (string, error) {
	request := pb.GetTaskNameRequest{TaskId: taskID}
	response, err := j.Client.GetTaskName(context.Background(), &request)
	if err != nil {
		return "", err
	}
	return response.TaskName, err
}
