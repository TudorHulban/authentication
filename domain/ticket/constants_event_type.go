package ticket

const (
	LevelEndUser        = 1
	LevelEndUserManager = 2
	LevelEndUserVIP     = 3

	LevelTeam           = 5
	LevelAgent          = 6
	LevelLead           = 7
	LevelTeamManager    = 9
	LevelProjectManager = 11
	LevelGroupManager   = 14
	LevelAccountManager = 16
	LevelC              = 19
)

const (
	EventTypeOpen                     = EventType(1)
	EventTypeWorkInProgress           = EventType(2)
	EventTypeAnalysis                 = EventType(3)
	EventTypeNoteInternal             = EventType(4)
	EventTypeWaitingFutherInformation = EventType(5)
	EventTypeResolution               = EventType(6)
	EventTypeWith3rdParty             = EventType(7)
	EventTypeBlocks                   = EventType(8)
	EventTypeUnBlocks                 = EventType(9)
	EventTypeEscalationInternal       = EventType(10)
	EventTypeEscalationCustomer       = EventType(11)
	EventTypeClose                    = EventType(12)
)

const msgUnknownEventType = "unknown event type"
