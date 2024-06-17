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

type ParamsAsHTMLTBody struct {
	RouteTicket     string
	CSSIDTicketBody string
}

func (t Tickets) AsHTMLTBody(params ParamsAsHTMLTBody) string {
	result := []string{
		fmt.Sprintf(
			"<tbody class=%s>",
			params.CSSIDTicketBody,
		),
	}

	for ix, ticket := range t {
		result = append(
			result,
			ticket.AsHTMLTRow(
				&ParamsTicketAsHTML{
					RouteTicket: params.RouteTicket,
					Index:       ix + 1,
				},
			),
		)
	}

	result = append(result, "</tbody>")

	return strings.Join(result, "")
}
