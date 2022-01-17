package events

// ApplicationStateTransitionEvent struct
type ApplicationStateTransitionEvent struct {
	Data      ApplicationStateTransitionEventData `json:"data,omitempty"`
	EventType EventType                           `json:"event_type,omitempty"`
}

func NewApplicationStateTransitionEvent() ApplicationStateTransitionEvent {
	return ApplicationStateTransitionEvent{
		Data:      ApplicationStateTransitionEventData{},
		EventType: ApplicationStateTransitionEventType,
	}
}

func (e ApplicationStateTransitionEvent) data() interface{} {
	return e.Data
}

func (e ApplicationStateTransitionEvent) eventType() EventType {
	return e.EventType
}
