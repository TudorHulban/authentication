package storememory

import (
	"fmt"
	"strings"

	"github.com/TudorHulban/authentication/domain/ticket"
)

type cacheTask map[ticket.PrimaryKeyTicket]ticket.TicketInfo

func (cache cacheTask) String() string {
	result := []string{
		fmt.Sprintf("cache Tasks: %d", len(cache)),
	}

	for pk, taskInfo := range cache {
		result = append(result,
			fmt.Sprintf("id: %v, name: %s", pk, taskInfo.Name),
		)
	}

	return strings.Join(result, "\n")
}
