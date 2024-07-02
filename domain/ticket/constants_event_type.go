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
	EventTypeOpen = EventType(iota + 1)
	EventTypeAssignTo
	EventTypeWorkInProgress
	EventTypeAnalysis
	EventTypeNoteInternal
	EventTypeWaitingFutherInformation
	EventTypeResolution
	EventTypeWith3rdParty
	EventTypeBlocks
	EventTypeUnBlocks
	EventTypeEscalationInternal
	EventTypeEscalationCustomer
	EventTypeClose
)

const msgUnknownEventType = "unknown event type"
