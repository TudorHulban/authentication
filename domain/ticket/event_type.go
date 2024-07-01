package ticket

import (
	"github.com/TudorHulban/authentication/helpers"
)

type EventType uint8

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

var setEventTypeUS = helpers.NewImmutableSetFrom[EventType, string](
	[]helpers.KV[EventType, string]{
		{
			Key:   1,
			Value: "Open",
		},
		{
			Key:   2,
			Value: "WIP",
		},
		{
			Key:   3,
			Value: "Analysis",
		},
		{
			Key:   4,
			Value: "NI",
		},
		{
			Key:   5,
			Value: "WFI",
		},
		{
			Key:   6,
			Value: "Resolution",
		},
		{
			Key:   7,
			Value: "With 3rd party",
		},
		{
			Key:   8,
			Value: "Blocking",
		},
		{
			Key:   9,
			Value: "Unblocking",
		},
		{
			Key:   10,
			Value: "Escalation Internal",
		},
		{
			Key:   11,
			Value: "Escalation Customer",
		},
		{
			Key:   12,
			Value: "Close",
		},
	},
)

var SetEventType = setEventTypeUS

func GetStringStatusFor(numeric EventType) string {
	value, exists := setEventTypeUS.Get(numeric)
	if !exists {
		return "unknown status" //TODO: add constant
	}

	return value
}
