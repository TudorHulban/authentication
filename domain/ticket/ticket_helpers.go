package ticket

import "github.com/TudorHulban/authentication/helpers"

func GetIDTicket(item *Ticket) uint64 {
	return uint64(item.PrimaryKey)
}

var CriteriaPK = func(pk helpers.PrimaryKey) func(item *Ticket) bool {
	return func(item *Ticket) bool {
		return GetIDTicket(item) == uint64(pk)
	}
}
