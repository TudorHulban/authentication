package task

import (
	"fmt"
	"strings"

	"github.com/TudorHulban/authentication/apperrors"
)

type Tickets []*Ticket

func (t Tickets) GetTaskByID(pk PrimaryKeyTicket) (*Ticket, error) {
	for _, task := range t {
		if task.PrimaryKeyTicket == pk {
			return task, nil
		}
	}

	return nil, apperrors.ErrEntryNotFound{
		Key: "GetTaskByID - PrimaryKeyTask",
	}
}

func (t Tickets) String() string {
	result := []string{
		fmt.Sprintf("Tasks: %d", len(t)),
	}

	for _, task := range t {
		result = append(result,
			fmt.Sprintf(
				"ID: %v, Name: %s",
				task.PrimaryKeyTicket,
				task.Name,
			),
		)
	}

	return strings.Join(
		result,
		"\n",
	)
}
