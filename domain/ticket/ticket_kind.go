package ticket

type TicketKind struct {
	Kind uint8

	OpeningEventType uint8
	ClosingEventType uint8
}

var (
	KindTicket = TicketKind{
		Kind:             1,
		OpeningEventType: 1,
		ClosingEventType: 10,
	}

	KindSale = TicketKind{
		Kind:             2,
		OpeningEventType: 1,
		ClosingEventType: 10,
	}

	KindLead = TicketKind{
		Kind:             3,
		OpeningEventType: 1,
		ClosingEventType: 10,
	}
)
