package ticket

// TODO: test for non existent ticket kinds
// TODO: move to swiss map
var TicketKindToEventType = map[TicketKind]map[EventType]*TicketEventTypeInfo{
	KindTicket: {
		EventTypeOpen: {
			DefaultEventTypeLevel: LevelEndUser,
			AllowedNextEventTypes: []EventType{
				EventTypeAssignTo,
				EventTypeWith3rdParty,
				EventTypeBlocks,
				EventTypeUnBlocks,
				EventTypeEscalationInternal,
				EventTypeEscalationCustomer,
				EventTypeClose,
			},
		},
		EventTypeAssignTo: {
			DefaultEventTypeLevel: LevelEndUser,
			AllowedNextEventTypes: []EventType{
				EventTypeNoteInternal,
				EventTypeWith3rdParty,
				EventTypeBlocks,
				EventTypeUnBlocks,
				EventTypeEscalationInternal,
				EventTypeEscalationCustomer,
				EventTypeClose,
			},
		},
		EventTypeWorkInProgress: {
			DefaultEventTypeLevel: LevelEndUser,
			AllowedNextEventTypes: []EventType{
				EventTypeNoteInternal,
				EventTypeWith3rdParty,
				EventTypeBlocks,
				EventTypeUnBlocks,
				EventTypeEscalationInternal,
				EventTypeEscalationCustomer,
				EventTypeClose,
			},
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
