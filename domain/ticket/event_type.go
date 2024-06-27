package ticket

import "github.com/TudorHulban/authentication/helpers"

// TODO: i18n
const (
	EventTypeInit         = "open"
	EventTypeInternalNote = "internal note"
	EventTypeAnswer       = "answer"
	EventType3rdParty     = "3rd party"
	EventTypeBlocks       = "blocking"
	EventTypeUnBlocks     = "unblocking"
	EventTypeEscalation   = "escalation"
	EventTypeClose        = "closure"
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
			Value: EventTypeAnswer,
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
