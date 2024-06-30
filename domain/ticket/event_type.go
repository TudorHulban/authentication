package ticket

import (
	"github.com/TudorHulban/authentication/helpers"
)

// TODO: i18n
const (
	EventTypeInit                     = "Open"
	EventTypeInWork                   = "In work"
	EventTypeAnalysis                 = "Analysis"
	EventTypeInternalNote             = "Internal note"
	EventTypeWaitingFutherInformation = "WFI"
	EventTypeResolution               = "Resolution"
	EventType3rdParty                 = "3rd party"
	EventTypeBlocks                   = "Blocking"
	EventTypeUnBlocks                 = "Unblocking"
	EventTypeEscalation               = "Escalation"
	EventTypeClose                    = "Closure"
)

var setEventType = helpers.NewImmutableSetFrom[uint8, string](
	[]helpers.KV[uint8, string]{
		{
			Key:   0,
			Value: EventTypeInit,
		},
		{
			Key:   1,
			Value: EventTypeInternalNote,
		},
		{
			Key:   2,
			Value: EventTypeResolution,
		},
		{
			Key:   3,
			Value: EventType3rdParty,
		},
		{
			Key:   4,
			Value: EventTypeBlocks,
		},
		{
			Key:   5,
			Value: EventTypeUnBlocks,
		},
		{
			Key:   6,
			Value: EventTypeEscalation,
		},
		{
			Key:   7,
			Value: EventTypeClose,
		},
	},
)

var SetEventType = setEventType
