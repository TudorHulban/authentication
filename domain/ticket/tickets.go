package ticket

import (
	"fmt"
	"strings"

	"github.com/TudorHulban/authentication/apperrors"
	"github.com/TudorHulban/authentication/helpers"
)

type Tickets []*Ticket

func (t Tickets) GetTaskByID(pk helpers.PrimaryKey) (*Ticket, error) {
	for _, task := range t {
		if task.PrimaryKey == pk {
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

	for _, ticket := range t {
		result = append(result,
			fmt.Sprintf(
				"ID: %v, Name: %s",
				ticket.PrimaryKey,
				ticket.Name,
			),
		)
	}

	return strings.Join(
		result,
		"\n",
	)
}
