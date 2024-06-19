package ticket

import "github.com/TudorHulban/authentication/helpers"

func GetIDTicket(item *Ticket) helpers.PrimaryKey {
	return item.PrimaryKey
}

var CriteriaPK = func(pk helpers.PrimaryKey) func(item *Ticket) bool {
	return func(item *Ticket) bool {
		return GetIDTicket(item) == pk
	}
}
