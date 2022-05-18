package message

import (
	"net/url"
	"path"
)

type TaskTracker interface {
	GetTaskName(taskID string) (string, error)
}

type TaskLink struct {
	Host    string
	Tracker TaskTracker
}

func (tl TaskLink) Parse(c *url.URL) (string, error) {
	if c.Host != tl.Host {
		return "", errHostNotMath
	}
	taskName, err := tl.Tracker.GetTaskName(path.Base(c.Path))
	if err != nil {
		return "", err
	}
	return taskName, nil
}
