package ticket

import "github.com/TudorHulban/authentication/helpers"

const (
	EventTypeInit         = "open ticket"
	EventTypeInternalNote = "internal note"
	EventTypeAnswer       = "answer"
	EventType3rdParty     = "3rd party"
	EventTypeBlocks       = "blocking"
	EventTypeUnBlocks     = "unblocking"
)

var setEventType = helpers.NewImmutableSetFrom[uint8, string](
	[]helpers.KV[uint8, string]{
		{
			Key:   0,
			Value: EventTypeInternalNote,
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
	},
)

var SetEventType = setEventType
