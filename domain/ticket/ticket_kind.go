package ticket

type TicketKind struct {
	Kind uint8

	OpeningEventType EventType
	ClosingEventType EventType
}

var (
	KindTicket = TicketKind{
		Kind:             1,
		OpeningEventType: EventTypeOpen,
		ClosingEventType: EventTypeClose,
	}

	KindSale = TicketKind{
		Kind:             2,
		OpeningEventType: EventTypeOpen,
		ClosingEventType: EventTypeClose,
	}

	KindLead = TicketKind{
		Kind:             3,
		OpeningEventType: EventTypeOpen,
		ClosingEventType: EventTypeClose,
	}
)
