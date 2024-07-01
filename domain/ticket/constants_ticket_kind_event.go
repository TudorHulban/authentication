package ticket

var TicketKindToEventType = map[TicketKind]map[EventType]*TicketEventTypeInfo{
	KindTicket: {
		EventTypeOpen: {
			DefaultEventTypeLevel: LevelEndUser,
			AllowedNextEventTypes: []uint8{
				2,
				7,
				8,
				9,
				10,
				11,
				12,
			},
		},
		EventTypeWorkInProgress: {
			DefaultEventTypeLevel: LevelEndUser,
		},
		EventTypeAnalysis: {
			DefaultEventTypeLevel: LevelTeam,
		},
		EventTypeNoteInternal: {
			DefaultEventTypeLevel: LevelTeam,
		},
		EventTypeWaitingFutherInformation: {
			DefaultEventTypeLevel: LevelEndUser,
		},
		EventTypeResolution: {
			DefaultEventTypeLevel: LevelEndUserManager,
		},
		EventTypeWith3rdParty: {
			DefaultEventTypeLevel: LevelEndUserManager,
		},
		EventTypeBlocks: {
			DefaultEventTypeLevel: LevelTeam,
		},
		EventTypeUnBlocks: {
			DefaultEventTypeLevel: LevelTeam,
		},
		EventTypeEscalationInternal: {
			DefaultEventTypeLevel: LevelEndUserVIP,
		},
		EventTypeClose: {
			DefaultEventTypeLevel: LevelEndUserManager,
		},
	},

	KindSale: map[EventType]*TicketEventTypeInfo{
		EventTypeOpen: &TicketEventTypeInfo{
			DefaultEventTypeLevel: 1,
		},
	},

	KindLead: map[EventType]*TicketEventTypeInfo{
		EventTypeOpen: &TicketEventTypeInfo{
			DefaultEventTypeLevel: 1,
		},
	},
}
