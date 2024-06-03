package task

import (
	"fmt"
	"strings"

	"github.com/TudorHulban/authentication/apperrors"
)

type Tasks []*Task

func (t Tasks) GetTaskByID(pk PrimaryKeyTask) (*Task, error) {
	for _, task := range t {
		if task.PrimaryKeyTask == pk {
			return task, nil
		}
	}

	return nil, apperrors.ErrEntryNotFound{
		Key: "GetTaskByID - PrimaryKeyTask",
	}
}

func (t Tasks) String() string {
	result := []string{
		fmt.Sprintf("Tasks: %d", len(t)),
	}

	for _, task := range t {
		result = append(result,
			fmt.Sprintf(
				"ID: %v, Name: %s",
				task.PrimaryKeyTask,
				task.Name,
			),
		)
	}

	return strings.Join(
		result,
		"\n",
	)
}
